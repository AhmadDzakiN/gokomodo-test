# TOKO API
A Toko API by Ahmad Dzaki Naufal for gokomodo technical test

## For more detailed information, please see docs/Toko API Documentation.pdf file

## Tech Stack
- Go
- PostgreSQL
- Docker (Optional)

## How to Start
- Clone this repo
- Copy file `params/.env.sample` to `params/.env`
- Modify `params/.env` as your needs
- Ensure your PostgreSQL database already up
- You can use database/init.sql file to create all needed tables and seed the data
- Run the project by `go run main.go` or `go build && ./gokomodo-assignment`

## List of Comamnds
### API:
- `go run main.go` or `go build && ./gokomodo-assignment`: To run the API server

### Docker:
- `docker build -t {your-postgresdb-image-name} -f postgresdb.dockerfile .`: To build PostgreSQL docker image
- `docker container run -d --name {your-postgresdb-container-name} -p 5432:5432 {your-postgresdb-image-name}`: To run/up PostgreSQL docker container based on previous point config with 5432 exposed port
- `docker-compose up -d --build`: To immediately setup and run PostgreSQL & API docker image & container

This project uses Gin Framework and Clean Architecture

## Project structure
.
├── database
│   └── init.sql
├── documents
│   └── Gokomodo-Test.postman_collection.json
├── internal
│   ├── app
│   │   ├── config
│   │   │   ├── app.go
│   │   │   ├── gorm.go
│   │   │   ├── validator.go
│   │   │   ├── viper.go
│   │   │   └── zerolog.go
│   │   ├── constant
│   │   │   └── constant.go
│   │   ├── delivery
│   │   │   ├── http
│   │   │   │   ├── handler
│   │   │   │   │   ├── buyer.go
│   │   │   │   │   └── seller.go
│   │   │   │   └── route
│   │   │   │       └── router.go
│   │   │   └── middleware
│   │   │       └── auth.go
│   │   ├── entity
│   │   │   ├── buyer.go
│   │   │   ├── order.go
│   │   │   ├── product.go
│   │   │   └── seller.go
│   │   ├── mocks
│   │   │   └── repository
│   │   │       ├── buyer.go
│   │   │       ├── order.go
│   │   │       ├── product.go
│   │   │       └── seller.go
│   │   ├── payloads
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── repository
│   │   │   ├── buyer_test.go
│   │   │   ├── buyer.go
│   │   │   ├── order_test.go
│   │   │   ├── order.go
│   │   │   ├── product_test.go
│   │   │   ├── product.go
│   │   │   ├── seller_test.go
│   │   │   └── seller.go
│   │   └── service
│   │       ├── buyer_test.go
│   │       ├── buyer.go
│   │       ├── seller_test.go
│   │       └── seller.go
│   └── pkg
│       ├── jwt
│       │   └── jwt.go
│       └── pagination
│           └── pagination.go
├── params
│   ├── .env
│   └── .env.sample
├── .gitignore
├── app.dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── postgresdb.dockerfile
└── README.md