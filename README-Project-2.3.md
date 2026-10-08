# Project 2.3 — Automating Docker Image Builds using Jenkins CI/CD Pipeline

## 1. Problem Statement

The objective of this project is to automate the process of building and publishing a Docker image for the backend microservice using a Jenkins CI/CD pipeline.

The pipeline automates the following activities:

- Checkout application source code from GitHub
- Build the backend application
- Build the Docker image
- Tag the Docker image using the Jenkins build number
- Authenticate with Docker Hub securely
- Push the Docker image to Docker Hub
- Trigger the Jenkins pipeline automatically using a GitHub Webhook
- Visualize the pipeline stages in Jenkins

### CI/CD Flow

\`\`\`text
Developer
    |
    | git push
    v
GitHub Repository
    |
    | Webhook
    v
Jenkins
    |
    +--> Checkout Code
    |
    +--> Build Application
    |
    +--> Build Docker Image
    |        |
    |        +--> Tag with BUILD_NUMBER
    |
    +--> Push to Docker Hub
             |
             v
      Docker Hub Registry
\`\`\`

---

# 2. Dependencies and Prerequisites

Before starting the project, ensure the following are available:

| Dependency | Purpose |
|---|---|
| Linux VM / Server | Jenkins and Docker environment |
| Jenkins | CI/CD automation |
| Java | Jenkins runtime |
| Docker Engine | Build and push Docker images |
| Git | Source code management |
| GitHub | Application source repository |
| Docker Hub | Container image registry |
| Docker Hub Personal Access Token | Secure Docker Hub authentication |
| Jenkins Pipeline | Pipeline automation |
| GitHub Webhook | Automated pipeline triggering |

---

# Task 1 — Prepare Application Repository

## 1.1 Clone Repository

### General Syntax

\`\`\`bash
git clone <GITHUB_REPOSITORY_URL>
cd <REPOSITORY_NAME>
\`\`\`

### Project

\`\`\`bash
git clone https://github.com/SnehalShinde11/docker-microservices-app.git
cd docker-microservices-app
\`\`\`

**Repository:**

\`\`\`text
https://github.com/SnehalShinde11/docker-microservices-app
\`\`\`

**Branch:**

\`\`\`text
main
\`\`\`

---

## 1.2 Backend Application Structure

The backend application is located inside the \`backend\` directory.

\`\`\`text
backend/
├── Dockerfile
├── Dockerfile.env
├── go.mod
├── main.go
└── main-env.go
\`\`\`

The Jenkins pipeline uses this backend directory as the Docker build context.

---

# Task 2 — Setup Jenkins Environment

## 2.1 Install Jenkins

Install Jenkins on the Linux VM/server.

### Official Jenkins Installation Documentation

\`\`\`text
https://www.jenkins.io/doc/book/installing/
\`\`\`

After installation, verify Jenkins:

\`\`\`bash
sudo systemctl status jenkins
\`\`\`

Start Jenkins:

\`\`\`bash
sudo systemctl start jenkins
\`\`\`

Enable Jenkins at system startup:

\`\`\`bash
sudo systemctl enable jenkins
\`\`\`

---

## 2.2 Verify Java

Check the installed Java version:

\`\`\`bash
java -version
\`\`\`

Jenkins requires a supported Java version to run.

---

## 2.3 Configure Docker Access for Jenkins

Jenkins must have permission to execute Docker commands.

Add the Jenkins user to the Docker group:

\`\`\`bash
sudo usermod -aG docker jenkins
\`\`\`

Restart Jenkins:

\`\`\`bash
sudo systemctl restart jenkins
\`\`\`

Verify Jenkins user's groups:

\`\`\`bash
id jenkins
\`\`\`

Test Docker access as the Jenkins user:

\`\`\`bash
sudo su -s /bin/bash jenkins -c 'docker ps'
\`\`\`

Verify Docker version:

\`\`\`bash
sudo su -s /bin/bash jenkins -c 'docker version'
\`\`\`

---

## 2.4 Access Jenkins

Retrieve the initial administrator password:

\`\`\`bash
sudo cat /var/lib/jenkins/secrets/initialAdminPassword
\`\`\`

Find the server IP:

\`\`\`bash
hostname -I
\`\`\`

### General Syntax

\`\`\`text
http://<JENKINS_SERVER_IP>:<JENKINS_PORT>
\`\`\`

### Project Example

\`\`\`text
http://<SERVER-IP>:8080
\`\`\`

Open the Jenkins URL in a browser and complete the initial Jenkins setup.

---

## 2.5 Verify GitHub Connectivity

Test whether the Jenkins server can access the GitHub repository.

### General Syntax

\`\`\`bash
git ls-remote <GITHUB_REPOSITORY_URL>
\`\`\`

### Project Example

\`\`\`bash
git ls-remote https://github.com/SnehalShinde11/docker-microservices-app.git
\`\`\`

A successful command should display the repository references and commit hashes.

---

## 2.6 Configure Docker Hub Credentials

Jenkins requires Docker Hub credentials to push the generated image.

Navigate to:

\`\`\`text
Jenkins
→ Manage Jenkins
→ Credentials
→ Jenkins
→ Global credentials
→ Add Credentials
\`\`\`

Select:

\`\`\`text
Kind: Username with password
\`\`\`

### General Configuration

\`\`\`text
Username: <DOCKER_HUB_USERNAME>
Password: <DOCKER_HUB_PERSONAL_ACCESS_TOKEN>
ID: <DOCKER_HUB_CREDENTIAL_ID>
\`\`\`

### Project Configuration

\`\`\`text
Username: snehalshinde11
Credential ID: dockerhub-creds
\`\`\`

The Jenkinsfile references the credential using:

\`\`\`groovy
credentialsId: 'dockerhub-creds'
\`\`\`

The Docker Hub Personal Access Token should be stored in Jenkins Credentials rather than directly inside the Jenkinsfile.

---

# Task 3 — Create Jenkins Pipeline Job

Create a new Jenkins Pipeline job.

Navigate to:

\`\`\`text
Jenkins Dashboard
→ New Item
\`\`\`

### General Syntax

\`\`\`text
Job Name: <JENKINS_JOB_NAME>
\`\`\`

### Project Configuration

\`\`\`text
Job Name: microservices-backend-ci
\`\`\`

Select:

\`\`\`text
Pipeline
\`\`\`

Click **OK**.

---

## 3.1 Configure Pipeline from SCM

Configure the pipeline to retrieve the Jenkinsfile directly from GitHub.

Select:

\`\`\`text
Definition:
Pipeline script from SCM
\`\`\`

Select:

\`\`\`text
SCM:
Git
\`\`\`

### General Repository Syntax

\`\`\`text
https://github.com/<GITHUB_USERNAME>/<REPOSITORY_NAME>.git
\`\`\`

### Project Repository

\`\`\`text
https://github.com/SnehalShinde11/docker-microservices-app.git
\`\`\`

### General Branch Syntax

\`\`\`text
*/<BRANCH_NAME>
\`\`\`

### Project Branch

\`\`\`text
*/main
\`\`\`

### General Jenkinsfile Path

\`\`\`text
<JENKINSFILE_PATH>
\`\`\`

### Project Jenkinsfile

\`\`\`text
Jenkinsfile
\`\`\`

Since the repository is public, repository credentials are not required for checkout.

---

# Task 4 — Define Pipeline Stages in Jenkinsfile

The Jenkinsfile defines the CI/CD pipeline and its execution stages.

### General Jenkinsfile Location

\`\`\`text
<REPOSITORY_ROOT>/Jenkinsfile
\`\`\`

The pipeline contains four actual stages:

\`\`\`text
1. Checkout Code
2. Build Application
3. Build Docker Image
4. Push to Docker Hub
\`\`\`

> Note: There is no separate tagging stage. The Docker image is tagged with \${BUILD_NUMBER} during the Docker build command itself.

---

## 4.1 Checkout Code

The first stage retrieves the source code from GitHub into the Jenkins workspace.

### General Syntax

\`\`\`groovy
stage('Checkout Code') {
    steps {
        git branch: '<BRANCH_NAME>',
            url: '<GITHUB_REPOSITORY_URL>'
    }
}
\`\`\`

The pipeline checks out the specified Git branch before executing the remaining stages.

---

## 4.2 Build Application

The backend Go application is compiled before building the Docker image.

### General Syntax

\`\`\`groovy
stage('Build Application') {
    steps {
        sh '''
            <BUILD_COMMAND>
        '''
    }
}
\`\`\`

### General Go Build Pattern

\`\`\`bash
docker run --rm \\
    -v "$WORKSPACE/<APPLICATION_DIRECTORY>:/app" \\
    -w /app \\
    <GO_IMAGE>:<GO_VERSION> \\
    go build -o <OUTPUT_FILE> <SOURCE_FILE>
\`\`\`

This approach runs the Go build inside a Go Docker container.

Advantages:

- No Go installation is required directly on the Jenkins server.
- A fixed Go version can be used.
- The Jenkins workspace is mounted into the build container.
- The generated binary is available in the workspace.

---

## 4.3 Build Docker Image

The Docker image is built from the backend application's Dockerfile.

### General Syntax

\`\`\`bash
docker build \\
    -t <DOCKER_HUB_USERNAME>/<IMAGE_NAME>:\${BUILD_NUMBER} \\
    <BUILD_CONTEXT>
\`\`\`

The Jenkins \${BUILD_NUMBER} environment variable is used as the Docker image tag.

For example:

\`\`\`text
Build #1  → image:1
Build #2  → image:2
Build #3  → image:3
\`\`\`

This provides a unique image tag for every Jenkins build.

### Important

The image is **tagged during the Docker build itself** using the \`-t\` option.

There is no separate tagging stage.

---

## 4.4 Push to Docker Hub

After successfully building the image, Jenkins authenticates with Docker Hub and pushes the image.

### General Docker Hub Repository

\`\`\`text
<DOCKER_HUB_USERNAME>/<IMAGE_NAME>
\`\`\`

### General Login Syntax

\`\`\`bash
echo "$DOCKER_PASSWORD" | docker login \\
    --username "$DOCKER_USERNAME" \\
    --password-stdin
\`\`\`

Jenkins credentials can be injected into the pipeline using:

### General Jenkins Credentials Syntax

\`\`\`groovy
withCredentials([
    usernamePassword(
        credentialsId: '<DOCKER_HUB_CREDENTIAL_ID>',
        usernameVariable: 'DOCKER_USERNAME',
        passwordVariable: 'DOCKER_PASSWORD'
    )
]) {
    sh '''
        <DOCKER_COMMANDS>
    '''
}
\`\`\`

### General Push Syntax

\`\`\`bash
docker push <DOCKER_HUB_USERNAME>/<IMAGE_NAME>:\${BUILD_NUMBER}
\`\`\`

### General Image Reference

\`\`\`text
<DOCKER_HUB_USERNAME>/<IMAGE_NAME>:<BUILD_NUMBER>
\`\`\`

The Docker Hub password/token is not hardcoded in the Jenkinsfile. Jenkins retrieves it securely from its Credentials store.

---

## 4.5 Complete Jenkins Pipeline Flow

\`\`\`text
                    GitHub Repository
                           |
                           | Checkout
                           v
                    +--------------+
                    |    Jenkins   |
                    +--------------+
                           |
                           v
                  +------------------+
                  | Checkout Code    |
                  +------------------+
                           |
                           v
                  +------------------+
                  | Build Application|
                  +------------------+
                           |
                           v
                  +------------------+
                  | Build Docker     |
                  | Image            |
                  |                  |
                  | Tag:             |
                  | BUILD_NUMBER     |
                  +------------------+
                           |
                           v
                  +------------------+
                  | Push to          |
                  | Docker Hub       |
                  +------------------+
                           |
                           v
                    Docker Hub
\`\`\`

---

# Task 5 — Verify CI Pipeline Execution

After configuring the Jenkins pipeline, manually trigger the first build.

Navigate to:

\`\`\`text
Jenkins Dashboard
→ microservices-backend-ci
→ Build Now
\`\`\`

The pipeline should execute the following stages:

\`\`\`text
Checkout Code
      ↓
Build Application
      ↓
Build Docker Image
      ↓
Push to Docker Hub
\`\`\`

---

## 5.1 Verify Build Number

Every Jenkins execution receives a unique build number.

Example:

\`\`\`text
Build #1
Build #2
Build #3
\`\`\`

The build number is used as the Docker image tag.

---

## 5.2 Verify Console Output

Open:

\`\`\`text
Jenkins
→ microservices-backend-ci
→ Build Number
→ Console Output
\`\`\`

A successful pipeline should end with:

\`\`\`text
Finished: SUCCESS
\`\`\`

---

## 5.3 Verify Docker Image

### General Image Reference

\`\`\`text
<DOCKER_HUB_USERNAME>/<IMAGE_NAME>:<BUILD_NUMBER>
\`\`\`

### Project Image

\`\`\`text
snehalshinde11/microservices-backend:<BUILD_NUMBER>
\`\`\`

Verify that the corresponding tag has been pushed to Docker Hub.

---

# Task 6 — Configure Automated GitHub Webhook Trigger

A GitHub Webhook allows GitHub to notify Jenkins whenever new code is pushed to the repository.

This removes the need to manually click **Build Now** after every code change.

---

## 6.1 Configure GitHub Webhook

First configure the Jenkins job.

Navigate to:

\`\`\`text
Jenkins
→ microservices-backend-ci
→ Configure
→ Build Triggers
\`\`\`

Enable:

\`\`\`text
GitHub hook trigger for GITScm polling
\`\`\`

Now configure the webhook in GitHub.

Navigate to:

\`\`\`text
GitHub Repository
→ Settings
→ Webhooks
→ Add webhook
\`\`\`

### General Payload URL Syntax

\`\`\`text
http://<JENKINS_SERVER_IP>:<JENKINS_PORT>/github-webhook/
\`\`\`

### Project Example

\`\`\`text
http://35.223.119.70:8080/github-webhook/
\`\`\`

> If the Jenkins server IP or port changes, update the webhook Payload URL accordingly.

Configure the webhook as follows:

\`\`\`text
Payload URL:
http://<JENKINS_SERVER_IP>:<JENKINS_PORT>/github-webhook/

Content type:
application/json

Which events:
Just the push event

Active:
Enabled
\`\`\`

Click **Add webhook**.

### Webhook Flow

\`\`\`text
Developer
    |
    | git push
    v
GitHub
    |
    | Webhook
    v
Jenkins /github-webhook/
    |
    v
Jenkins Pipeline
    |
    +--> Checkout
    +--> Build
    +--> Docker Build
    +--> Docker Push
\`\`\`

---

## 6.2 Test GitHub Webhook Trigger

Make a small change to the backend application and push it to GitHub.

Navigate to the project:

\`\`\`bash
cd ~/microservices-app
\`\`\`

Check the current backend response:

\`\`\`text
Backend Service is running
\`\`\`

Modify the response:

\`\`\`bash
sed -i 's/Backend Service is running/Backend Service is running - Jenkins Webhook Test/' backend/main.go
\`\`\`

Verify the change:

\`\`\`bash
git diff
\`\`\`

Stage the change:

\`\`\`bash
git add backend/main.go
\`\`\`

Commit the change:

\`\`\`bash
git commit -m "Test GitHub webhook trigger"
\`\`\`

Push the change:

\`\`\`bash
git push origin main
\`\`\`

### Expected Flow

\`\`\`text
git push
    ↓
GitHub
    ↓
Webhook
    ↓
Jenkins
    ↓
New Build Automatically Started
    ↓
Checkout Code
    ↓
Build Application
    ↓
Build Docker Image
    ↓
Push to Docker Hub
\`\`\`

Verify the new build under:

\`\`\`text
Jenkins
→ microservices-backend-ci
→ Build History
\`\`\`

The new build should start automatically without clicking **Build Now**.

---

# Task 7 — Pipeline Stage Visualization

Jenkins provides a visual representation of the pipeline stages.

Navigate to:

\`\`\`text
Jenkins Dashboard
→ microservices-backend-ci
\`\`\`

The pipeline should display stages similar to:

\`\`\`text
+----------------+-------------------+--------------------+-------------------+
| Checkout Code  | Build Application | Build Docker Image | Push to Docker Hub|
+----------------+-------------------+--------------------+-------------------+
        ✓                  ✓                    ✓                    ✓
\`\`\`

Each successful stage should be displayed as completed.

If the stage visualization is not available, verify that the Jenkins Pipeline Stage View plugin is installed:

\`\`\`text
Manage Jenkins
→ Plugins
→ Installed
→ Pipeline: Stage View
\`\`\`

---

# Project Details

| Item | Details |
|---|---|
| Name | Snehal Shinde |
| Project | Project 2.3 |
| Assignment | Automating Docker Image Builds using Jenkins CI/CD Pipeline |
| GitHub Repository | \`SnehalShinde11/docker-microservices-app\` |
| Branch | \`main\` |
| Jenkins Job | \`microservices-backend-ci\` |
| Pipeline Definition | \`Jenkinsfile\` |
| Docker Hub Username | \`snehalshinde11\` |
| Docker Hub Repository | \`snehalshinde11/microservices-backend\` |
| Jenkins Credential ID | \`dockerhub-creds\` |
| Application | Backend Microservice |
| Application Language | Go |
| Docker Image | \`snehalshinde11/microservices-backend:<BUILD_NUMBER>\` |
| CI/CD Tool | Jenkins |
| Containerization | Docker |
| Source Control | GitHub |
| Container Registry | Docker Hub |
| Automated Trigger | GitHub Webhook |
| Webhook Endpoint | \`/github-webhook/\` |
