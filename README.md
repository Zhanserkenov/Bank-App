## Project Description

This project is a simple banking application API built in Go (Golang). It provides basic functionalities for user registration, login, account management, and transactions. Below is an overview of the project structure and functionalities:

## Database Structure and Relationships

The database structure for the Bank App consists of the following entities:

1. **User**: Represents a registered user of the bank application.
   - Attributes:
     - ID (Primary Key)
     - Username
     - Email
     - Password

2. **Account**: Represents a bank account associated with a user.
   - Attributes:
     - ID (Primary Key)
     - Type
     - Name
     - Balance
     - UserID (Foreign Key referencing User.ID)

3. **Transaction**: Represents a financial transaction between accounts.
   - Attributes:
     - ID (Primary Key)
     - From (Foreign Key referencing Account.ID)
     - To (Foreign Key referencing Account.ID)
     - Amount

## API Structure

The Bank App API provides the following endpoints:

1. **POST /login**: Endpoint for user authentication.
   - Request Body:
     - Username
     - Password
   - Response:
     - JWT token on successful authentication.
     - Error message if authentication fails.

2. **POST /register**: Endpoint for user registration.
   - Request Body:
     - Username
     - Email
     - Password
   - Response:
     - JWT token on successful registration and authentication.
     - Error message if registration fails.

3. **POST /transaction**: Endpoint for conducting transactions between accounts.
   - Request Body:
     - UserID
     - From (Account ID)
     - To (Account ID)
     - Amount
   - Response:
     - Success message and updated account details on successful transaction.
     - Error message if transaction fails.

4. **GET /transactions/{userID}**: Endpoint for retrieving transactions associated with a user.
   - Request Parameter:
     - userID
   - Response:
     - List of transactions associated with the user.
     - Error message if user or transactions are not found.

5. **GET /user/{id}**: Endpoint for retrieving user details.
   - Request Parameter:
     - id
   - Response:
     - User details including associated accounts.
     - Error message if user is not found or token is invalid.

## Team
1. **Zhumatay Sayana** - **22B030367**
2. **Zhakyayev Islambek** - **22B031315**
3. **Zhanserkeno Arsen** - **22B031492**
