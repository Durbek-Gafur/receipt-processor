#!/bin/bash

# Stop and remove any container using port 8081
if [ $(docker ps -q -a -f name=receipt_processor_durbek) ]; then
    docker stop receipt_processor_durbek
    docker rm receipt_processor_durbek
fi

# Build the Docker image
docker build -t receipt_processor_durbek .

# Run the Docker container, mapping port 8081 of the container to port 8081 of the host machine
docker run -d -p 8081:8081 --name receipt_processor_durbek receipt_processor_durbek

# Print a message to indicate where the server can be accessed
echo "Server is running at http://localhost:8081"
