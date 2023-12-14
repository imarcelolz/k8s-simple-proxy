# K8S Simple Proxy

This project utilizes Go, a statically typed, compiled language developed by Google. It also uses Kubernetes, a portable, extensible, open-source platform for managing containerized workloads and services.

## Features

- Uses `kubectl proxy` to create a proxy server that forwards traffic from a local address to the Kubernetes API Server.
- Uses `curl` to send requests to the local server.
- Includes a Jupyter Notebook setup using Docker.

## Requirements

- Go
- Kubernetes
- Docker
- Jupyter Notebook

## Installation & Usage

1. Start the Kubernetes proxy server:

```bash
kubectl proxy --port=8080 &
```

2 Send a request to the local server:

```bash
curl http://localhost:8080/api/
```

3. Start Jupyter Notebook using Docker:
```bash
docker run -it -p 8888:8888 gopherdata/gophernotes
```

## Contributing
Contributions are welcome. Please open an issue to discuss your idea or submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
