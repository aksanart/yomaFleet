# Vehicle Services

## Description
The Vehicle Services provide data CRUD operations to vehicles

## Requirements
- Docker
- Docker Compose
- Tracker Services up

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
- CRUD operations for vehicle data.
- Utilizes MongoDB, which is not exposed outside the Docker private network and communicates only with other services within the network.
- Services do not expose ports; other services can connect to them via Docker's private network.
