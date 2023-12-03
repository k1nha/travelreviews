CREATE TABLE IF NOT EXISTS reviews 
(
  id uuid PRIMARY KEY NOT NULL,
  place_id uuid NOT NULL REFERENCES places(id),
  stars INT NOT NULL,
  description VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL
)


CREATE TABLE IF NOT EXISTS places 
(
  id uuid PRIMARY KEY NOT NULL,
  name varchar(255) NOT NULL,
  address varchar(255) NOT NULL,
  city varchar(255) NOT NULL,
  created_at DATETIME NOT NULL
)