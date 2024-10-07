# Synapsis Restaurant API

This API was developed by **Pramudya Wibisono** as part of the **Back End Engineer Challenge - PT Synapsis Sinergi Digital**. It provides a backend solution for managing restaurant operations such as user registration, login, food listing, and managing food carts.

## API Documentation

The API is structured with public and private routes. Below is an overview of the available routes:

### Public Routes

These routes are accessible without authentication:

- **POST** `/api/v1/public/register`  
  Allow users to register account.

- **POST** `/api/v1/public/login`  
  Allow users to login and receive JWT token.

- **GET** `/api/v1/public/foods`  
  Returns a list of available foods in the restaurant.

### Private Routes

These routes require authentication using a Bearer token in the Authorization header:

- **POST** `/api/v1/private/food-carts`  
  Add a food item to the user's cart.

- **DELETE** `/api/v1/private/food-carts/:id`  
  Remove a specific item from the food cart by its ID.

- **GET** `/api/v1/private/food-carts`  
  Retrieve the current list of items in the user's food cart.

- **POST** `/api/v1/private/food-carts/checkout`  
  Checkout all items in the user's cart and create an order.

### Index Route

- **GET** `/api/v1/`  
  Returns a welcome message:  
  ```json
  {
    "message": "Welcome to Synapsis Restaurant API"
  }

## Postman Collection

- To try out the API, you can use the Postman collection available at the following link:  
[Synapsis Restaurant API Postman Collection](https://github.com/pramudyaws/synapsis-restaurant/blob/main/Synapsis%20Restaurant%20API.postman_collection.json)

- For private routes, ensure that you replace the Authorization header with the correct Bearer token, which is obtained after successful login.

## Technologies Used
- **Go (Golang)**: Backend programming language
- **Gin**: HTTP web framework for building the API
- **GORM**: Object Relational Mapping (ORM) library for database interactions
- **JWT**: JSON Web Tokens for authentication and authorization
- **Supabase (PostgreSQL)**: Used for database management, powered by PostgreSQL via Supabase
- **Docker**: For containerization
- **AWS**: Amazon Web Services used for deployment
