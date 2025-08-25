# 🛒 Kitchen Kirana Tracking System

A small backend service built with **Go + Gin + MongoDB** to manage kitchen/kirana items, track purchases, and generate usage alerts.  
This project demonstrates clean folder structure, MongoDB integration, and input validation.

---

##  Features

- Manage items (Add / List / Update / Delete)
- Track grocery stock levels
- Alerts when items go below threshold
- Record purchases & auto-update stock
- Quick buy option
- Predict consumption trend (basic)

---
## Setup Instructions
### 1. Clone repo
```bash
git clone https://github.com/saurabhraut1212/kitchen_kirana.git
cd kitchen_kirana
```
### 2. Install dependencies
```bash
go mod tidy
```
### 3. Run server
```bash
go run cmd/server/main.go
```
---
## API Endpoints
---
### 1️⃣ Items
| Method | Endpoint         | Description       |
| ------ | ---------------- | ----------------- |
| POST   | `/api/items`     | Add new item      |
| GET    | `/api/items`     | List all items    |
| GET    | `/api/items/:id` | Get single item   |
| PUT    | `/api/items/:id` | Update item       |
| DELETE | `/api/items/:id` | Delete item       |
| GET    | `/api/alerts`    | Items below stock |
---
### 2️⃣ Purchases
| Method | Endpoint                   | Description                    |
| ------ | -------------------------- | ------------------------------ |
| POST   | `/api/purchases`           | Record purchase (update stock) |
| POST   | `/api/purchases/quick/:id` | Quick buy (default +1 qty)     |
---
### 3️⃣ Predict Consumption
| Method | Endpoint                 | Description              |
| ------ | ------------------------ | ------------------------ |
| GET    | `/api/items/:id/predict` | Predict item consumption |
---
## Tech Stack
- Go (Gin) → Web framework
- MongoDB → NoSQL database
- Go Mongo Driver → DB connector
- Validation → binding:"required" on structs

---
## Postman testing
https://web.postman.co/workspace/My-Workspace~388302e8-5eb7-4c3f-821d-5523c39dad56/collection/26119400-bf00e9c2-c4ae-44d5-b968-3c634a5161d6?action=share&source=copy-link&creator=26119400
