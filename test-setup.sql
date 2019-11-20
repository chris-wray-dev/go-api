CREATE DATABASE project_test;
\c project_test;

CREATE TABLE locations (
  location_id SERIAL PRIMARY KEY, 
  location_name VARCHAR(30) UNIQUE NOT NULL, 
  location_text TEXT NOT NULL
);

INSERT INTO locations (location_id, location_name, location_text) 
VALUES (1, 'Hacienda', 'This is the Hacienda. Many famous bands have performed here');

INSERT INTO locations (location_id, location_name, location_text) 
VALUES (2, 'The Twisted Wheel', 'This is the The Twisted Wheel. Many famous bands have performed here');

INSERT INTO locations (location_id, location_name, location_text) 
VALUES (3, 'The Free Trade Hall', 'This is the The Free Trade Hall. Many famous bands have performed here');

INSERT INTO locations (location_id, location_name, location_text) 
VALUES (4, 'Another Place', 'This is the Another Place. Many famous bands have performed here');

INSERT INTO locations (location_id, location_name, location_text) 
VALUES (5, 'One More Place', 'This is the One More Place. Many famous bands have performed here');

SELECT * FROM locations