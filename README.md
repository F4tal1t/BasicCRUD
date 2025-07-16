# BasicCRUD - Car Inventory API

![Go](https://img.shields.io/badge/Go-1.24-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-blue)
![Fiber](https://img.shields.io/badge/Fiber-Web%20Framework-green)

BasicCRUD is a Go-based REST API application that performs basic CRUD operations on a PostgreSQL database. It provides a simple and efficient API to manage a car inventory system.

## ✨ Features

- **Create** a new car entry
- **Retrieve** a car by ID
- **Update** car details
- **Delete** a car entry
- **Swagger Documentation** for API endpoints
- **Environment-based Configuration** for security
- **Concurrent Request Handling** with mutex protection
- **Comprehensive Error Handling**

## 🛠️ Requirements

- [Go](https://golang.org) 1.24 or later
- [PostgreSQL](https://www.postgresql.org) 12 or later

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone <repository-url>
cd BasicCRUD
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Set Up Environment Variables

Copy the `.env.example` to `.env` and update the values with your PostgreSQL credentials:

```bash
cp .env.example .env
```

Edit the `.env` file:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_NAME=mycarsdb
DB_SSL_MODE=disable
PORT=8080
```

### 4. Database Setup

1. **Create a new database**

   ```sql
   CREATE DATABASE mycarsdb;
   ```

2. **Run the included schema script**

   ```bash
   psql -U <username> -d mycarsdb -f schema.sql
   ```

### 5. Run the Application

```bash
go run main.go
```

The server will start at `http://localhost:8080` (or the port defined in your `.env` file).

## 📚 API Documentation

### Interactive Documentation
Visit `http://localhost:8080/swagger/index.html` in your browser to view the interactive API documentation.

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST   | `/cars`  | Create a new car |
| GET    | `/cars/{id}` | Get a car by ID |
| PUT    | `/cars/{id}` | Update a car |
| DELETE | `/cars/{id}` | Delete a car |

### Request/Response Examples

#### Create a Car
```bash
curl -X POST http://localhost:8080/cars \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Camaro",
    "model": "SS",
    "brand": "Chevrolet",
    "year": 2023,
    "price": 45000.00
  }'
```

#### Get a Car
```bash
curl http://localhost:8080/cars/1
```

#### Update a Car
```bash
curl -X PUT http://localhost:8080/cars/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Camaro",
    "model": "ZL1",
    "brand": "Chevrolet",
    "year": 2023,
    "price": 65000.00
  }'
```

#### Delete a Car
```bash
curl -X DELETE http://localhost:8080/cars/1
```

## 🧪 Testing

Run the included tests:

```bash
go test ./...
```

Run benchmarks:

```bash
go test -bench=. ./handlers
```

## 🏗️ Project Structure

```
BasicCRUD/
├── config/           # Database configuration
├── handlers/         # HTTP handlers
├── middleware/       # Custom middleware
├── models/           # Data models
├── utils/            # Utility functions
├── docs/             # Auto-generated Swagger docs
├── main.go           # Application entry point
├── schema.sql        # Database schema
├── .env.example      # Environment variables template
├── .gitignore        # Git ignore file
└── README.md         # This file
```

## 🔒 Security Features

- Environment-based configuration (no hardcoded credentials)
- SQL injection protection through parameterized queries
- Input validation and sanitization
- Secure headers middleware
- Concurrent request protection

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🐛 Issues

If you encounter any issues, please [create an issue](../../issues) on GitHub.

