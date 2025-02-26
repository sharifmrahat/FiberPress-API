# FiberPress-API

FiberPress is a simple and efficient **blog API** built using **Golang** with the Fiber framework and MongoDB as the database.

## **API URL**

```
Base URL:
```

## **Tech Stack**

- **Golang** (Fiber Framework)
- **MongoDB** (NoSQL Database)
- **JWT Authentication** (for secure access)
- **Godotenv** (for environment variable management)

## **Setup Instructions**

### **1. Clone the repository**

```sh
git clone https://github.com/sharifmrahat/FiberPress-API.git
cd FiberPress-API
```

### **2. Install dependencies**

```sh
go mod tidy
```

### **3. Set up environment variables**

Create a `.env` file in the root directory and add the following:

```
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=fiberpress
SERVER_PORT=3000
JWT_SECRET=supersecretkey
```

### **4. Run the application**

```sh
go run main.go
```

## **Endpoints**

### **Authentication**

| Method | Endpoint    | Description           | Auth Required |
| ------ | ----------- | --------------------- | ------------- |
| POST   | `/register` | Register a new user   | ❌ No         |
| POST   | `/login`    | Login and get a token | ❌ No         |

### **User Management**

| Method | Endpoint     | Description      | Auth Required |
| ------ | ------------ | ---------------- | ------------- |
| GET    | `/users/:id` | Get user profile | ✅ Yes        |
| PUT    | `/users/:id` | Update profile   | ✅ Yes        |

### **Posts**

| Method | Endpoint     | Description             | Auth Required |
| ------ | ------------ | ----------------------- | ------------- |
| GET    | `/posts`     | Get all published posts | ❌ No         |
| POST   | `/posts`     | Create a new post       | ✅ Yes        |
| GET    | `/posts/:id` | Get post by ID          | ❌ No         |
| PUT    | `/posts/:id` | Update a post           | ✅ Yes        |
| DELETE | `/posts/:id` | Soft delete a post      | ✅ Yes        |

### **Categories**

| Method | Endpoint      | Description           | Auth Required |
| ------ | ------------- | --------------------- | ------------- |
| GET    | `/categories` | Get all categories    | ❌ No         |
| POST   | `/categories` | Create a new category | ✅ Yes        |

---

## **Contributing**

1. Fork the repository
2. Create a new branch (`feature-branch`)
3. Commit your changes (`git commit -m "Added new feature"`)
4. Push the branch (`git push origin feature-branch`)
5. Create a Pull Request

---

## **License**

This project is licensed under the MIT License.
