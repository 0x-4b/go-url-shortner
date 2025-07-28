# URL Shortener

[![Go](https://img.shields.io/badge/Go-1.24%2B-blue.svg)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-1.10-green.svg)](https://gin-gonic.com/)
[![Redis](https://img.shields.io/badge/Redis-6.0%2B-red.svg)](https://redis.io/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A fast, minimal URL shortener built with Go, Gin, and Redis.  
Built for backend lovers — no frontend, no bloat. Just clean REST API endpoints and a powerful Redis-backed short link engine.

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Technical Details](#technical-details)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Author](#author)

---

## Overview

**URL Shortener** is a high-performance web service that converts long URLs into short, shareable links.  
It features a RESTful API with custom short code generation, per-user tracking, and Redis-based in-memory storage for fast redirects.

---

## Features

- 🚀 **Fast REST API** — Built with the Gin framework
- 🔗 **Custom Short Codes** — SHA256 + Base58 encoding
- 💾 **Redis Storage** — In-memory cache for blazing fast access
- 👤 **User-based Links** — Each link is tied to a user ID
- 🛡️ **Input Validation** — Basic format checking for URLs
- 🧪 **Unit Tested** — Core logic and storage tested

---

## Installation

### Prerequisites

- [Go 1.24+](https://golang.org/dl/)
- [Redis](https://redis.io/download)

### Quick Start

```bash
# Clone the repo
git clone https://github.com/0x-4b/go-url-shortner.git
cd go-url-shortner

# Install Go dependencies
go mod download

# Start Redis server (in a separate terminal)
redis-server

# Run the app
go run main.go
```

---

## Usage

### Shorten a URL

```bash
curl -X POST http://localhost:9808/create-short-url   -H "Content-Type: application/json"   -d '{"long_url":"https://example.com/very/long/url", "user_id":"user123"}'
```

**Response:**
```json
{
  "message": "short url created successfully",
  "short_url": "http://localhost:9808/abc123"
}
```

### Access the Short URL

```bash
curl -i http://localhost:9808/abc123
```

**Response:** `302 Redirect` to original long URL

### Health Check

```bash
curl http://localhost:9808/
```

**Response:**
```json
{
  "message": "Welcome to the URL Shortner API"
}
```

---

## Project Structure

```
go-url-shortner/
├── handler/
│   └── handlers.go                # HTTP route handlers
├── shortner/
│   ├── shorturl_generator.go     # SHA256 + Base58 encoding logic
│   └── shorturl_generator_test.go
├── store/
│   ├── store_service.go          # Redis interaction logic
│   └── store_service_test.go
├── main.go                       # Entry point
├── go.mod                        # Dependencies
└── README.md                     # You're here
```

---

## Technical Details

### Shortening Logic

- **Hashing**: SHA256 hash of (long URL + user ID)
- **Encoding**: Encoded into Base58 (Bitcoin alphabet)
- **Short Code**: First 8 characters of encoded string

### Storage

- **Backend**: Redis
- **Key Format**: `short_code -> long_url`
- **Expiration**: 6 hours (can be changed)

### Performance

- ⚡ Sub-millisecond lookups from Redis
- ✅ Handles concurrent requests cleanly
- 🧱 No external dependencies besides Redis

---

## Roadmap

- [ ] Click analytics
- [ ] Custom user-defined slugs
- [ ] Expiration timestamps
- [ ] QR code generation
- [ ] Bulk shortening support
- [ ] Optional persistent DB (maybe, one day… 🫠)

---

## Contributing

Contributions welcome!

1. Fork the repo
2. Create a branch (`git checkout -b feature/something`)
3. Write code + tests
4. Push and open a PR

Please write idiomatic Go and add tests where it makes sense.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Author

**[@0x-4b](https://github.com/0x-4b)**
