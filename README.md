# Golang starter

This project is an example of a web application built using Go (Golang) with exclusively native code. The goal of this project is to demonstrate how to create a web application with a layered architecture in Go without relying on any third-party libraries or frameworks.

- [Golang starter](#golang-starter)
  - [🚚 Features](#-features)
  - [📝 Prerequisites](#-prerequisites)
  - [🚀 Quick start](#-quick-start)
  - [🗂️ Project Structure](#️-project-structure)
- [📜 License](#-license)

## 🚚 Features

- Pure Go code.
- Layered architecture.
- Simple web server with basic routing.
- Fetching data from an external API.

## 📝 Prerequisites

Before you begin, ensure you have met the following requirements:

- [Golang](https://go.dev/doc/install) (v1.21 or higher).

## 🚀 Quick start

1. Clone the repository:
   ```sh
   git clone https://github.com/anthonypillot/starter-golang.git
   ```
2. Navigate to the project directory:
   ```sh
   cd starter-golang
   ```
3. Run the application:
   ```sh
   cd src
   go run main.go
   ```

The web server will start, and you can access the application by navigating to `http://localhost:8080` in your web browser.

## 🗂️ Project Structure

- [`main.go`](./src/main.go): The entry point of the application.
- [`controller`](./src/controller): Contains the controller functions for handling requests.
- [`service`](./src/service): Contains the business logic for the application.
- [`dao`](./src/dao): Contains the data access objects for interacting with the database or other data sources like APIs.
- [`util`](./src/util): Contains utility functions used throughout the application.
- [`model`](./src/model): Contains the data models used in the application.

# 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
