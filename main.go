package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Node struct {
	NodeName string     `json:"node_name"`
	Amps     [][]string `json:"amps"`
}

func getNodeAmps(c *gin.Context) {
	nodeName := c.Param("node_name")

	data, err := os.ReadFile("topology.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read topology"})
		return
	}

	var nodes []Node
	if err := json.Unmarshal(data, &nodes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse topology"})
		return
	}

	var result []map[string]interface{}
	ampSet := make(map[string]bool)

	for _, node := range nodes {
		if strings.EqualFold(node.NodeName, nodeName) {
			// Add the node itself
			result = append(result, map[string]interface{}{
				"id":       node.NodeName,
				"title":    node.NodeName,
				"subtitle": "Node",
				"color":    "green",
				"icon":     "server",
			})

			// Add amps
			for _, path := range node.Amps {
				for _, amp := range path {
					if !ampSet[amp] {
						result = append(result, map[string]interface{}{
							"id":       amp,
							"title":    amp,
							"subtitle": "Amp",
							"color":    "blue",
							"icon":     "bolt",
						})
						ampSet[amp] = true
					}
				}
			}
			break
		}
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Node not found or no amps"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func getNodeEdges(c *gin.Context) {
	nodeName := c.Param("node_name")

	data, err := os.ReadFile("topology.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read topology"})
		return
	}

	var nodes []Node
	if err := json.Unmarshal(data, &nodes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse topology"})
		return
	}

	var result []map[string]interface{}
	edgeIndex := 0

	for _, node := range nodes {
		if strings.EqualFold(node.NodeName, nodeName) {
			for _, path := range node.Amps {
				if len(path) > 0 {
					// Node to first amp
					result = append(result, map[string]interface{}{
						"id":            fmt.Sprintf("e%d", edgeIndex),
						"source":        node.NodeName,
						"target":        path[0],
						"mainStat":      fmt.Sprintf("%s → %s", node.NodeName, path[0]),
						"sourceArrow":   "none",
						"targetArrow":   "arrow",
						"secondaryStat": "from node",
					})
					edgeIndex++
				}
				for i := 0; i < len(path)-1; i++ {
					result = append(result, map[string]interface{}{
						"id":            fmt.Sprintf("e%d", edgeIndex),
						"source":        path[i],
						"target":        path[i+1],
						"mainStat":      fmt.Sprintf("%s → %s", path[i], path[i+1]),
						"sourceArrow":   "none",
						"targetArrow":   "arrow",
						"secondaryStat": "amp link",
					})
					edgeIndex++
				}
			}
			break
		}
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Node not found or no edges"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func getAllNodes(c *gin.Context) {
	data, err := os.ReadFile("topology.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read topology"})
		return
	}

	var nodes []Node
	if err := json.Unmarshal(data, &nodes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse topology"})
		return
	}

	var names []string
	for _, node := range nodes {
		names = append(names, node.NodeName)
	}

	c.JSON(http.StatusOK, names)
}

func main() {
	router := gin.Default()

	router.GET("/nodes", getAllNodes)
	router.GET("/node/:node_name/amps", getNodeAmps)
	router.GET("/node/:node_name/edges", getNodeEdges)

	router.Run(":8080")
}
