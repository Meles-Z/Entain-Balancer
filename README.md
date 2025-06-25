# Entain-Balancer

👋 Hi Technical Team,

Thank you for taking the time to review my project.

I’ve designed this project using a modular architecture to ensure high scalability, ease of maintenance, and smooth integration of new features. I'm eager to hear your feedback on what I've done well and where I can improve.

---

#### Below is the best way to run this project in your development environment

Before proceeding, I recommend that you check the Makefile.

## How to Run the Project

1. Build the project  
   This will pull the required Docker images (`golang:1.24-alpine`, `postgres:17`), build the Go application, and start containers in detached mode.

   ``make build``

2. Populate the database  
   This creates predefined users (IDs 1, 2, and 3). If they already exist, they will be skipped.

   ``make populate_db``


3. Send API requests  
   The backend server runs at:

   ``http://localhost:8080``

   You can now send requests and receive responses from the available endpoints.

## ✅ Optionals

1. Start containers without rebuilding (optional)  
   Use this command if you’ve already built the project and only want to start the containers.

   ``make run``

2. View logs (optional)  
   Shows real-time logs from your running containers.

   ``make logs``

3. Wipe the database (optional)  
   Clears all specified tables data.

   ``make wipe_db``


---

## See more commands in Makefile 

