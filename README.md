# SimpleTask API

A simple backend service for managing tasks and users built with Go, Gin, and GORM.

## Features

- CRUD operations for tasks and users
- Task filtering by status and user
- Pagination and sorting
- Idempotency support
- Task summary statistics

## Setup

1. Install Go (1.19+)
2. Clone the repository
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Run

```bash
go run .
```

The server will start on `http://localhost:8080`

## Testing with Postman

A Postman collection is included (`simpletask.postman_collection.json`) with pre-configured requests for all API endpoints. To use it:

1. Import the collection file into Postman
2. Set the base URL to `http://localhost:8080` (or update as needed)
3. Use the requests to test the API endpoints

## API Endpoints

### Users
- `POST /user` - Create user
- `GET /user/:userid` - Get user
- `PATCH /user/:userid` - Update user
- `DELETE /user/:userid` - Delete user

### Tasks
- `POST /tasks` - Create task
- `GET /tasks` - List tasks (supports filtering, pagination, sorting)
- `GET /tasks/:taskid` - Get task
- `PATCH /tasks/:taskid` - Update task
- `DELETE /tasks/:taskid` - Delete task
- `GET /tasks/summary` - Get task count summary

### Query Parameters for GET /tasks
- `status` - Filter by status (pending, in_progress, done)
- `id` - Filter by user ID
- `order` - Sort by due_date (asc, desc)
- `limit` - Number of results (default: 10)
- `offset` - Results offset (default: 0)
