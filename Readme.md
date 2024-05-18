# Blogging Platform

## Setup

### Prerequisites

* Install Go (https://go.dev/doc/install)
* Install SQLite3 (https://www.sqlite.org/download.html)
* (Optional) Install Docker ([https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/))
* (Optional) Install Docker Compose ([https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/))

### Local Setup

1. Clone the repository:

```sh
git clone [https://github.com/blog-platform.git](https://github.com/blog-platform.git)
cd blog-platform
```

2. Navigate to the project directory:

```sh
cd blog-platform
```

3. Install dependencies:

```sh
go mod tidy
```

4. Set up the environment variable `JWT_SECRET`:

```sh
export JWT_SECRET="your_jwt_secret"
```

5. Set up the SQLite database:

```sh
sqlite3 ./database/blog.db < ./database/migrations.sql
```

6. Run the application:

```sh
go run main.go
```

### Using Docker Compose

1. Build and start the application with Docker Compose:

```sh
docker-compose up --build
```

2. The application will be available at `http://localhost:3000`.

3. To stop the application:

```sh
docker-compose down
```

### API Endpoints

#### Authentication

* `POST /login`: Login and receive a JWT token.

#### Posts

* `POST /posts`: Create a new post.
* `GET /posts`: Retrieve a list of posts with optional filtering by author and creation date, and pagination.
* `GET /posts/:id`: Retrieve a single post by ID.
* `PUT /posts/:id`: Update an existing post.
* `DELETE /posts/:id`: Delete a post by ID.

#### Users

* `POST /users`: Create a new user.
* `PATCH /users/:id/role`: Update a user's role.
* `DELETE /users/:id`: Delete a user by ID.

### Generating Swagger Documentation

1. Install `swag` if you haven't already:

```sh
go install github.com/swaggo/swag/cmd/swag@latest [invalid URL removed]
```

2. Generate the Swagger documentation:

```sh
swag init
```

3. The generated documentation will be available at `http://localhost:3000/swagger/index.html` when the application is running.

### Running Tests

1. Navigate to the project directory:

```sh
cd /path/to/your/project
```

2. Run the tests:

```sh
go test ./...
```

3. For more detailed output, use the `-v` (verbose) flag:

```sh
go test -v ./...
```

## Docker Compose Configuration

**docker-compose.yml**

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "3000:3000"
    environment:
      - JWT_SECRET=secret
    volumes:
      - ./database:/app/database
    entrypoint: ["sh", "-c", "sqlite3 /app/database/blog.db < /app/database/migrations.sql && ./main"]
```

**Curl Examples**

Here are some curl examples to interact with the API endpoints:

**Authentication**

```sh
curl -X POST http://localhost:3000/login \
-H "Content-Type: application/json" \
-d '{"username":"your_username", "password":"your_password"}'
```

**Posts**

**Create a Post**

```sh
curl -X POST http://localhost:3000/posts \
-H "Authorization: Bearer your_jwt_token" \
-H "Content-Type: application/json" \
-d '{"title":"Your Post Title", "content":"Your Post Content", "author":"Author Name"}'
```

**Get All Posts**

```sh
curl -X GET http://localhost:3000/posts \
-H "Authorization: Bearer your_jwt_token"
```

**Get a Post by ID**

```sh
curl -X GET http://localhost:3000/posts/1 \
-H "Authorization: Bearer your_jwt_token"
```

**Update a Post**

```sh
curl -X PUT http://localhost:3000/posts/1 \
-H "Authorization: Bearer your_jwt_token" \
-H "Content-Type: application/json" \
-d '{"title":"Updated Post Title", "content":"Updated Post Content"}'
```

**Delete a Post**

```sh
curl -X DELETE http://localhost:3000/posts/1 \
-H "Authorization: Bearer your_jwt_token"
```

**Users**

* `POST /users`: Create a new user. You'll likely need admin privileges for this action.
* `PATCH /users/:id/role`: Update a user's role. You'll likely need admin privileges for this action.
* `DELETE /users/:id`: Delete a user. You'll likely need admin privileges for this action.

**Notes:**

* Replace `your_jwt_token` with the actual JWT token you receive after successful login.
* Replace `your_username` and `your_password` with your actual login credentials.
* Remember to adjust file paths (`./database/blog.db` and `./database/migrations.sql`) if they differ in your project setup.
