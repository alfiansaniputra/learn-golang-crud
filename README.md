# Belajar Golang API

## Project Overview
This project is a simple API built with Go (Golang) to demonstrate clean code practices, professional structure, and usability. It includes features such as user authentication and data retrieval.

## Features
- **Login API**: Authenticate users using either email or username and password.
- **Check Data API**: Retrieve all user data from the database, handling nullable fields gracefully.

## Technologies Used
- **Go**: Programming language.
- **GORM**: ORM for database interactions.
- **MySQL**: Database.
- **Dotenv**: For managing environment variables.

## Project Structure
```
├── config
│   └── db.go          # Database connection setup
├── controllers
│   └── auth_controller.go # API handlers
├── models
│   └── user.go        # User model
├── routes
│   └── auth_routes.go # API routes
├── utils
│   ├── hash.go        # Password hashing utilities
│   ├── response.go    # Standardized API responses
│   └── token.go       # JWT token generation
├── .env               # Environment variables
├── main.go            # Application entry point
└── README.md          # Project documentation
```

## Setup Instructions
1. Clone the repository:
   ```bash
   git clone <repository-url>
   ```
2. Navigate to the project directory:
   ```bash
   cd 1-belajar-golang-api
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Set up the `.env` file with your database credentials:
   ```env
   DATABASE_MAIN_HOST=localhost
   DATABASE_MAIN_NAME=your_database_name
   DATABASE_MAIN_USER=your_username
   DATABASE_MAIN_PASSWORD=your_password
   DATABASE_MAIN_PORT=3306
   DATABASE_MAIN_DIALECT=mysql
   ```
5. Run the application:
   ```bash
   go run main.go
   ```

## API Endpoints
### Login
**POST** `/login`
- Request Body:
  ```json
  {
    "identifier": "email_or_username",
    "password": "your_password"
  }
  ```
- Response:
  ```json
  {
    "token": "jwt_token",
    "user": {
      "id": 1,
      "username": "example",
      "email": "example@example.com",
      "photo": "-",
      "fcm_token": "-",
      "created_at": "2025-06-22T00:00:00Z",
      "updated_at": "2025-06-22T00:00:00Z"
    }
  }
  ```

### Check Data
**GET** `/checkdata`
- Response:
  ```json
  [
    {
      "id": 1,
      "username": "example",
      "email": "example@example.com",
      "photo": "-",
      "fcm_token": "-",
      "created_at": "2025-06-22T00:00:00Z",
      "updated_at": "2025-06-22T00:00:00Z"
    }
  ]
  ```

## License
This project is licensed under the MIT License.
