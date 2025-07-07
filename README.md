# TakeHomeProject – Setup Instructions

This project includes a Golang backend using Gin and a ReactJS frontend, backed by a MySQL database.

## Prerequisites

- Go 1.21 or later
- Node.js 18 or later
- MySQL Server (e.g., MySQL 8.0+)

---

## Backend Setup (Go + MySQL)

1. Clone the repository and navigate to the backend directory:

   ```bash
   git clone https://github.com/bobo-1232/TakeHomeProject.git
   cd TakeHomeProject/backend
   ```

2. Set up the MySQL database:

   - Create a database named `golangdb`
   - Use the provided SQL script to create and seed the tables:

   ```bash
   mysql -u your_username -p golangdb < init.sql
   ```

   Replace `your_username` with your MySQL username. You’ll be prompted for your password.

3. Update the database connection string in `main.go`:

   ```go
   dsn := "username:password@tcp(127.0.0.1:3306)/golangdb"
   ```

   Replace `username` and `password` with your local MySQL credentials.

4. Run the backend server:

   ```bash
   go mod tidy
   go run main.go
   ```

   The API will now be available at: `http://localhost:8080`

---

## Frontend Setup (React)

1. Navigate to the frontend directory:

   ```bash
   cd ../frontend/go-learning
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Start the frontend development server:

   ```bash
   npm start
   ```

   The React app will be available at: `http://localhost:3000`

---

## Test the Setup

- Verify backend: open [http://localhost:8080/person/1/info](http://localhost:8080/person/1/info)
- Verify frontend: open [http://localhost:3000](http://localhost:3000)

---

## Test POST /person/create (Task A)

You can test the `POST /person/create` endpoint using a tool like **Postman** or **cURL**:

### Request

**POST** `http://localhost:8080/person/create`  
**Content-Type:** `application/json`

### Example JSON Body

```json
{
  "name": "Sarah",
  "phone_number": "555-123-4567",
  "city": "Denver",
  "state": "CO",
  "street1": "789 Broadway",
  "street2": "Unit 5",
  "zip_code": "80203"
}
```

### Expected Response

- HTTP `200 OK` on success
- Check the database to confirm the new person and related entries were inserted
