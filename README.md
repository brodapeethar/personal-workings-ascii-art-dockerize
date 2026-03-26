# personal-workings-ascii-art-dockerize
An ASCII art web service written in Go that converts text into stylized ASCII output. Packaged with Docker for consistent builds, portability, and quick deployment.

## Table of Contents

Description
Features
Project Structure
Prerequisites
Installation
Usage
API Documentation
Technologies Used
Author

##  Description
ASCII Art Web is a web-based application that transforms regular text into stylized ASCII art. Users can select from three different banner styles (standard, shadow, thinkertoy) and instantly generate artistic text representations. The application is fully containerized using Docker, making it easy to deploy and run anywhere.
Live Demo: Access at http://localhost:3000 after running the container.

## Features

 * Multiple Banner Styles: Choose from standard, shadow, or thinkertoy fonts
 * Web Interface: Clean, responsive UI with dark theme
 * Dockerized: Runs consistently across all platforms
 * Secure: Runs as non-root user in container
 * Health Checks: Built-in container health monitoring
 * Optimized: Multi-stage Docker build for minimal image size (~20MB)
 * Newline Support: Handles \n for multi-line ASCII art
 * Special Characters: Supports all ASCII printable characters (32-126)


 ## Project Structure
 .
├── banners
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
├── Dockerfile-updated
├── dockerignore
├── generator
│   ├── art-generator.go
│   └── art-generator_test.go
├── go.mod
├── handlers
│   ├── formHandler.go
│   ├── homeHandler.go
│   └── utils.go
├── main.go
├── README.md
├── static
│   └── style.css
└── templates
    └── index.html

## 🔧 Prerequisites
Before running this project, ensure you have:

Docker installed (version 20.10 or higher)
Docker daemon running
Port 3000 available on your machine

Check Docker Installation:
bashdocker --version
# Expected: Docker version 20.10.x or higher

docker ps
# Should not return permission errors
Installing Docker:
Ubuntu/Debian:
bashsudo apt update
sudo apt install docker.io -y
sudo systemctl start docker
sudo systemctl enable docker

* Add your user to docker group (to avoid using sudo):
bashsudo usermod -aG docker $USER
newgrp docker


## Installation
Option 1: Clone Repository
bash# Clone the repository
git clone https://github.com/yourusername/ascii-art-web-dockerize.git
cd ascii-art-web-dockerize

Option 2: Manual Setup

Create project directory:

bashmkdir ascii-art-web-dockerize
cd ascii-art-web-dockerize

Create folder structure:

bashmkdir -p banners generator handlers static templates

Add all project files to their respective folders as shown in Project Structure


## Usage
Step 1: Build Docker Image
bashdocker build -t ascii-art-web:1.0 .
Expected Output:
[+] Building 15.2s (18/18) FINISHED
=> => naming to docker.io/library/ascii-art-web:1.0
Verify image created:
bashdocker images | grep ascii-art-web
Step 2: Run Docker Container

bashdocker run -d --name ascii-container -p 3000:3000 ascii-art-web:1.0
Flags explained:

-d = Run in detached mode (background)
--name ascii-container = Name the container
-p 3000:3000 = Map port 3000 (host:container)
ascii-art-web:1.0 = Image to use

Verify container is running:
bashdocker ps

Step 3: Access the Application
Open your browser and navigate to:
http://localhost:3000
You should see the ASCII Art Web interface!

Step 4: Generate ASCII Art

Enter your text in the textarea
Select a banner style (standard, shadow, or thinkertoy)
Click "GENERATE ART"
View your ASCII art output below the form



## API Documentation
Routes
GET /

Description: Displays the home page with the form
Response: HTML page with textarea and banner selector

POST /ascii-art

Description: Generates ASCII art from submitted text
Parameters:

text (string, required): The text to convert to ASCII art
banner (string, required): Banner style (standard, shadow, or thinkertoy)


Response: HTML page with generated ASCII art
Error Codes:

400 Bad Request: Missing text or banner selection
404 Not Found: Banner file not found
405 Method Not Allowed: Wrong HTTP method (must be POST)
500 Internal Server Error: ASCII art generation error



GET /static/*

Description: Serves static files (CSS, JS, images)
Example: GET /static/style.css


## Technologies Used
Backend

Go (Golang) - Programming language
net/http - HTTP server
html/template - Template rendering
os - File operations
strings - String manipulation

Frontend

HTML5 - Markup
CSS3 - Styling (dark theme with terminal aesthetics)
Vanilla JavaScript - Client-side interactions (form handling)

DevOps

Docker - Containerization
Alpine Linux - Base image (minimal footprint)
Multi-stage builds - Optimized image size


Notes

Application runs on port 3000 by default (can be changed in main.go)
Supports ASCII characters 32-126 (printable characters)
Handles newlines using \n in input text
Docker image uses multi-stage build for optimal size (~20MB)
Runs as non-root user for security
Includes health check for production deployments


🤝 Contributing
Contributions are welcome! Please follow these steps:

Fork the repository
Create a feature branch (git checkout -b feature/AmazingFeature)
Commit your changes (git commit -m 'Add some AmazingFeature')
Push to the branch (git push origin feature/AmazingFeature)
Open a Pull Request


📄 License
This project is part of the Zone01 curriculum.



Author
Email: olajidepeter7012@gmail.com
GitHub: @brodapeethar


## Additional Resources

Go Documentation
Docker Documentation
ASCII Art Generator
Dockerfile Best Practices