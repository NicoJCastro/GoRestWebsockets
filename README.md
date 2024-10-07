# REST API and WebSockets with Go

## Description
This project is a REST API and WebSocket server built with Go. It provides endpoints for user authentication, post management, and real-time communication using WebSockets.

## Features
- User Signup and Login
- CRUD operations for posts
- Real-time communication with WebSockets
- Middleware for authentication

## Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/your-repo.git
    ```
2. Navigate to the project directory:
    ```sh
    cd your-repo
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage
1. Run the server:
    ```sh
    go run main.go
    ```
2. The server will start on `http://localhost:8080`.

## Endpoints
### Public Endpoints
- `GET /` - Home
- `POST /signup` - User Signup
- `POST /login` - User Login
- `GET /posts/{id}` - Get Post by ID
- `GET /posts` - List Posts

### Protected Endpoints (require authentication)
- `GET /api/v1/me` - Get User Info
- `POST /api/v1/posts` - Insert Post
- `PUT /api/v1/posts/{id}` - Update Post by ID
- `DELETE /api/v1/posts/{id}` - Delete Post by ID

## WebSocket
- `GET /ws` - WebSocket endpoint for real-time communication

## Middleware
- `CheckAuthMiddleware` - Middleware to check authentication for protected endpoints

## Contributing
1. Fork the repository.
2. Create a new branch:
    ```sh
    git checkout -b feature-branch
    ```
3. Make your changes and commit them:
    ```sh
    git commit -m "Description of changes"
    ```
4. Push to the branch:
    ```sh
    git push origin feature-branch
    ```
5. Open a pull request.

## License
This project is licensed under the MIT License.