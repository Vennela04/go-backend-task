🚀 Go Backend Development Task

A simple backend API built using **Go (Golang)**, **Fiber**, and **PostgreSQL**.

This project demonstrates:
- REST API creation in Go using Fiber
- Database integration with PostgreSQL
- Layered architecture (Handler → Service → Repository)
- JSON request handling and response formatting
- Age calculation from date of birth (DOB)
- Tested endpoints via Postman

🧱 Project Structure

go-backend-task/
├── cmd/server/ # Main entry point
├── config/ # Configuration setup (DB & port)
├── db/migrations/ # SQL schema creation
├── internal/ # Application core (handlers, routes, repo, services)
│ ├── handler/
│ ├── logger/
│ ├── repository/
│ ├── routes/
│ └── service/
└── README.md

⚙️ Setup & Run Instructions

1️⃣ Prerequisites
- Go 1.21+  
- PostgreSQL installed and running  
- Postman (for API testing)

2️⃣ Create Database and Table
Open your terminal or pgAdmin and run:

```sql
CREATE DATABASE backenddb;
\c backenddb

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);

3️⃣ Run the Project
Clone the repository:
git clone https://github.com/<your-username>/go-backend-task.git

cd go-backend-task

Install dependencies:
go mod tidy

Run the server:
go run ./cmd/server

Your API will be live on:
http://localhost:8080

🧪 API Endpoints

➕ Create User
POST /users

Request Body:
{
  "name": "Alice",
  "dob": "1990-05-10"
}
Response:
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}

📄 Get All Users
GET /users
Response
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 35
  }
]
🧠 Highlights
✅ Built with Go Fiber
✅ Integrated with PostgreSQL
✅ Clean, modular code structure
✅ Returns dynamic age based on DOB
✅ Tested with Postman

🧰 Tech Stack
Component	Technology
Language	Go (Golang)
Framework	Fiber
Database	PostgreSQL
Logger	    Zap
API Testing	Postman

👩‍💻 Author
Vennela S
💼 Backend Developer | Go | Python | AI/ML
📧 vennela04sanjay@gmail.com
🔗 https://github.com/Vennela04
