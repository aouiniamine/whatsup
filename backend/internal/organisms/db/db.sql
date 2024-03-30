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

CREATE TABLE messages(
   id serial PRIMARY KEY,
   sender INT,
   recipient INT,
   message_content VARCHAR(255) NOT NULL,
   CONSTRAINT fk_sender
      FOREIGN KEY(sender) 
	    REFERENCES users(id)
	        ON DELETE SET NULL,
   CONSTRAINT fk_recipient
      FOREIGN KEY(sender) 
	    REFERENCES users(id)
	        ON DELETE SET NULL
)