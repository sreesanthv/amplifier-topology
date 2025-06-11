package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Topology struct {
	NodeName string     `json:"node_name"`
	Amps     [][]string `json:"amps"`
}

var topologyData []Topology

func loadTopology(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &topologyData)
}

// GET /nodes
func getNodes(c *gin.Context) {
	var nodes []string
	for _, entry := range topologyData {
		nodes = append(nodes, entry.NodeName)
	}
	c.JSON(http.StatusOK, gin.H{"nodes": nodes})
}

// GET /node/:node_name/amps (Grafana format)
func getNodeAmps(c *gin.Context) {
	nodeName := c.Param("node_name")
	nodeMap := make(map[string]bool)

	nodes := []map[string]string{
		{"id": nodeName, "title": nodeName, "subTitle": "Node"},
	}

	for _, entry := range topologyData {
		if entry.NodeName == nodeName {
			for _, path := range entry.Amps {
				for _, amp := range path {
					if !nodeMap[amp] {
						nodeMap[amp] = true
						subTitle := "Amplifier"
						if amp == "amp 4" { // optional: treat amp 4 as splitter
							subTitle = "Splitter"
						}
						nodes = append(nodes, map[string]string{
							"id":       amp,
							"title":    amp,
							"subTitle": subTitle,
						})
					}
				}
			}
			break
		}
	}
	c.JSON(http.StatusOK, gin.H{"nodes": nodes})
}

// GET /node/:node_name/edges (Grafana format)
func getNodeEdges(c *gin.Context) {
	nodeName := c.Param("node_name")
	var edges []map[string]string
	idCounter := 0

	for _, entry := range topologyData {
		if entry.NodeName == nodeName {
			for _, path := range entry.Amps {
				if len(path) > 0 {
					// Connect node to first amp
					edges = append(edges, map[string]string{
						"id":     fmt.Sprintf("edge_%d", idCounter),
						"source": nodeName,
						"target": path[0],
					})
					idCounter++
				}
				for i := 0; i < len(path)-1; i++ {
					edges = append(edges, map[string]string{
						"id":     fmt.Sprintf("edge_%d", idCounter),
						"source": path[i],
						"target": path[i+1],
					})
					idCounter++
				}
			}
			break
		}
	}
	c.JSON(http.StatusOK, gin.H{"edges": edges})
}

func main() {
	if err := loadTopology("topology.json"); err != nil {
		log.Fatalf("Failed to load topology: %v", err)
	}

	router := gin.Default()

	router.GET("/nodes", getNodes)
	router.GET("/node/:node_name/amps", getNodeAmps)
	router.GET("/node/:node_name/edges", getNodeEdges)

	log.Println("Server running at http://localhost:8080")
	router.Run(":8080")
}
