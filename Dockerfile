# --- build stage ---
FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY main.go .

RUN go build -o stub

# --- runtime stage ---
FROM alpine:3.19

WORKDIR /app

COPY --from=build /app/stub .

EXPOSE 8081

CMD ["./stub"]
