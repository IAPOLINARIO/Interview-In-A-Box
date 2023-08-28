# BrewInnovate Hiring Project

Welcome to my project submission for the BrewInnovate hiring process! This repository contains the code for a RESTful API designed to manage books, implemented as part of the technical assessment.

Built with Go and PostgreSQL, the application offers CRUD operations for books and demonstrates my abilities in software development, containerization, and DevOps. The code is well-organized and easily expandable, showcasing the best practices and design patterns I adhere to.

Below you'll find instructions on how to set up the project, details about the endpoints, Docker Compose configuration, and sample requests using `curl`.

I'm excited to share this project with the team at BrewInnovate, and I look forward to the opportunity to discuss it further. Enjoy exploring the code, and feel free to reach out with any questions!

## Table of Contents

- [Configuration](#configuration)
- [Docker](#docker)
- [Endpoints](#endpoints)
- [Examples](#examples)

## Configuration

Make sure you have Go installed and a PostgreSQL server running.

Set the following environment variables to connect to the PostgreSQL server:

- `POSTGRES_USER`: Database user
- `POSTGRES_PASSWORD`: Database password
- `POSTGRES_HOST`: Database host
- `POSTGRES_PORT`: Database port
- `POSTGRES_DB`: Database

## Docker

The application can also be run using Docker Compose. The provided `docker-compose.yml` file configures two services: the application itself (`app`) and a PostgreSQL database (`db`).

### Services

#### `app`

- **Build**: Context is set to the current directory and uses the `Dockerfile`.
- **Ports**: Exposed on port `8090`.
- **Dependencies**: Depends on the `db` service.
- **Environment**: Uses the `.env` file for environment variables.

#### `db`

- **Image**: Uses the `postgres:15.3` image.
- **Restart**: Always restarts if it stops.
- **Volumes**: Mounts the PostgreSQL data to `postgres_data` and the `startup.sh` script to the initialization directory.
- **Environment**: Uses the `.env` file for environment variables.

### Volumes

- `postgres_data`: Volume used to persist PostgreSQL data.

### Running with Docker Compose

To run the application using Docker Compose, first make sure you have Docker and Docker Compose installed. Then run the following command:

```
docker-compose up --build
```

The application will be available at `localhost:8090`.

## Endpoints

- `GET /`: Health check
- `GET /books`: Get all books
- `GET /books/{id}`: Get a book by ID
- `POST /books`: Create a new book
- `PUT /books/{id}`: Update a book by ID
- `DELETE /books/{id}`: Delete a book by ID

## Examples

Below are some sample requests using `curl` for each endpoint:

### Health Check

```
curl http://localhost:8090/
```

### Get All Books

```
curl http://localhost:8090/books
```

### Get a Book by ID

Replace `1` with the desired ID.

```
curl http://localhost:8090/books/1
```

### Create a New Book

```
curl -X POST -H "Content-Type: application/json" -d '{"title":"Book Title", "author":"Author Name"}' http://localhost:8090/books
```

### Update a Book by ID

Replace `1` with the desired ID.

```
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Title", "author":"Updated Author"}' http://localhost:8090/books/1
```

### Delete a Book by ID

Replace `1` with the desired ID.

```
curl -X DELETE http://localhost:8090/books/1
```

### Note

Make sure to configure the `.env` file with the required environment variables as mentioned in the [Configuration](#configuration) section.
