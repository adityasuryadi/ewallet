CREATE TABLE "users" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255),
    address VARCHAR(255),
    dob DATE,
    role VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
