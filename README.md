# Todo Application

This is a full-stack Todo application built with Go (Fiber) for the backend and React for the frontend.

## Features

- Create, read, update, and delete todo items
- Mark todos as completed
- MongoDB integration for data persistence
- RESTful API
- Cross-Origin Resource Sharing (CORS) enabled

## Tech Stack

### Backend
- Go
- Fiber (web framework)
- MongoDB (database)
- godotenv (for environment variable management)

### Frontend
- React
- Chakra UI (for styling)
- React Query (for state management and data fetching)

## Getting Started

### Prerequisites

- Go (1.16+)
- Node.js (14+)
- MongoDB

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/todo-application.git
   cd todo-application
   ```

2. Set up the backend:
   ```
   cd backend
   go mod tidy
   ```

3. Create a `.env` file in the backend directory with the following content:
   ```
   MONGODB_URI=your_mongodb_connection_string
   PORT=4000
   ```

4. Set up the frontend:
   ```
   cd client
   npm install
   ```

### Running the Application

1. Start the backend server:
   ```
   In the parent directory
   go run main.go
   ```

2. In a new terminal, start the frontend development server:
   ```
   cd client
   npm run dev
   ```

3. Open your browser and navigate to `http://localhost:5173` (or the port specified by Vite).

## API Endpoints

- `GET /api/todos`: Fetch all todos
- `POST /api/todos`: Create a new todo
- `PATCH /api/todos/:id`: Update a todo (mark as completed)
- `DELETE /api/todos/:id`: Delete a todo

## Deployment

For production deployment:

1. Build the frontend:
   ```
   cd client
   npm run build
   ```

2. Set the `ENV` environment variable to "production" on your server.

3. Run the Go application on your server, which will serve both the API and the static frontend files.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
