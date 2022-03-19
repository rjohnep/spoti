DROP TABLE IF EXISTS playlists;
CREATE TABLE IF NOT EXISTS playlists (
  id         serial PRIMARY KEY,
  title      VARCHAR (128) UNIQUE NOT NULL,
  artist     VARCHAR (255) NOT NULL,
  price      DECIMAL (5,2) NOT NULL
);


INSERT INTO playlists
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
