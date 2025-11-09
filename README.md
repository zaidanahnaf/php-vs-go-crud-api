# php-vs-go-crud-api
A comparative study between PHP (Native) and Go (Gin Framework) in building a simple CRUD API for UMKM (Micro, Small, and Medium Enterprises). This project aims to analyze performance, structure, and best practices in REST API development.


# API CRUD Comparison: PHP vs. Go (Gin Framework)

This documentation covers the implementation of a simple CRUD API using **PHP Native** and **Go (Gin Framework)** to compare performance, code structure, and efficiency in building REST APIs.

---

## 1. Setup & Installation

### A. PHP API Setup

#### 1. Installation & Preparation
- Ensure you have **XAMPP (Apache + MySQL + PHP)** or a server that supports PHP.
- Create a project folder inside `htdocs`.
- Create a MySQL database named `umkm_db` with a `umkm` table.

#### 2. PHP Directory Structure
```sh
/php-api/
 ├── index.php  # API entry point
 ├── conn.php      # Database connection
 ├── users/
     ├── create.php  # Insert UMKM data
     ├── read.php    # Retrieve UMKM data
     ├── update.php  # Update UMKM data
     ├── delete.php  # Delete UMKM data
```

#### 3. PHP API Endpoints
| Method | Endpoint       | Description              |
|--------|--------------|-------------------------|
| GET    | `/umkm`      | Retrieve all UMKM data  |
| POST   | `/umkm`      | Add new UMKM data       |
| PUT    | `/umkm/:id`  | Update existing UMKM    |
| DELETE | `/umkm/:id`  | Delete UMKM data        |

#### 4. Running the PHP API
- Start **Apache & MySQL** in XAMPP.
- Open a browser or Postman and access `http://localhost/php-api/umkm`.

---

### B. Go API Setup (Gin Framework)

#### 1. Installation & Preparation
- Ensure **Go** is installed.
- Install **Gin Framework**:
  ```sh
  go get -u github.com/gin-gonic/gin
  ```
- Install **MySQL driver for Go**:
  ```sh
  go get -u github.com/go-sql-driver/mysql
  ```

#### 2. Go Directory Structure
```sh
/go-api/
 ├── main.go         # API entry point
 ├── config/
 │   ├── conn.go       # Database connection
 ├── handlers/
 │   ├── umkm_handler.go  # CRUD logic
 ├── models/
 │   ├── umkm.go     # UMKM model
 ├── routes/
 │   ├── routes.go   # API routing
```

#### 3. Go API Endpoints
| Method | Endpoint       | Description              |
|--------|--------------|-------------------------|
| GET    | `/umkm`      | Retrieve all UMKM data  |
| POST   | `/umkm`      | Add new UMKM data       |
| PUT    | `/umkm/:id`  | Update existing UMKM    |
| DELETE | `/umkm/:id`  | Delete UMKM data        |

#### 4. Running the Go API
- Run the following command:
  ```sh
  go run main.go
  ```
- Open Postman and access `http://localhost:8080/umkm`.

---

## 2. PHP vs. Go Comparison

| Aspect           | PHP (Native) | Go (Gin Framework) |
|-----------------|-------------|------------------|
| **Performance**  | Slower due to interpreted nature | Faster due to compiled nature |
| **Setup**        | Easy, requires PHP server | Requires Go compiler & dependencies |
| **Code Structure** | Can be messy without a framework | Well-structured with Gin |
| **Error Handling** | Loose, more forgiving | Strict, safer from runtime errors |
| **Scalability**  | Less optimal for large-scale apps | Ideal for microservices |

---

## 3. Conclusion & Recommendations
- **PHP is suitable for quick and simple projects**, especially with **Laravel**.
- **Go is more optimized for high-performance and scalable APIs**, particularly with **Gin Framework**.
- If interested in comparing more technologies, try **Node.js (Express.js)** or **Python (FastAPI)**.

---

## 4. Next Steps
- **To explore PHP further:** Use **Laravel** for a more structured API.
- **To explore Go further:** Implement **middleware and JWT Authentication**.

🔥 **Tried it? If you have questions or want to add more features, feel free to discuss!**
