# Supermarket Checkout
A microservice for fetching the price of checkout items.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/jamieClayton1/supermarket-checkout.git
cd supermarket-checkout
```

2. Install dependencies:

```bash
go mod tidy
```
4. Build the project:

```bash
go build -o supermarket-checkout ./cmd/main

```

## Usage

To run the application locally, use the following commands:

```bash
go run ./cmd/api
```

To run the built binary:

```bash
./supermarket-checkout
```

To test the endpoint: 


Start scanning:
```
curl --location 'http://localhost:80/checkout/scan' \
--header 'Content-Type: application/json' \
--data '{
    "sku": "A",
}'
```

Use the response ``basket_id`` to add more items to the basket:

```
curl --location 'http://localhost:80/checkout/scan' \
--header 'Content-Type: application/json' \
--data '{
    "sku": "B",
    "basket_id":"63728163-1166-42fc-842f-3ec8f4053245"
}'
```

To check the price of your basket, provide the ``basket_id``:

```
curl --location 'http://localhost:80/checkout/63728163-1166-42fc-842f-3ec8f4053245/price' \
--header 'Content-Type: application/json' 
```