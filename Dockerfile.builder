FROM golang:1.23.2 AS base-graphkubel-builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./internal/common /app/internal/common