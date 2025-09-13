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

*go mod download

**.env
```
PORT=8080
DB_HOST=localhost
DB_PORT=3306
DB_NAME=school_census
DB_USER=root
DB_PASSWORD=your_password
JWT_SECRET=your_super_secret_jwt_key
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret
CLIENT_URL=http://localhost:3000
   ```bash
   git clone https://github.com/your-username/school-census-backend.git
   cd school-census-backend
