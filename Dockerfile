# Start from golang base image
FROM golang:1.17 as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Muhamad Ryan Adi Wiguna <ryan.adi911@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apt update

# Set the current working directory inside the container 
WORKDIR /apps

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./app/

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates && apk add --no-cache bash

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /apps/main . 

#Command to run the executable
ENTRYPOINT ["./main"]

# FROM golang:1.17

# # create a directory named /app
# RUN mkdir /app

# # make /app our working directory
# WORKDIR /app

# # copy all files from this directory to /app
# COPY ./ /app
# # build/compile
# RUN cd ./app && go build -o rest-api

# # the command that should be executed when
# # the container is started
# CMD cd ./app && ./rest-api
