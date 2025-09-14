# School Census Go Backend

A high-performance, secure backend API for a School Census Management System built with Go and Gin.

## 🚀 Features

- **JWT Authentication** - Secure email/password login with refresh tokens
- **Protected Routes** - Middleware for role-based access control
- **RESTful API** - Clean and consistent API design
- **File Upload** - Cloudinary integration for image management
- **CORS Enabled** - Configured for cross-origin requests
- **Database ORM** - GORM for efficient database operations
- **Environment Configuration** - Secure management of sensitive data

## 🛠️ Tech Stack

- **Framework**: Gin Gonic
- **Database**: MySQL with GORM
- **Authentication**: JWT with bcrypt password hashing
- **File Storage**: Cloudinary
- **Environment**: Go 1.21+
- **CORS**: Gin CORS middleware

## 📦 Installation

1. **Clone the repository**
 ```bash
   git clone https://github.com/your-username/school-census-backend.git
   cd school-census-backend
```
2. ** Install packages**
go mod tidy
go run .
