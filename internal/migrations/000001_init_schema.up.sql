CREATE TABLE result (
    id SERIAL PRIMARY KEY,
    gender VARCHAR(10) NOT NULL,
    title VARCHAR(10) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    street_number INTEGER NOT NULL,
    street_name VARCHAR(100) NOT NULL,
    city VARCHAR(50) NOT NULL,
    state VARCHAR(50) NOT NULL,
    country VARCHAR(50) NOT NULL,
    postcode VARCHAR(20) NOT NULL,
    latitude DECIMAL(9,6) NOT NULL,
    longitude DECIMAL(9,6) NOT NULL,
    timezone_offset VARCHAR(10) NOT NULL,
    timezone_description VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    salt VARCHAR(50) NOT NULL,
    md5 VARCHAR(50) NOT NULL,
    sha1 VARCHAR(50) NOT NULL,
    sha256 VARCHAR(50) NOT NULL,
    dob DATE NOT NULL,
    age INTEGER NOT NULL,
    registered_date DATE NOT NULL,
    registered_age INTEGER NOT NULL,
    phone VARCHAR(20) NOT NULL,
    cell VARCHAR(20) NOT NULL,
    national_id_name VARCHAR(50) NOT NULL,
    national_id_value VARCHAR(20) NOT NULL,
    picture_large_url VARCHAR(200) NOT NULL,
    picture_medium_url VARCHAR(200) NOT NULL,
    picture_thumbnail_url VARCHAR(200) NOT NULL,
    nationality VARCHAR(10) NOT NULL
);

CREATE TABLE info (
    id SERIAL PRIMARY KEY,
    seed VARCHAR(50),
    results INTEGER,
    page INTEGER,
    version VARCHAR(10)
);

CREATE TABLE result_into (
    id SERIAL PRIMARY KEY,
    result_id INTEGER REFERENCES result(id),
    info_id INTEGER REFERENCES info(id)
)