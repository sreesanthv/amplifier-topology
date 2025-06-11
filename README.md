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

## 📊 Grafana Setup

### 1\. Install Grafana

You can run Grafana using Docker:

```
docker run -d -p 3000:3000 grafana/grafana
```

### 2\. Install `yesoreyeram-infinity-datasource`

*   Go to Grafana → **Configuration → Plugins**
*   Search for `Infinity` and install it.

### 3\. Add Infinity Data Source

*   Go to **Settings → Data Sources**
*   Select **Infinity**
*   Set **URL** `http://localhost:8080`
*   Save & Test

### 4\. Import Dashboard

*   Open Grafana → Dashboards → **Import**
*   Upload or paste the content of `grafana_dashboard_export.json`
*   Set Infinity data source → Import

### 5\. Use the Dashboard

*   Select a node from the dropdown
*   It loads corresponding amps and links in a Node Graph panel