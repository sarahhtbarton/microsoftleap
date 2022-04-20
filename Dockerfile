# syntax=docker/dockerfile:1

# What image do you want to start building on?
FROM golang:1.18

# Make a folder in your image where your app's source code can live
RUN mkdir /app

# Tell your container where your app's source code will live. This is the working directory
WORKDIR /app

# What source code do you what to copy, and where to put it?
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /sbapp ./cmd/server/

EXPOSE 8080

CMD [ "/sbapp" ]