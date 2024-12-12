# GraphQL with Golang on Docker

A simple GraphQL server using Golang, coupled with Docker for easy execution in any environment.

## Description

This is a basic GraphQL server implemented with Golang.

### Changes:
- The server now serves a main.go file to make it easier to test GraphQL through a browser.

## Technologies used

- Golang
- GraphQL
- Docker

## Prerequisites

To run this project, you must have Docker installed on your machine. If you don't have it, you can download it from [here](https://www.docker.com/products/docker-desktop).

## Instructions to run the project

1. *Clone this repository:*

If you haven't cloned the repository yet, you can do so with the following command:

bash
git clone git clone git clone https://github.com/SantiagoZ98/GraphQL_go2.git

2. **Build the Docker image:**

Before running the container, build the Docker image with the following command:

bash
docker build -t santiagozurita26/my-go-arq .

3. *Push the image to Docker Hub (if necessary):*

If you want to push the image to Docker Hub for others to use, you can do so with:

bash
docker push santiagozurita26/my-go-arq:tagname

4. **Run the Docker container:**

After building the image, run the container with the following command:

bash
docker pull santiagozurita26/my-go-arq

5. **Test the connection:**

-**When a user connects,** they will receive the message "Hello World, my name is Santiago Zurita, I am a student at UCE."
