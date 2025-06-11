# 📡 Amplifier Topology API (Grafana Graph Visualization)

This is a simple mock API server built with Go + Gin to serve cable node topology (Node → RPD → Amplifiers) data in a format compatible with **Grafana's Node Graph Panel**.

## 🚀 Features

*   Serves topology from `topology.json`
*   Grafana-compatible endpoints for nodes and edges
*   REST endpoints to explore data

## 🔧 Getting Started

### 1\. Clone the Repository

```
git clone https://github.com/sreesanthv/amplifier-topology.git
cd amplifier-topology
```

### 2\. Build & Run with Docker Compose

```
docker-compose up --build
```

API will be running at: [http://localhost:8080](http://localhost:8080)

## 🧪 API Endpoints

### GET /nodes

Returns a list of all node names.

### GET /node/:node\_name/amps

Returns amplifier nodes in Grafana format.

### GET /node/:node\_name/edges

Returns amplifier edges in Grafana format.

### Example:

```
curl http://localhost:8080/node/Peringave/amps
curl http://localhost:8080/node/Peringave/edges
```

## 📊 Grafana Integration (using yesoreyeram-infinity-datasource)

1.  Install the **Infinity datasource plugin** in Grafana: [yesoreyeram-infinity-datasource](https://grafana.com/grafana/plugins/yesoreyeram-infinity-datasource).
2.  Go to **Connections → Data sources** and click **Add data source**.
3.  Select **Infinity**.
4.  Set type to **JSON** and choose **URL** as the source.
5.  Set the endpoint to:
    *   `http://localhost:8080/node/Peringave/amps` → for nodes
    *   `http://localhost:8080/node/Peringave/edges` → for edges
6.  Set **Format as**: `Table`
7.  Use in a **Node Graph Panel** with the following options:
    *   **Nodes query**: from `/amps` endpoint
    *   **Edges query**: from `/edges` endpoint

**Tip:** You may apply filters or dynamic variables using `${node}` in the Infinity URL to make it dynamic.