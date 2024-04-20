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

```
curl --location 'http://localhost:80/checkout/price' \
--header 'Content-Type: application/json' \
--data '{
    "item_skus": ["A", "B", "B", "C", "D"]
}'
```