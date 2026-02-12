# Clothes Store

A premium, modern Online store platform built with Go and MongoDB. Features a luxury design aesthetic with a fully functional shopping experience and a comprehensive administrative dashboard.

## Team
- Yskak Zhanibek
- Nauanov Alikhan
- Zhumagali Beibarys

## Features

### Customer Experience
- **Modern Shop**: Advanced filtering (category, gender, color, size) and sorting.
- **Product Details**: High-quality imagery, size selection, and stock status.
- **Cart & Wishlist**: Persistent client-side shopping cart and wishlist management.
- **Checkout**: Seamless checkout flow with address management and order confirmation.
- **User Accounts**: Registration, login, and order history tracking.

### Admin Dashboard
- **Analytics**: Key performance indicators (Total Sales, Orders, Users).
- **Product Management**: Complete CRUD with local image uploads and advanced validation.
- **Order Management**: Track and update order statuses.
- **User Management**: Overview of registered users.

## Tech Stack

- **Backend**: Go (Gin Web Framework)
- **Database**: MongoDB
- **Authentication**: JWT (JSON Web Tokens) with Secure Cookies
- **Frontend**: Semantic HTML5, Vanilla CSS (Modern CSS variables), JavaScript (ES6+)
- **Icons**: Lucide Icons

## Project Structure

```bash
├── cmd/
│   └── server/          # Entry point (main.go)
├── internal/
│   ├── api/             # Routing and Middleware
│   ├── config/          # Environment configuration
│   ├── db/              # Database connection
│   ├── handlers/        # HTTP Handlers
│   ├── models/          # Data structures
│   ├── repository/      # Database operations
│   └── services/        # Business logic
├── static/
│   ├── assets/          # Images, Banners, UI elements
│   ├── css/             # Stylesheets
│   └── js/              # Client-side logic
└── templates/           # HTML fragments
```

## Setup & Installation

### Prerequisites
- Go 1.25+
- MongoDB instance

### 1. Environment Configuration
Create a `.env` file in the root directory:
```env
PORT=8000
MONGODB_URI=Nelzya
JWT_SECRET=No
ADMIN_EMAIL=No
```

### 2. Run the Application
```bash
go run cmd/server/main.go
```
The server will start at http://localhost:8000.

## Code Quality
- **Clean Architecture**: Separation of concerns between layers.
- **Optimized Assets**: Localized assets for faster loading and reliability.
- **Sanitized**: Codebase is free of redundant comments and junk files.
