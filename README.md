# Clean Architecture in Go

This repository provides an example implementation of Clean Architecture in Go, ensuring separation of concerns between business logic, infrastructure, and presentation layers.

## 📌 Overview
The project is structured into several layers:

- **external/**: Implements the SPTrans API.
- **interfaces/**: Defines interfaces for the SPTrans API.
- **internal/**:
  - **controller/**: Registers routes and includes tests.
  - **router/**: Sets up the Gin engine.

## 🚀 Getting Started
Follow these steps to get the project up and running:

1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo.git
   cd your-repo
   ```
2. Build the project and install dependencies:
   ```sh
   make build
   ```
3. Start the application (after the first build):
   ```sh
   make start
   ```

## 🛠 Tech Stack
- **Go** - Main programming language
- **Gin** - Web framework
- **Docker** - Containerization
- **CI/CD** - Continuous Integration & Deployment
- **Clean Architecture** - Software design pattern

