# Receipt Points Web Service

A Go-based web service that processes receipts and calculates points based on specific business rules. The service accepts receipt data in JSON format, calculates points based on predefined criteria, and returns a unique ID for each receipt.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Go](https://golang.org/doc/install) (optional, for local development)

## Project Structure

```
/fetch-take-home-exercise
├── main.go                 # Application entry point
├── handlers/              
│   └── receipt_handler.go  # HTTP request handlers
├── services/              
│   └── receipt_service.go  # Business logic layer
├── models/                
│   └── receipt.go         # Data models
└── utils/                 
    └── utils.go           # Helper functions
```

## Running the Application

### With Docker

1. Build the image:
   ```sh
   docker build -t receipt-processor .
   ```

2. Run the container:
   ```sh
   docker run -p 8080:8080 receipt-processor
   ```

The service will be available at `http://localhost:8080`.

## API Usage

### Process Receipt

```sh
curl -X POST http://localhost:8080/receipts/process \
  -H "Content-Type: application/json" \
  -d '{
    "retailer": "M&M Corner Market",
    "purchaseDate": "2022-03-20",
    "purchaseTime": "14:33",
    "items": [
      {
        "shortDescription": "Gatorade",
        "price": "2.25"
      },
      {
        "shortDescription": "Gatorade",
        "price": "2.25"
      },
      {
        "shortDescription": "Gatorade",
        "price": "2.25"
      },
      {
        "shortDescription": "Gatorade",
        "price": "2.25"
      }
    ],
    "total": "9.00"
  }'
```

### Get Points

```sh
curl http://localhost:8080/receipts/{id}/points
```

## Running Tests

Execute the test suite:

```sh
go test ./services
```