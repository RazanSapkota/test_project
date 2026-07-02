# Basic Gin API (MVC Structure)

This project contains a basic Gin API scaffold using an MVC-style layout with:

- Route layer
- Controller layer

Service layer can be added later as the project grows.

## Project Structure

```text
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── controllers/
│   │   └── test_controller.go
│   └── routes/
│       └── routes.go
├── go.mod
├── go.sum
└── README.md
```

## Run

```bash
go run ./cmd/api
```

Server starts on `http://localhost:8080`

## Test Route

```http
GET /test
```

Expected response:

```json
{
	"message": "Gin API is working"
}
```

