pipeline {

    agent any

    stages {

        stage('Checkout Code') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/SnehalShinde11/docker-microservices-app.git'
            }
        }

        stage('Build Application') {
            steps {
                sh '''
                    docker run --rm \
                        -v "$WORKSPACE/backend:/app" \
                        -w /app \
                        golang:1.26 \
                        go build -o backend main.go

                    rm -f backend/backend
                '''
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '''
                    docker build \
                        -t snehalshinde11/microservices-backend:${BUILD_NUMBER} \
                        ./backend
                '''
            }
        }

        stage('Push to Docker Hub') {
            steps {
                withCredentials([
                    usernamePassword(
                        credentialsId: 'dockerhub-creds',
                        usernameVariable: 'DOCKER_USERNAME',
                        passwordVariable: 'DOCKER_PASSWORD'
                    )
                ]) {
                    sh '''
                        echo "$DOCKER_PASSWORD" | docker login \
                            --username "$DOCKER_USERNAME" \
                            --password-stdin

                        docker push \
                            snehalshinde11/microservices-backend:${BUILD_NUMBER}

                        docker logout
                    '''
                }
            }
        }
    }
}

