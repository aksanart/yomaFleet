# Project Overview

## Description
This project contains four distinct services: API Gateway (apigw), Dashboard, Tracker, and Vehicle services. Each service plays a critical role in the overall system.

## Services
- **API Gateway (apigw)**: Acts as the main entry point, handles user authentication and authorization, and orchestrates communication between the Tracker and Vehicle services.
- **Dashboard**: Provides a user interface for visualizing and interacting with data from the Tracker and Vehicle services. Displays live tracking updates.
- **Tracker**: Provides data CRUD operations to track vehicles and sends data to a Redis stream for live tracking purposes.
- **Vehicle**: Manages vehicle-related data and operations.

## Running the Services
Follow these steps to run the services one by one:

1. **API Gateway (apigw)**:
    ```sh
    cd apigw
    ./run.sh
    ```

2. **Tracker**:
    ```sh
    cd tracker
    ./run.sh
    ```

3. **Vehicle**:
    ```sh
    cd vehicle
    ./run.sh
    ```

4. **Dashboard**:
    ```sh
    cd dashboard
    go run main.go
    ```

## Reviewing Service Details
For detailed information about each service, navigate to the respective folder and read the `README.md` file:
1. **API Gateway**:
    ```sh
    cd apigw
    cat README.md
    ```

2. **Tracker**:
    ```sh
    cd tracker
    cat README.md
    ```

3. **Vehicle**:
    ```sh
    cd vehicle
    cat README.md
    ```

4. **Dashboard**:
    ```sh
    cd dashboard
    cat README.md
    ```

Ensure each service is running successfully before starting the next one.