# Project 2.2 — Containerizing Microservices and Publishing Images to Registries

## 1. Problem Statement

Containerize a multi-service application consisting of a frontend, backend, and database. Build independent Docker images for each service, tag the images appropriately, publish the images to Docker Hub and Amazon Elastic Container Registry (AWS ECR), verify their availability, and deploy the application using Docker Compose.

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

---

## 3. Dependencies and Prerequisites

Install and configure:

- Git
- Docker Engine
- Docker Compose
- Docker Hub account
- AWS CLI
- AWS account with permission to create and push to ECR repositories
- AWS credentials configured for the CLI

Verify the installed tools:

```bash
git --version
docker --version
docker compose version
aws --version
```

---

## 4. Clone the Repository

### General Syntax

```bash
git clone <repository-url>
cd <repository-directory>
```

### Project Implementation

```bash
git clone https://github.com/SnehalShinde11/docker-microservices-app.git
cd docker-microservices-app
```

---

# Task 1 — Prepare Microservices Application

## 1.1 Frontend Service

Location:

```text
frontend/
├── Dockerfile
└── index.html
```

The frontend uses Nginx to serve the HTML application.

The page identifies the application as **Snehal Microservices Application** and displays the frontend, backend, and database services.

### General Nginx Dockerfile Structure

```dockerfile
FROM nginx:alpine

COPY <source-file> /usr/share/nginx/html/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
```

The actual project Dockerfile is available in:

```text
frontend/Dockerfile
```

---

## 1.2 Backend Service

Location:

```text
backend/
├── Dockerfile
├── Dockerfile.env
├── go.mod
├── main.go
└── main-env.go
```

The Go backend provides:

```text
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

### General Go Build Syntax

```bash
go build -o <output-binary> <source-file>
```

The project also contains `Dockerfile.env` for the environment-aware backend version.

---

## 1.3 Database Service

Location:

```text
database/
└── Dockerfile
```

The database image is based on:

```text
postgres:16-alpine
```

The database password is supplied through a Docker secret and is not stored directly in the Dockerfile.

---

# Task 2 — Create Dockerfiles for Each Service

## 2.1 Frontend Dockerfile

The frontend Dockerfile uses Nginx and copies `index.html` into the Nginx web root.

Build context:

```text
./frontend
```

### General Docker Build Syntax

```bash
docker build -t <image-name>:<tag> <build-context>
```

### Project Implementation

```bash
docker build -t microservices-frontend:1.0 ./frontend
```

---

## 2.2 Backend Dockerfile

The backend Dockerfile packages the Go application into a Docker image.

Build context:

```text
./backend
```

### Project Implementation

```bash
docker build -t microservices-backend:1.0 ./backend
```

The project also contains `Dockerfile.env` for the environment-aware backend version.

To specify a custom Dockerfile:

```bash
docker build \
  -f <dockerfile> \
  -t <image-name>:<tag> \
  <build-context>
```

Project example:

```bash
docker build \
  -f backend/Dockerfile.env \
  -t microservices-backend:v2 \
  ./backend
```

---

## 2.3 Database Dockerfile

The database Dockerfile uses PostgreSQL 16 Alpine and exposes port `5432`.

### General Dockerfile Base Image Syntax

```dockerfile
FROM <base-image>:<tag>
```

### Project Implementation

```dockerfile
FROM postgres:16-alpine
```

---

# Task 3 — Build Docker Images

From the project root:

```bash
cd ~/microservices-app

docker build -t microservices-frontend:1.0 ./frontend
docker build -t microservices-backend:1.0 ./backend
docker build -t microservices-database:1.0 ./database
```

### Verify Images

General syntax:

```bash
docker images
```

Project verification:

```bash
docker images | grep microservices
```

---

# Task 4 — Tag Docker Images

Docker images are tagged with the required registry and version information.

### General Docker Registry Tag Syntax

```bash
docker tag <local-image>:<tag> <registry>/<repository>:<tag>
```

### Docker Hub Syntax

```bash
docker tag <local-image>:<tag> \
  <dockerhub-username>/<repository>:<tag>
```

### Project Implementation

```bash
docker tag microservices-frontend:1.0 snehalshinde11/microservices-frontend:1.0
docker tag microservices-backend:1.0 snehalshinde11/microservices-backend:1.0
docker tag microservices-database:1.0 snehalshinde11/microservices-database:1.0
```

### AWS ECR Syntax

```bash
docker tag <local-image>:<tag> \
  <aws-account-id>.dkr.ecr.<region>.amazonaws.com/<repository>:<tag>
```

### Project Implementation

```bash
docker tag microservices-frontend:1.0 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-frontend:1.0
docker tag microservices-backend:1.0 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-backend:1.0
docker tag microservices-database:1.0 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-database:1.0
```

---

# Task 5 — Publish, Verify and Deploy Container Images

## 5.1 Push Images to Docker Hub

### General Syntax

Login:

```bash
docker login
```

Push:

```bash
docker push <dockerhub-username>/<repository>:<tag>
```

### Project Implementation

Login to Docker Hub:

```bash
docker login
```

Push the images:

```bash
docker push snehalshinde11/microservices-frontend:1.0
docker push snehalshinde11/microservices-backend:1.0
docker push snehalshinde11/microservices-database:1.0
```

---

## 5.2 Push Images to AWS ECR

### General Syntax

Create an ECR repository:

```bash
aws ecr create-repository \
  --repository-name <repository-name> \
  --region <aws-region>
```

Authenticate Docker with ECR:

```bash
aws ecr get-login-password \
  --region <aws-region> | \
  docker login \
  --username AWS \
  --password-stdin <aws-account-id>.dkr.ecr.<aws-region>.amazonaws.com
```

Push an image:

```bash
docker push \
  <aws-account-id>.dkr.ecr.<aws-region>.amazonaws.com/<repository>:<tag>
```

### Project Implementation

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

Push the images:

```bash
docker push 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-frontend:1.0
docker push 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-backend:1.0
docker push 307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-database:1.0
```

---

## 5.3 Verify Image Availability

### General Docker Run Syntax

```bash
docker run -d \
  --name <container-name> \
  -p <host-port>:<container-port> \
  <image>:<tag>
```

### Frontend

Run the ECR image:

```bash
docker run -d \
  --name test-frontend \
  -p 8081:80 \
  307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-frontend:1.0
```

Verify the running container:

```bash
docker ps
```

Test the application:

```bash
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
```

Test the backend:

```bash
curl http://localhost:8080
```

Verify the health endpoint:

```bash
curl http://localhost:8080/health
```

Expected health response:

```text
Backend Service is healthy
```

### Database

The database container uses environment variables and a mounted secret file.

General environment variable syntax:

```bash
docker run \
  -e <VARIABLE>=<VALUE> \
  <image>:<tag>
```

General bind mount syntax:

```bash
docker run \
  --mount type=bind,source=<host-path>,target=<container-path>,readonly \
  <image>:<tag>
```

Project implementation:

```bash
docker run -d \
  --name test-database \
  -p 5432:5432 \
  --env POSTGRES_DB=microservices \
  --env POSTGRES_USER=snehal \
  --env POSTGRES_PASSWORD_FILE=/run/secrets/db_password \
  --mount type=bind,source="$(pwd)/secrets/db_password.txt",target=/run/secrets/db_password,readonly \
  307555122680.dkr.ecr.us-east-1.amazonaws.com/microservices-database:1.0
```

Verify the container:

```bash
docker ps
```

Check the database logs:

```bash
docker logs test-database
```

The logs should show:

```text
database system is ready to accept connections
```

---

## 5.4 Deploy Using Docker Compose

The primary Compose configuration is:

```text
docker-compose.yml
```

It runs:

- Frontend on host port `8081`
- Backend on host port `8080`
- PostgreSQL on host port `5432`

### General Docker Compose Syntax

Validate a Compose file:

```bash
docker compose config
```

Start services:

```bash
docker compose up -d
```

Check service status:

```bash
docker compose ps
```

View logs:

```bash
docker compose logs <service-name>
```

Stop services:

```bash
docker compose down
```

### Project Implementation

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

---

# 5.5 Configure Environment Variables

The backend service is configured through Docker Compose environment variables.

### General Docker Run Environment Variable Syntax

```bash
docker run \
  -e <VARIABLE1>=<VALUE1> \
  -e <VARIABLE2>=<VALUE2> \
  <image>:<tag>
```

### General Docker Compose Environment Syntax

```yaml
services:
  backend:
    environment:
      VARIABLE1: value1
      VARIABLE2: value2
```

### Project Implementation

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

---

# 5.6 Build and Test Backend Version 2

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

---

# 5.7 Deploy Backend V2 Using Docker Compose

The V2 Compose configuration is:

```text
docker-compose-v2.yml
```

It uses:

| Service | Image | Host Port |
|---|---|---:|
| Frontend | `microservices-frontend:1.0` | 8083 |
| Backend | `microservices-backend:v2` | 8084 |
| Database | `microservices-database:1.0` | 5434 |

### Project Implementation

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

---

# 5.8 Publish Backend V2 to Docker Hub

The general Docker Hub tag and push syntax is already covered in **Task 4** and **Task 5.1**.

### Project Implementation

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

---

# Project Details

| Detail | Information |
|---|---|
| Name | Snehal Shinde |
| Project | Project 2.2, Project 2.3, Project 2.4 |
| Assignments | Containerizing Microservices and Publishing Images to Registries; Automating Docker Image Builds using Jenkins CI/CD Pipeline; Deploying Multi-Container Applications using Docker Compose |
| GitHub Repository | `SnehalShinde11/docker-microservices-app` |
| Branch | `main` |
| Docker Hub Username | `snehalshinde11` |

---

# Project 2.3 — Automating Docker Image Builds using Jenkins CI/CD Pipeline

Project 2.3 automates the backend Docker image build and publishing process using Jenkins CI/CD.

### Project 2.3 README

[Project 2.3 — Automating Docker Image Builds using Jenkins CI/CD Pipeline](./README-Project-2.3.md)

---

# Project 2.4 - Deploying Multi-Container Applications using Docker Compose
