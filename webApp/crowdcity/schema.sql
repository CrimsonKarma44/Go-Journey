CREATE TABLE Users
(
    ID       INT PRIMARY KEY AUTO_INCREMENT, -- Automatically increments for each new user
    Username VARCHAR(255) NOT NULL,          -- Stores the username (cannot be NULL)
    Email    VARCHAR(255) NOT NULL UNIQUE,   -- Stores the email (cannot be NULL and must be unique)
    Password VARCHAR(255) NOT NULL           -- Stores the password (cannot be NULL)
);
