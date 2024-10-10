# Tracker Services

## Description
The Tracker Services provide data CRUD operations to track vehicles and send data to a Redis stream for live tracking purposes.

## Requirements
- Docker
- Docker Compose

## Getting Started
Follow these instructions to get the service up and running.

### Running the Service
1. Ensure you have Docker and Docker Compose installed on your machine.
2. Clone this repository to your local machine.
3. Navigate to the project directory.
4. Run the shell script to start the services:
    ```sh
    ./run.sh
    ```

### Testing
To testing unit-test you can enable test in docker-compose

## Features
- CRUD operations for vehicle tracking data.
- Live tracking data streamed to Redis.
- Utilizes MongoDB, which is not exposed outside the Docker private network and communicates only with other services within the network.
- Services do not expose ports; other services can connect to them via Docker's private network.

## Usage
Once the services are running, you can interact with the CRUD endpoints for vehicle tracking and listen to the Redis stream for live tracking