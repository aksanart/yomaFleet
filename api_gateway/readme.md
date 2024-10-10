# API Gateway

## Description
The API Gateway is responsible for validating authenticated users, saving user information, and acting as a gateway from the frontend. It also functions as a saga orchestrator between the Tracker Service and Vehicle Service via gRPC and displays live tracking using Redis streams.

## Requirements
- Docker
- Docker Compose
- Tracker Service up and running
- Vehicle Service up and running

## Getting Started
Follow these instructions to get the API Gateway up and running.

### Running the Service
1. Ensure you have Docker and Docker Compose installed on your machine.
2. Clone this repository to your local machine.
3. Navigate to the project directory.
4. Make sure the Tracker Service and Vehicle Service are up and running.
5. Run the shell script to start the API Gateway services:
    ```sh
    ./run.sh
    ```

### Testing
To enable testing with `docker-compose`:
1. Ensure you have a `docker-compose.test.yml` file set up for testing.
2. Run the following command to start the test setup:
    ```sh
    docker-compose -f docker-compose.test.yml up --build
    ```
3. Verify the services are running correctly and execute your tests.

## Features
- User authentication and authorization validation.
- User information saving.
- Gateway for frontend communication.
- Saga orchestration between Tracker Service and Vehicle Service via gRPC.
- Displays live tracking data using Redis streams.
- Services do not expose ports, except the API Gateway which is exposed publicly to be accessed on the dashboard.
- Other services can connect to them via Docker's private network.

## Postman Collection
To interact with the API Gateway using Postman, import the provided Postman collection:
1. The Postman collection file is included in the service repository as `we_plus_apigw.postman_collection.json`.
2. Open Postman and navigate to the "Import" button.
3. Select the `we_plus_apigw.postman_collection.json` file from the repository.
4. You can now use the imported collection to test the API endpoints.

## Usage
Once the services are running, you can interact with the API Gateway for user management and live tracking display. The API Gateway also orchestrates the communication between the Tracker and Vehicle services for seamless data handling.
