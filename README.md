# 📡 Amplifier Topology API (Grafana Graph Visualization)

This is a simple mock API server built with Go + Gin to serve cable node topology (Node → RPD → Amplifiers) data in a format compatible with **Grafana's Node Graph Panel**.

## 🚀 Features

*   Serves topology from `topology.json`
*   Grafana-compatible endpoints for nodes and edges
*   REST endpoints to explore data

## 🔧 Getting Started

### 1\. Clone the Repository

```
git clone <your-repo-url>
cd <project-directory>
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

## 📊 Grafana Integration

1.  Add **Node Graph Panel** in Grafana.
2.  Use a **Simple JSON** or **JSON API** plugin to connect to this API.
3.  Set the panel to consume the `/amps` and `/edges` endpoints.

## 🐳 Docker

### Dockerfile

```
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app main.go
CMD ["./app"]
```

### docker-compose.yml

```
version: '3.8'

services:
  api:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - .:/app
    restart: unless-stopped
```

## 📝 Sample Topology File (topology.json)

```
[
  {
    "node_name": "Peringave",
    "amps": [
      ["amp 1", "amp 2", "amp 3", "amp 4", "amp 5", "amp 6"],
      ["amp 1", "amp 2", "amp 3", "amp 4", "amp 7", "amp 8"]
    ]
  }
]
```