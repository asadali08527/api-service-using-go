
# API Service with JWT Authentication and Authorization

This repository contains a simple Go API service demonstrating user authentication, authorization, and role-based access control (RBAC) using JWT (JSON Web Tokens). The application allows admins to manage users and provides secure access to resources through role-based permissions.

## Table of Contents

- [Project Structure](#project-structure)
- [Environment Setup](#environment-setup)
- [Database Configuration](#database-configuration)
- [JWT Authentication and Authorization](#jwt-authentication-and-authorization)
- [API Endpoints](#api-endpoints)
- [Packages Documentation](#packages-documentation)
  - [config.go](#configgo)
  - [main.go](#maingo)
  - [controllers/admin_controller.go](#controllersadmin_controllergo)
  - [controllers/user_controller.go](#controllersuser_controllergo)
  - [middleware/jwt_middleware.go](#middlewarejwt_middlewarego)
  - [middleware/role_middleware.go](#middlewarerole_middlewarego)
  - [services/admin_service.go](#servicesadmin_servicego)
  - [services/user_service.go](#servicesuser_servicego)
  - [utils/jwt_utils.go](#utilsjwt_utilsgo)
  - [models/user.go](#modelsusergo)
  - [db/db.go](#dbdbgo)
- [Postman API Demo](#postman-api-demo)
- [Security Considerations](#security-considerations)

---

## Project Structure

```
/api-service
|-- config/
|   |-- config.go
|-- controllers/
|   |-- admin_controller.go
|   |-- user_controller.go
|-- middleware/
|   |-- jwt_middleware.go
|   |-- role_middleware.go
|-- services/
|   |-- admin_service.go
|   |-- user_service.go
|-- models/
|   |-- user.go
|-- db/
|   |-- db.go
|-- utils/
|   |-- jwt_utils.go
|-- main.go
|-- README.md
```

---

## Environment Setup

1. **Clone the repository**:
    ```bash
    git clone https://github.com/your-repo/api-service.git
    cd api-service
    ```

2. **Install dependencies**:
    Make sure you have Go installed and run:
    ```bash
    go mod tidy
    ```

3. **Set environment variables**:
    Update the environment variables in your `.env` file or export them in your terminal session:
    ```bash
    export JWT_SECRET="your_secret_key"
    export DB_URL="your_postgres_connection_string"
    ```

4. **Run the application**:
    ```bash
    go run main.go
    ```

---

## Database Configuration

The application uses PostgreSQL as the database. Make sure to configure the `DB_URL` correctly and have the necessary permissions to create and modify tables. The `db.go` file automatically migrates the `User` model on application startup.

---

## JWT Authentication and Authorization

The application implements JWT-based authentication to verify users and provides role-based access to resources. Admins can manage users (CRUD operations), and authenticated users can view or update their profiles.

---

## API Endpoints

| Method | Endpoint                | Description                                          | Access     |
|--------|-------------------------|------------------------------------------------------|------------|
| POST   | `/register`              | Register a new user                                  | Public     |
| POST   | `/register/admin`        | Register a new admin                                 | Public     |
| POST   | `/login`                 | Log in as a user or admin and receive JWT token       | Public     |
| GET    | `/api/profile`           | Get the authenticated user's profile                 | User/Admin |
| PUT    | `/api/profile`           | Update the authenticated user's profile              | User/Admin |
| GET    | `/api/admin/users`       | Get all users (Admin only)                           | Admin      |
| POST   | `/api/admin/users`       | Create a new user (Admin only)                       | Admin      |
| DELETE | `/api/admin/users/{id}`  | Delete a user by ID (Admin only)                     | Admin      |
| POST   | `/api/admin/users/{id}/revoke` | Revoke a user's token (Admin only)             | Admin      |

---

## Packages Documentation

### config.go

- **Purpose**: Stores environment variables such as JWT secret and database connection URL. Environment variables should be stored securely.
  
### main.go

- **Purpose**: Main entry point for the application. It initializes the database, services, controllers, and sets up the routes and middlewares.
  
### controllers/admin_controller.go

- **Purpose**: Admin-related operations like creating users, getting all users, deleting users, and revoking tokens. Admin-only routes are protected by `AdminRoleMiddleware`.

### controllers/user_controller.go

- **Purpose**: Handles user-related operations such as registering, logging in, viewing, and updating user profiles. JWT is used to authenticate and authorize requests.

- **Key Endpoints**:
  - `/register`: Registers a new user with the default "user" role.
  - `/login`: Logs in users and returns a JWT.
  - `/api/profile`: Allows users to view and update their profiles.

### middleware/jwt_middleware.go

- **Purpose**: Middleware that ensures the incoming request contains a valid JWT token in the `Authorization` header. If the token is valid, the user information is stored in the request context.

### middleware/role_middleware.go

- **Purpose**: Middleware that ensures the logged-in user has the "admin" role. It protects admin-only routes.

### services/admin_service.go

- **Purpose**: Provides business logic for admin operations like creating users, retrieving all users, deleting users, and revoking JWT tokens.

### services/user_service.go

- **Purpose**: Provides business logic for user operations such as registering, authenticating users, and managing user profiles.

### utils/jwt_utils.go

- **Purpose**: Core utility for JWT operations such as generating and validating tokens, and extracting user information from the token.

  - **GenerateJWT**: Generates a JWT token containing user-specific claims (username, role, etc.).
  - **ValidateToken**: Validates the JWT and extracts user claims (email, role, etc.).

### models/user.go

- **Purpose**: Contains the `User` model, including fields like ID, Username, Password, Role, and JWT token. The `User` model is mapped to the database table using GORM.

### db/db.go

- **Purpose**: Initializes the PostgreSQL database connection and automates the migration of the `User` model to create the necessary table on application startup.

---

## Postman API Demo

You can test the API using Postman or any HTTP client.

1. **Register a User**:
    - Endpoint: `POST /register`
    - Body:
      ```json
      {
        "username": "testuser",
        "password": "password123",
        "email": "test@example.com"
      }
      ```

2. **Login to Get JWT**:
    - Endpoint: `POST /login`
    - Body:
      ```json
      {
        "username": "testuser",
        "password": "password123"
      }
      ```

3. **Access Protected Routes**:
    - Use the token from login in the Authorization header:
      ```
      Authorization: Bearer <your_token_here>
      ```

4. **Create User as Admin**:
    - Login as an admin and use the token in the request to access admin-only routes like creating users or revoking tokens.

---

## Security Considerations

- **JWT Secret Management**: Store the JWT secret (`JWT_SECRET`) in environment variables or a secret management tool.
- **Token Expiry**: Tokens are set to expire after 24 hours. Ensure you implement a refresh token mechanism if needed.
- **Database Credentials**: Avoid hardcoding database credentials in code. Use environment variables for sensitive information.

---

### Example .env file:

```env
JWT_SECRET=your_secret_key
DB_URL=your_postgres_connection_string
```

---

### License

This project is licensed under the MIT License.

---

