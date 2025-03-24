# Payment Processing System

## Overview
The **Payment Processing System** is a backend microservices-based project that enables secure payment transactions using **Golang, Gin, PostgreSQL, Kafka, and Stripe**. It follows a **scalable and modular architecture**, ensuring high availability and reliability.

## Tech Stack
- **Golang** – Backend development
- **Gin** – Web framework for API handling
- **PostgreSQL** – Database for storing user and transaction details
- **Kafka** – Event-driven messaging system
- **Stripe API** – Payment processing
- **Docker** – Containerization
- **Kubernetes (future)** – For orchestration and deployment

## Services
1. **User Service** – Handles authentication, authorization, and user management.
2. **Payment Service** – Processes payments and integrates with Stripe.
3. **Transaction Service (Upcoming)** – Manages transaction records and history.
4. **Notification Service (Upcoming)** – Sends real-time notifications to users.

## Current Features
- ✅ **User Authentication & Authorization** (JWT-based Sign-Up & Sign-In)
- ✅ **Payment Integration with Stripe** (Initial implementation using Payment Intents)
- ✅ **Kafka Producer** (Publishes payment events)
- ✅ **Webhook for Stripe** (Processes payment updates)

## Future Implementations (In Progress)
- 🔄 **Kafka Consumer** – Listen to payment events and update the database.
- 🔄 **Transaction Logging** – Store transaction records for reconciliation.
- 🔄 **Retry Mechanism** – Implement **Kafka Dead Letter Queue (DLQ)** for failed payments.
- 🔄 **Kubernetes Deployment** – Deploy services in a scalable environment.
- 🔄 **Custom Payment Flow** – Transition from Stripe to an in-house payment gateway.

## Setup Instructions
### Prerequisites
- Install **Go (v1.20+)**
- Install **Docker & Docker Compose**
- Install **PostgreSQL**
- Install **Kafka** (Locally or via Docker)

### Running the Services
```sh
# Clone the repository
git clone https://github.com/Rahul13900/Payment-Processing-System.git
cd payment-processing-system

# Start PostgreSQL and Kafka (Docker recommended)
docker-compose up -d

# Run User Service
go run user-service/main.go

# Run Payment Service
go run payment-service/main.go
```

### Testing Webhooks Locally
```sh
stripe listen --forward-to http://localhost:8080/webhook/stripe
```

## API Endpoints
### User Service
- `POST /signup` – Register a new user
- `POST /signin` – Authenticate user and return JWT

### Payment Service
- `POST /payments` – Initiate a payment request
- `POST /webhook/stripe` – Listen for Stripe events


