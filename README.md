# Chat 💬

A simple web-based chat application built in **Go** with server-side rendering using **HTML templates**.  
The project demonstrates how to build a lightweight real-time chat service with clean architecture and environment-based configuration.

---

## ✨ Features

- **Web-based chat interface** with templated HTML views  
- **Go backend** for handling requests and rendering templates  
- **Environment-based configuration** using `.env`  
- **Container-ready** with simple deployment setup  
- **Extendable** structure with `cmd/`, `internal/`, and `pkg` packages  

---

## 📂 Project Structure

```
chat/
├── assets/js/              # Client-side JavaScript (if needed)
├── cmd/                    # Server entrypoints
├── internal/               # Internal Go packages
├── pkg/utils/              # Utility functions
├── templates/              # HTML templates for UI
├── .env.example            # Example environment variables
├── LICENSE                 # License file
├── go.mod / go.sum         # Go modules and dependencies
└── main.go                 # Application entrypoint
```

---

## ⚙️ Getting Started

### Prerequisites

- [Go 1.18+](https://go.dev/)  
- [Docker](https://www.docker.com/) (optional)  

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/alirezadp10/chat.git
   cd chat
   ```

2. Copy the environment file:

   ```bash
   cp .env.example .env
   ```

   Update `.env` with your configuration (server port, DB connection, etc.).

3. Install Go dependencies:

   ```bash
   go mod download
   ```

---

### ▶️ Running Locally

```bash
go run main.go
```

By default, the app runs on `http://localhost:8080`.

---

### 🐳 Running with Docker

Build and run the containerized setup:

```bash
docker build -t chat-app .
docker run -p 8080:8080 chat-app
```

---

## 🛠️ Tech Stack

- **Go** – Backend server  
- **HTML templates** – Server-side rendered views  
- **JavaScript (optional)** – Client-side enhancements  
- **Docker** – Deployment-ready containerization  

---

## 🤝 Contributing

Contributions are welcome! To get started:

1. Fork the repo  
2. Create a new branch (`feature/my-feature`)  
3. Commit your changes  
4. Submit a Pull Request  

---

## 📜 License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.
