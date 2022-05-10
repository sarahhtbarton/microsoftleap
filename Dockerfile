# syntax=docker/dockerfile:1

FROM golang:1.18 AS builder

RUN mkdir /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./sbapp ./cmd/server/

FROM scratch

COPY --from=builder /app/sbapp .

EXPOSE 8080

CMD [ "/sbapp" ]