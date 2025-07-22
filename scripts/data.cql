-- Create a keyspace
CREATE KEYSPACE IF NOT EXISTS store WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

-- Create a table
CREATE TABLE IF NOT EXISTS store.albums (
                                                   id text PRIMARY KEY,
                                                   title text,
                                                   artist text,
                                                   price float,
                                                   stockedSince date
);

CREATE INDEX ON store.albums(id);

-- Insert some data
INSERT INTO store.albums
(id, title, artist, price, stockedSince)
VALUES ('1', 'Blue Train', 'John Coltrane', 56.99, toTimeStamp(now()));
INSERT INTO store.albums
(id, title, artist, price, stockedSince)
VALUES ('2', 'Jeru', 'Gerry Mulligan', 17.99, toTimeStamp(now()));
INSERT INTO store.albums
(id, title, artist, price, stockedSince)
VALUES ('3', 'Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99, toTimeStamp(now()));