# Docker Setup for Receipt Processor

This repository contains the Docker setup for the Receipt Processor application. You can build and run the Docker container using a provided `run.sh` script or manually using Docker commands. Below are the instructions for both methods.

## Overview

This Docker setup builds and runs the Receipt Processor application, mapping port `8081` from the container to port `8081` on the host machine. The container name used in this setup is `receipt_processor_durbek`.

## Prerequisites

- Docker installed on your machine
- Docker Compose (optional, for additional management)
- Ensure the `run.sh` script has execute permissions (`chmod +x run.sh`)

## Using `run.sh` Script

The `run.sh` script automates the process of stopping any existing container, building the Docker image, and running the container. Follow these steps:

1. **Clone the Repository:**
   ```
   git clone https://github.com/Durbek-Gafur/receipt-processor.git
   cd receipt-processor
   ```

2. **Make the Script Executable:**
   Set the execute permissions for the `run.sh` script:
   ```
   chmod +x run.sh
   ```

3. **Run the Script:**
   Execute the script to build and run the Docker container:
   ```
   ./run.sh
   ```

4. **Verify the Container is Running:**
   Check the status of the container with:
   ```
   docker ps
   ```

   You should see the container `receipt_processor_durbek` running and mapping port `8081`.

5. **Access the Application:**
   Open your web browser and navigate to [http://localhost:8081](http://localhost:8081) to access the Receipt Processor application.

## Manual Build and Run Instructions

If you prefer to manage Docker manually, follow these instructions:

1. **Clone the Repository:**
   ```
   git clone https://github.com/Durbek-Gafur/receipt-processor.git
   cd receipt-processor
   ```

2. **Build the Docker Image:**
   Run the following command to build the Docker image:
   ```
   docker build -t receipt_processor_durbek .
   ```

3. **Run the Docker Container:**
   Start the Docker container with the following command:
   ```
   docker run -d -p 8081:8081 --name receipt_processor_durbek receipt_processor_durbek
   ```

4. **Verify the Container is Running:**
   Check the status of the container with:
   ```
   docker ps
   ```

   You should see the container `receipt_processor_durbek` running and mapping port `8081`.

5. **Access the Application:**
   Open your web browser and navigate to [http://localhost:8081](http://localhost:8081) to access the Receipt Processor application.

## Troubleshooting

- **Container Conflict:**
  If you encounter a conflict with the container name, ensure no existing container with the same name is running or stopped. Use the following commands to stop and remove the existing container:
  ```
  docker stop receipt_processor_durbek
  docker rm receipt_processor_durbek
  ```

- **Rebuild the Image:**
  If you make changes to the application code or Dockerfile, rebuild the Docker image with:
  ```
  docker build -t receipt_processor_durbek .
  ```

- **Check Container Logs:**
  If the application is not behaving as expected, view the container logs with:
  ```
  docker logs receipt_processor_durbek
  ```

## Future Improvements

1. **Better Error Handling:**
   - Implement comprehensive error handling in the application to provide clearer and more actionable error messages.
   - Capture and log errors to facilitate troubleshooting and debugging.

2. **Stack Trace Management:**
   - Ensure stack traces are included in error logs for debugging purposes.
   - Consider using a centralized logging system to collect and analyze logs from multiple containers.

3. **Environment Variables:**
   - Enhance the Docker setup to support environment variables for configuration. Update the Dockerfile and `run.sh` to accept environment variables and pass them to the container.
   - Example of setting environment variables in `docker run`:
     ```
     docker run -d -p 8081:8081 --name receipt_processor_durbek -e ENV_VAR_NAME=value receipt_processor_durbek
     ```

4. **Automated Testing:**
   - Integrate automated tests in the Docker build process to ensure code quality and functionality before deploying the container.

5. **Configuration Management:**
   - Use configuration management tools or service discovery mechanisms to handle different configurations for various environments (development, staging, production).

