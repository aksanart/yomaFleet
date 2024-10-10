# Dashboard Service

## Description
The Dashboard Service provides a user interface for visualizing and interacting with the data from the Tracker and Vehicle services. It also handles live updates for tracking using Redis streams.

## Requirements
- Go installed on your machine
- API Gateway up and running
- Tracker Service up and running
- Vehicle Service up and running

## Getting Started
Follow these instructions to get the Dashboard Service up and running.

### Running the Service
1. Ensure you have Go installed on your machine.
2. Clone this repository to your local machine.
3. Navigate to the project directory.
4. Make sure the API Gateway, Tracker Service, and Vehicle Service are up and running.
5. Run the following command to start the Dashboard Service:
    ```sh
    go run main.go
    ```

### Features
- User interface for data visualization and interaction.
- Live tracking updates using Redis streams.
- Dashboard Service port is exposed publicly to be accessed on the dashboard.

## Usage
Once the services are running, you can access the Dashboard Service to visualize and interact with data from the Tracker and Vehicle services. The Dashboard Service also handles live updates for tracking using Redis streams.
