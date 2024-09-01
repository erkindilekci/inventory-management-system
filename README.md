# Inventory Management App

## Overview
This project is a web application that leverages a modern tech stack to provide a robust and scalable solution. The backend is built with Go and Postgres, while the frontend is developed using React and TypeScript.

## Screenshot
![imsscreenshot](https://github.com/user-attachments/assets/73eb210e-dbd1-4700-9c68-02b63d0ea3a2)

## Technologies

### Backend
- **Go**: A statically typed, compiled programming language designed for simplicity and efficiency.
- **Postgres**: A powerful, open-source object-relational database system with a strong reputation for reliability and performance.
- **Echo**: A high-performance, minimalist web framework for Go, designed for ease of use and scalability.
- **github.com/golang-jwt/jwt/v5**: A Go implementation of JSON Web Tokens (JWT) for secure authentication.
- **github.com/jackc/pgx/v4**: A PostgreSQL driver and toolkit for Go, providing efficient and feature-rich database interactions.
- **github.com/labstack/gommon**: A set of common packages for Go, including logging, color, and bytes utilities.

### Frontend
- **React**: A JavaScript library for building user interfaces, maintained by Facebook and a community of individual developers and companies.
- **TypeScript**: A strongly typed programming language that builds on JavaScript, giving you better tooling at any scale.
- **Tailwind CSS**: A utility-first CSS framework for rapidly building custom user interfaces.
- **js-cookie**: A simple, lightweight JavaScript API for handling cookies.

## Getting Started

### Prerequisites
- Go 1.23.0 or later
- Node.js 14.x or later
- Postgres 12 or later

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/erkindilekci/inventory-management-system.git
   cd yourproject
   ```

2. **Backend Setup:**
   ```sh
   cd backend
   go mod download
   go build -o ./ims cmd/imsapi/main.go
   ```

3. **Frontend Setup:**
   ```sh
   cd frontend
   npm install
   npm run dev
   ```

### Running the Application

1. **Start the Backend:**
   ```sh
   cd backend
   ./ims
   ```

2. **Start the Frontend:**
   ```sh
   cd frontend
   npm run dev
   ```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
