# Entain-Balancer

Hi all 👋

This project is structured for high scalability and maintainability.

Below is the best way to run this project in your development environment:

---

## ✅ Requirements

1. You must have **Docker** installed on your machine.

---

## 🚀 How to Run the Project

1. Start Docker  
   Make sure Docker is running before you proceed.

2. Build the project  
   This will pull the required Docker images (`golang:1.24-alpine`, `postgres:17`), build the Go application, and start containers in detached mode.

   ``make build``

3. Populate the database  
   This creates predefined users (IDs 1, 2, and 3). If they already exist, they will be skipped.

   ``make populate_db``

4. Start containers without rebuilding (optional)  
   Use this command if you’ve already built the project and only want to start the containers.

   ``make up``

5. View logs (optional)  
   Shows real-time logs from your running containers.

   ``make logs``

6. Wipe the database (optional)  
   Clears all tables and resets the database schema. You will need to rebuild afterward.

   ``make wipe_db``

   After that:

   ``make build``

7. Send API requests  
   The backend server runs at:

   ``http://localhost:8080``

   You can now send requests and receive responses from the available endpoints.

---

## ✅ Summary

- Dockerized Go application
- PostgreSQL database with GORM
- Structured logging with Zap
- Clean and scalable folder architecture
- Makefile for automated workflow

---
