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

🧪 How to Test

You can test the API using Postman, Insomnia, or any other API client of your choice.

🔍 Get User Balance
GET http://localhost:8080/user/2/balance

Sample Response:

{
  "userId": 2,
  "balance": "2000.00",
  "createdAt": "2025-06-25 06:36:37",
  "updatedAt": "2025-06-25 06:39:58"
}

💸 Create a Transaction
POST http://localhost:8080/user/2/transaction

Source-Type: game  // Options: game, server, payment, etc.
Content-Type: application/json

{
  "state": "win",         // Options: "win" or "lost"
  "amount": "100.00",
  "transactionId": "trx-1"  // Must be unique for each request (idempotency)
}

⚠️ Note: transactionId must always be unique due to idempotency. Reusing the same ID will not create a new transaction.

📈 Load Testing

Use the hey CLI tool to test how many requests per second your service can handle:

hey -n 1000 -c 30 -m POST \
  -H "Content-Type: application/json" \
  -H "Source-Type: game" \
  -d '{"state":"win", "amount":"10.00", "transactionId":"some-id-123"}' \
  http://localhost:8080/user/1/transaction

Sample Response:

Status code distribution:
  [200] 1 responses
  [409] 985 responses
  [500] 4 responses

This output shows that only one request succeeded. The remaining 409 errors are due to repeated use of the same transactionId, as the system ensures idempotency by allowing only one transaction per unique ID.



## See more commands in Makefile 

