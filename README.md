# Receipt Processor

A web service that processes receipts and awards points based on defined rules.

## Running the Application

### Prerequisites

- Docker must be installed on your system.

### Building and Running with Docker

1. Build the Docker image:

```bash
docker build -t receipt-processor .
```

2. Run the Docker container:

```bash
docker run -p 8080:8080 receipt-processor
```

3. The service will be accessible at http://localhost:8080.

## API Endpoints

### Process Receipt
- Path: /receipts/process
- Method: POST
- Description: Submits a receipt for processing and returns a unique ID.

### Get Points
- Path: /receipts/{id}/points
- Method: GET
- Description: Returns the number of points awarded to the receipt.

## Example Request and Response

### Process Receipt
```bash
curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
        {"shortDescription": "Mountain Dew 12PK", "price": "6.49"},
        {"shortDescription": "Emils Cheese Pizza", "price": "12.25"},
    ],
    "total": "18.74"
}'
```

### Get Points
```bash
curl http://localhost:8080/receipts/{id}/points
```
Replace {id} with the ID returned from the Process Receipt request.