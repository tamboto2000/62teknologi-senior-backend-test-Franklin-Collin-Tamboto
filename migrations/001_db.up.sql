CREATE TABLE IF NOT EXISTS businesses (
	id VARCHAR(22) UNIQUE PRIMARY KEY,
	alias VARCHAR(100) NOT NULL,
	name VARCHAR(100) NOT NULL,	
	image_url TEXT NOT NULL,	
	is_closed BOOLEAN NOT NULL,	
	latitude DOUBLE PRECISION NOT NULL,
	longitude DOUBLE PRECISION NOT NULL,
	transactions VARCHAR(22)[] NOT NULL,
	address1 VARCHAR(100) NOT NULL,
	address2 VARCHAR(100),
	address3 VARCHAR(100),
	city_id INT NOT NULL,
	zip_code VARCHAR(10) NOT NULL,
	country_code VARCHAR(5) NOT NULL,
	state_code VARCHAR(5) NULL,
	cross_streets VARCHAR(100),
	phone VARCHAR(20) NOT NULL,
	is_claimed BOOLEAN NOT NULL DEFAULT FALSE,
	date_opened DATE NOT NULL,
	date_closed DATE,	
	message_url TEXT NOT NULL,
	message_use_case TEXT NOT NULL,
	is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS business_hours (
	id SERIAL NOT NULL PRIMARY KEY,
	business_id VARCHAR(22) NOT NULL,
	day INT NOT NULL,
	start_hour VARCHAR(4) NOT NULL,
	end_hour VARCHAR(4) NOT NULL,
	is_overnight BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS business_special_hours (
	id SERIAL NOT NULL PRIMARY KEY,
	business_id VARCHAR(22) NOT NULL,
	day DATE NOT NULL,
	is_closed BOOLEAN NOT NULL DEFAULT FALSE,
	start_hour VARCHAR(4),
	end_hour VARCHAR(4),
	is_overnight BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
	id SERIAL NOT NULL PRIMARY KEY,
	alias VARCHAR(50) NOT NULL,
	title VARCHAR(50) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS cities (
	id SERIAL NOT NULL PRIMARY KEY,
	name VARCHAR(50),
	country_id INT NOT NULL,
	state_id INT,
	latitude DOUBLE PRECISION NOT NULL,
	longitude DOUBLE PRECISION NOT NULL
);

CREATE TABLE IF NOT EXISTS countries (
	id SERIAL NOT NULL PRIMARY KEY,
	name VARCHAR(50),
	code VARCHAR(5),
	latitude DOUBLE PRECISION NOT NULL,
	longitude DOUBLE PRECISION NOT NULL
);

CREATE TABLE IF NOT EXISTS states (
	id SERIAL NOT NULL PRIMARY KEY,
	name VARCHAR(50),
	code VARCHAR(5),
	country_id INT NOT NULL,
	latitude DOUBLE PRECISION NOT NULL,
	longitude DOUBLE PRECISION NOT NULL
);

INSERT INTO categories (alias, title, created_at, updated_at) VALUES 
	('ramen', 'Ramen', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('seafood', 'Seafood', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('japanese', 'Japanese', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);