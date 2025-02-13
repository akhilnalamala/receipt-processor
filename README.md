# Fetch-Rewards Take Home Assessment Solution.

This is my solution for the Fetch-Rewards take home assessment, which is the Receipt Processor Challenge.

## Intro

The required web server has been implemented in Go. I have chosen not to Dockerize it as the assessment details mentioned that the evaluator will have Go installed on their system. 

## Dependencies

I have installed four external packages to help implement this:

- [chi](https://github.com/go-chi/chi) - to set up routes  
- [render](https://github.com/go-chi/render) - to manage HTTP requests and responses  
- [validator](https://github.com/go-playground/validator) - for validation of the `Item` and `Receipt` structs  
- [uuid](https://github.com/google/uuid) - to generate UUIDs for the receipts  

## Installation and Running the Server

1. Clone the repo
```
git clone https://github.com/akhilnalamala/receipt-processor.git
```

2. Navigate to directory
```
cd receipt-processor
```

3. Install dependencies
```
go mod tidy
```

4. Run the server
```
go run main.go
```

- This runs and serves the app on localhost:8080, meaning the 2 endpoints are available at 
```
http://localhost:8080/receipts/process - POST
http://localhost:8080/receipts/{id}/points - GET
```