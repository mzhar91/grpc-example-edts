<!-- ABOUT THE PROJECT -->

## About The Project

EDTS sharing session about basic gRPC and golang framework

<!-- GETTING STARTED -->

## Getting Started

### Prerequisites (OSX)
- [x] Go installed | go1.20.2 darwin/amd64

### Installation
1. Clone the repo
   ```sh
   git clone https://github.com/mzhar91/grpc-example-edts.git
   ```
2. Install go packages in example-order-api folder
   ```sh
   cd example-order-api && go install && go mod vendor
   ```
3. Install go packages in example-payment-api folder
   ```sh
   cd example-payment-api && go install && go mod vendor
   ```

### Run Application
1. Run go order service
   ```sh
   cd example-order-api && go run main.go
   ```
2. Run go payment service
   ```sh
   cd example-payment-api && go run main.go
   ```

<!-- CONTACT -->

## Contact

Harry Kurniawan - harry.kurniawan@sg-dsa.com | k.harry791@gmail.com