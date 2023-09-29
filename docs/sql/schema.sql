DROP TABLE IF EXISTS users;
CREATE TABLE users (
	id VARCHAR(255) PRIMARY KEY NOT NULL,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	password TEXT NOT NULL,
	profile_image TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	created_by VARCHAR(255) NOT NULL,
	updated_at TIMESTAMP,
	updated_by VARCHAR(255),
	deleted_at TIMESTAMP,
	deleted_by VARCHAR(255),
	is_deleted BOOLEAN NOT NULL
);

DROP TABLE IF EXISTS chat_message;
CREATE TABLE chat_message (
	id VARCHAR(255) PRIMARY KEY NOT NULL,
	message TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	created_by VARCHAR(255) NOT NULL,
	updated_at TIMESTAMP,
	updated_by VARCHAR(255),
	deleted_at TIMESTAMP,
	deleted_by VARCHAR(255),
	is_deleted BOOLEAN NOT NULL
);

DROP TABLE IF EXISTS chat;
CREATE TABLE chat (
	id VARCHAR(255) PRIMARY KEY NOT NULL,
	from_user_id VARCHAR(255) NOT NULL,
	to_user_id VARCHAR(255) NOT NULL,
	message_id VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	created_by VARCHAR(255) NOT NULL,
	updated_at TIMESTAMP,
	updated_by VARCHAR(255),
	deleted_at TIMESTAMP,
	deleted_by VARCHAR(255),
	is_deleted BOOLEAN NOT NULL
);