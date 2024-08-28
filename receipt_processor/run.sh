#!/bin/bash

# Stop and remove any container using port 8081
docker ps -q --filter "publish=8081" | xargs -r docker stop | xargs -r docker rm

# Build the Docker image
docker build -t receipt_processor .

# Run the Docker container, mapping port 8081 of the container to port 8081 of the host machine
docker run -d -p 8081:8081 --name receipt_processor receipt_processor

# Print a message to indicate where the server can be accessed
echo "Server is running at http://localhost:8081"
