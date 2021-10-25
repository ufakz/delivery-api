# Delivery-API
This is a simulated delivery API for the backend of a webstore that uses the [Gokada.ng business API](https://www.gokada.ng/business) to generate estimates for orders and store order details in a [Firebase Firestore](https://firebase.google.com/products/firestore) cloud database.

## Installation
Clone the repository and install dependencies from the root directory using
```
go install
```
## API Documentation
The API has two endpoints:
- `/delivery_estimate` which returns the estimate of a delivery order based on the distance between the source and destination addresses
- `/delivery_order` which stores details of the order in the firestore database and returns the estimate of the delivery order

The API documentation is automatically generated using [swagger](https://swagger.io). The documentation files are [here](./docs). The automatically generated API spec file is [here](./delivery.yaml).

To generate the API spec yaml file, run the following command from the root directory:
```
swagger generate spec -o ./delivery.yaml --scan-models
```

To serve the documentation on a browser interface. Run the command:
```
swagger serve -F=swagger swagger.yaml
```

## Testing
The API uses the standard go `testing` library. Currently only tests for the delivery order [exists](./delivery/main_test.go). To run the test, simply run the `go test -v ./delivery` command.
