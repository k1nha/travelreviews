CREATE TABLE IF NOT EXISTS reviews 
(
  id uuid PRIMARY KEY NOT NULL,
  place_id uuid NOT NULL,
  stars INT NOT NULL,
  description varchar(255) NOT NULL,
  created_at datetime NOT NULL
)
