# Go-based downstream stub (Deterministic Service Simulator)

## 📌 What is this?

A minimal, high-performance HTTP stub written in Go, designed to simulate downstream services with:

- Deterministic latency
- Controlled concurrency
- Near-zero overhead runtime (compiled binary)

This stub is intentionally not a real service — it is a precision tool for performance experiments.

## ⚙️ Features

- 🧵 Controlled concurrency via semaphore
- ⏱ Deterministic delay per request
- 🚀 Compiled binary (no JVM, no runtime)
- 🐳 Docker-ready
- 📉 Minimal overhead → does not distort benchmark results


## ▶️ Run locally (no build yet)

**Run**
```bash
go run main.go
```

**Test:**
```bash
curl http://localhost:8081/delay
```

## ⚙️ Build binary (recommended)

**Build**
```bash
go build -o stub
```

**Run:**

```bash
./stub
```

👉 This is your production stub

## 📦 Dockerize (minimal, correct)

Dockerfile is [here](https://github.com/popov-rnd/stub-go/blob/main/Dockerfile)

**Build image**

```bash
docker build -t stub:go .
```

**Run container**

```bash
docker run -p 8081:8081 \
  -e STUB_DELAY_MS=500 \
  -e STUB_CONCURRENCY=2000 \
  stub:go
```

**Test:**
```bash
curl http://localhost:8081/delay
```

## 🔑 Configuration

- *STUB_DELAY_MS*: Delay per request (milliseconds), defauts to 500;
- *STUB_CONCURRENCY*: Max concurrent requests (semaphore), defaults to 10000.


## ⚠️ Important Notes

- This is not a production service
- This is a controlled experiment component
- Behavior is intentionally simplified and deterministic

