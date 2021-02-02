# dm-challenge

DM challenge project
=======
# DM-Challenge

Delivery Much backend project challenge.

### How to run:
1. After clone the project, just run **docker-compose up -d**
    - This will run the following services:
      - **mongoDB**, the database, on port 27017
      - **mongo-express**, to manage MongoDB, on port 8081
      - **rabbitmq**, the message-broker
        - The service runs on port 5672
        - The UI Management runs on port 15672
      - **stock-service**, service that changes the product stock
      
2. To run the project, you might do:
    - Just run **go run .** at the project root folder
    - Run **go build**, it will generate the binary file called **dm-challenge**, at the project root folder, and then execute this file
