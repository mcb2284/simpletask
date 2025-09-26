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
