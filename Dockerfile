# syntax=docker/dockerfile:1

# What image do you want to start building on?
FROM golang:1.18

# Make a folder in your image where your app's source code can live
RUN mkdir /app

# Tell your container where your app's source code will live. This is the working directory
WORKDIR /app

# What source code do you what to copy, and where to put it?
# COPY takes two arguments (It's a little hard to tell, but . and /src/app are separated by a space). The first argument (the .) says where to copy the app's source code from. In this case . is a relative path that points to the directory the Dockerfile is currently in.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# RUN go build -o /docker-gs-ping

# EXPOSE 8080

# CMD [ "/docker-gs-ping" ]