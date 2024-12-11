# GraphQL with Golang in Docker

A simple GraphQL server using golang, dockerized for easy execution in any environment.

## Description

This is a basic GraphQL server implemented with golang. 

### Changes:
- Now, the server serves an main.go file for easy graphQL testing through a browser.

## Technologies Used

- Golang
- GraphQL
- Docker

## Prerequisites

To run this project, you need to have Docker installed on your machine. If you don't have it, you can download it from [here](https://www.docker.com/products/docker-desktop).

## Instructions to Run the Project

1. *Clone this repository:*

   If you haven't cloned the repository yet, you can do so with the following command:

   bash
   git clone git clone git clone https://github.com/SantiagoZ98/GraphQL_go2.git

2. **Build the Docker image:**

   Before running the container, build the Docker image using the following command:

   bash
   docker build -t santiagozurita26/my-go-arq2 .

3. *Push the image to Docker Hub (if needed):*

   If you'd like to upload the image to Docker Hub for others to use, you can do so with:

   bash
   docker push santiagozurita26/my-go-arq:tagname

4. **Run the Docker container:**

   After building the image, run the container with the following command:

   bash
   docker pull santiagozurita26/my-go-arq

5. **Test the connection:**

-**When a client connects,** it will receive the message "Hello, World!".
-**If the client sends a message back to the server,** the server will log that message to the console


#### Server:
1. The server will receive the message and print it in the terminal.
2. The server will respond with the message `"Hello World, my name is Santiago Zurita!"`.

