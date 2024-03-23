CREATE DATABASE whatsup
CREATE TABLE users (
    id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	email VARCHAR ( 50 ) UNIQUE NOT NULL
)

CREATE TABLE connection_req (
    id serial PRIMARY KEY,
	validator int NOT NULL,
	user_id INT REFERENCES users(id),
	req_time TIMESTAMP

)