### Project Title: Restaurant Reservation System

#### Project Description:
Design and implement a simple restaurant reservation system using Go, the Gin framework, GORM and PostgreSQL. 
The system should be able to handle basic operations such as creating, reading, updating, and deleting reservations. The application should also handle simple user authentication and authorization.

#### Requirements:

##### User Registration and Login: 
Users should be able to register with their email and password. Registered users should be able to login.
Use Gin middleware for handling authentication and session management.

##### Reservations: 
Reservations should at least include the following details:

User ID
Restaurant Name
Reservation Date & Time
Number of Guests

##### Database: 
Use PostgreSQL as the database and GORM for handling ORM. The database should have at least two tables - Users and Reservations. 

Users table should store user details including hashed passwords, and Reservations table should store reservation details.

##### CRUD Operations: 
Implement the following CRUD operations:

###### Creating a new reservation.
Reading existing reservations. Admin users should be able to view all reservations, whereas Regular Users should only be able to view their own reservations.
Updating an existing reservation. Only the user who created a reservation or Admin users should be able to update it.
Deleting a reservation. Only the user who created a reservation or Admin users should be able to delete it.

##### RESTful API: All the above operations should be exposed via RESTful APIs built using the Gin framework.

##### Error Handling: Implement proper error handling. The application should return meaningful error messages and status codes for different types of errors (validation errors, database errors, etc.).

##### Logging: Implement logging using a suitable Go logging library. Log important events such as user login, CRUD operations, errors, etc.