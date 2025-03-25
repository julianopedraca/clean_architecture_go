# Clean Architecture Go
Clean Architecture is a software design pattern that separates the application's business logic from its infrastructure and presentation layers.
This repository provides an example implementation of Clean Architecture in Go.

## Overview
The project is structured into several layers:
- **external**: implements sp trans api
- **interfaces**: defines the interfaces for the sp trans api
- **internal**:
  - **controller**: Register routes and tests
  - **router**: setup gin engine

## Getting Started
To get started with the project, follow these steps:
1. Clone the repository
2. Run `make build` to install dependencies and run the aplication
3. After the first build you can use `make start`

## techstack
Go, Gin, CI/CD, Docker, Clean Architecture
