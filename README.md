# Project 2.2 — Containerizing Microservices and Publishing Images to Registries

## 1. Problem Statement

Containerize a multi-service application consisting of a frontend, backend, and database. Build independent Docker images for each service, tag the images appropriately, publish them to Docker Hub and Amazon Elastic Container Registry (AWS ECR), verify their availability, and deploy the application using Docker Compose.

The project also demonstrates environment-based backend configuration and a version 2 backend image.

## 2. Solution Approach

The application is organized into three independent services:

| Service | Technology | Container Port | Purpose |
|---|---|---:|---|
| Frontend | Nginx | 80 | Serves the application UI |
| Backend | Go | 8080 | Provides application and health endpoints |
| Database | PostgreSQL 16 Alpine | 5432 | Provides the database service |

Each service has its own Dockerfile and Docker image. Docker Compose is used to run the services together, while Docker secrets are used for the database password instead of placing the password in the Dockerfile.

The images are published to:
- Docker Hub
- Amazon ECR in the `us-east-1` region

## 3. Dependencies and Prerequisites

Install and configure:

- Git
- Docker Engine
- Docker Compose
- Docker Hub account
- AWS CLI
- AWS account with permission to create and push to ECR repositories
- AWS credentials configured for the CLI

Verify the tools:

```bash
git --version
docker --version
docker compose version
aws --version
```

## 4. Clone the Repository

```bash
git clone https://github.com/SnehalShinde11/docker-microservices-app.git
cd docker-microservices-app
```

## Task 1 — Prepare Microservices Application

### 1.1 Frontend Service

Location:

```
frontend/
├── Dockerfile
└── index.html
```

The frontend uses Nginx to serve the HTML application.

The page identifies the application as **Snehal Microservices Application** and displays the frontend, backend, and database services.

### 1.2 Backend Service

Location:

```
backend/
├── Dockerfile
├── Dockerfile.env
├── go.mod
├── main.go
└── main-env.go
```

The Go backend provides:

```
GET /
GET /health
```

The environment-aware implementation in `main-env.go` supports variables such as:

- `APP_NAME`
- `APP_ENV`
- `PORT`
- `DB_HOST`
- `DB_PORT`
- `DB_NAME`
- `DB_USER`

### 1.3 Database Service

Location:

```
database/
└── Dockerfile
```

The database image is based on:

```
postgres:16-alpine
```

The database password is supplied through a Docker secret and is not stored in the Dockerfile.

## Task 2 — Create Dockerfiles for Each Service

### 2.1 Frontend Dockerfile

The frontend Dockerfile uses Nginx and copies `index.html` into the Nginx web root.

Build context:

```
./frontend
```

### 2.2 Backend Dockerfile

The backend Dockerfile packages the Go application into a Docker image.

Build context:

```
./backend
```

The project also contains `Dockerfile.env` for the environment-aware backend version.

### 2.3 Database Dockerfile

The database Dockerfile uses PostgreSQL 16 Alpine and exposes port 5432.

## Task 3 — Build Docker Images

From the project root:

```bash
cd ~/microservices-app

docker build -t microservices-frontend:1.0 ./frontend
docker build -t microservices-backend:1.0 ./backend
docker build -t microservices-database:1.0 ./database
```

Verify the images:

```bash
docker images | grep microservices
```

## Task 4 — Tag Docker Images

Docker images are tagged with the required registry and version information.

### 4.1 General Docker Registry Tag Syntax

```text
docker tag <local-image>:<tag> <registry>/<repository>:<tag>
```

For Docker Hub:

```bash
docker tag microservices-frontend:1.0 snehalshinde11/microservices-frontend:1.0
docker tag microservices-backend:1.0 snehalshinde11/microservices-backend:1.0
docker tag microservices-database:1.0 snehalshinde11/microservices-database:1.0
```

For AWS ECR:

```bash
docker tag microservices-frontend:1.0 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-frontend:1.0
docker tag microservices-backend:1.0 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-backend:1.0
docker tag microservices-database:1.0 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-database:1.0
```

# Task 5 — Publish, Verify and Deploy Container Images

## 5.1 Push Images to Docker Hub

Login to Docker Hub:

```bash
docker login
```

Tag the images:

```bash
docker tag microservices-frontend:1.0 snehalshinde11/microservices-frontend:1.0
docker tag microservices-backend:1.0 snehalshinde11/microservices-backend:1.0
docker tag microservices-database:1.0 snehalshinde11/microservices-database:1.0
```

Push the images:

```bash
docker push snehalshinde11/microservices-frontend:1.0
docker push snehalshinde11/microservices-backend:1.0
docker push snehalshinde11/microservices-database:1.0
```

## 5.2 Push Images to AWS ECR

Create the ECR repositories:

```bash
aws ecr create-repository --repository-name microservices-frontend --region us-east-1
aws ecr create-repository --repository-name microservices-backend --region us-east-1
aws ecr create-repository --repository-name microservices-database --region us-east-1
```

Authenticate Docker with ECR:

```bash
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 307555122680.dkr.ecr.us-east-1.amazonaws.com
```

Tag the images:

```bash
docker tag microservices-frontend:1.0 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-frontend:1.0
docker tag microservices-backend:1.0 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-backend:1.0
docker tag microservices-database:1.0 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-database:1.0
```

Push the images:

```bash
docker push 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-frontend:1.0
docker push 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-backend:1.0
docker push 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-database:1.0
```

## 5.3 Verify Image Availability

### Frontend

Run the ECR image:

```bash
docker run -d \
  --name test-frontend \
  -p 8081:80 \
  307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-frontend:1.0

docker ps
curl http://localhost:8081
```

The response should contain the **Snehal Microservices Application** frontend content.

### Backend

Run the ECR image:

```bash
docker run -d \
  --name test-backend \
  -p 8080:8080 \
  307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-backend:1.0

curl http://localhost:8080
```

Verify the health endpoint:

```bash
curl http://localhost:8080/health
```

Expected health response:

```
Backend Service is healthy
```

### Database

Run the ECR database image with the database password supplied through a mounted secret file:

```bash
docker run -d \
  --name test-database \
  -p 5432:5432 \
  --env POSTGRES_DB=microservices \
  --env POSTGRES_USER=snehal \
  --env POSTGRES_PASSWORD_FILE=/run/secrets/db_password \
  --mount type=bind,source="$(pwd)/secrets/db_password.txt",target=/run/secrets/db_password,readonly \
  307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-database:1.0

docker ps
docker logs test-database
```

The logs should show:

```
database system is ready to accept connections
```

## 5.4 Deploy Using Docker Compose

The primary Compose configuration is `docker-compose.yml`.

It runs:

- Frontend on host port `8081`
- Backend on host port `8080`
- PostgreSQL on host port `5432`

Validate the Compose configuration:

```bash
cd ~/microservices-app
docker compose config
```

Remove standalone verification containers if they are still running:

```bash
docker rm -f test-frontend test-backend test-database
```

Start the application:

```bash
docker compose up -d
```

Check the services:

```bash
docker compose ps
```

Test the application:

```bash
curl http://localhost:8081
curl http://localhost:8080
curl http://localhost:8080/health
docker compose logs database
```

## 5.5 Configure Environment Variables

The backend service is configured through Docker Compose environment variables:

```text
APP_ENV=production
PORT=8080
DB_HOST=database
DB_PORT=5432
DB_NAME=microservices
DB_USER=snehal
```

Verify the variables inside the running backend container:

```bash
docker compose exec backend env | grep -E 'APP_ENV|PORT|DB_HOST|DB_PORT|DB_NAME|DB_USER'
```

The database password is handled separately through the Docker secret:

```text
/run/secrets/db_password
```

This avoids placing the password directly in the Dockerfile.

## 5.6 Build and Test Backend Version 2

Build the environment-aware backend image:

```bash
docker build \
  -f backend/Dockerfile.env \
  -t microservices-backend:v2 \
  ./backend
```

Run a standalone V2 container:

```bash
docker run -d \
  --name backend-v2-test \
  -p 8082:8080 \
  -e APP_NAME="Snehal Microservices Backend V2" \
  -e APP_ENV="production" \
  -e PORT="8080" \
  -e DB_HOST="database" \
  -e DB_PORT="5432" \
  -e DB_NAME="microservices" \
  -e DB_USER="snehal" \
  microservices-backend:v2
```

Test the V2 backend:

```bash
curl http://localhost:8082
curl http://localhost:8082/health
```

The application name should reflect the value supplied through `APP_NAME`.

## 5.7 Deploy Backend V2 Using Docker Compose

The V2 Compose configuration is `docker-compose-v2.yml`.

It uses:

| Service | Image | Host Port |
|---|---|---:|
| Frontend | `microservices-frontend:1.0` | 8083 |
| Backend | `microservices-backend:v2` | 8084 |
| Database | `microservices-database:1.0` | 5434 |

Validate and start the V2 stack:

```bash
docker compose -f docker-compose-v2.yml config
docker compose -f docker-compose-v2.yml up -d
```

Check the services:

```bash
docker compose -f docker-compose-v2.yml ps
```

Test the V2 backend:

```bash
curl http://localhost:8084
curl http://localhost:8084/health
```

The V2 Compose configuration sets:

```text
APP_NAME=Snehal Microservices Backend V2
APP_ENV=production
PORT=8080
DB_HOST=database
DB_PORT=5432
DB_NAME=microservices
DB_USER=snehal
```

## 5.8 Publish Backend V2 to Docker Hub

Tag the V2 backend image:

```bash
docker tag microservices-backend:v2 \
  snehalshinde11/microservices-backend:v2
```

Push it to Docker Hub:

```bash
docker push snehalshinde11/microservices-backend:v2
```

Verify the local tags:

```bash
docker images | grep microservices-backend
```

## Project Details

| Detail | Information |
|---|---|
| Name | Snehal Shinde |
| Project | Project 2.2 |
| Assignment | Containerizing Microservices and Publishing Images to Registries |
| GitHub Repository | `SnehalShinde11/docker-microservices-app` |
| Branch | `main` |
| Docker Hub Username | `snehalshinde11` |
| Docker Hub Frontend | `snehalshinde11/microservices-frontend:1.0` |
| Docker Hub Backend | `snehalshinde11/microservices-backend:1.0` |
| Docker Hub Database | `snehalshinde11/microservices-database:1.0` |
| Docker Hub Backend V2 | `snehalshinde11/microservices-backend:v2` |
| AWS Registry | `307555122680.dkr.ecr.us-east-1.amazonaws.com` |
| AWS Region | `us-east-1` |
| ECR Frontend Repository | `microservices-frontend` |
| ECR Backend Repository | `microservices-backend` |
| ECR Database Repository | `microservices-database` |
| Containerization | Docker |
| Orchestration | Docker Compose |
| Frontend | Nginx |
| Backend | Go |
| Database | PostgreSQL 16 Alpine |
