-- Create a keyspace
CREATE KEYSPACE IF NOT EXISTS store WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

-- Create a table
CREATE TABLE IF NOT EXISTS albums (
                                                   id UUID PRIMARY KEY,
                                                   title text,
                                                   artist text,
                                                   price float,
                                                   stockedSince date
);

CREATE INDEX ON albums(id);

-- Insert some data
INSERT INTO albums
(id, title, artist, price, stockedSince)
VALUES (uuid(), 'Blue Train', 'John Coltrane', 56.99, toTimeStamp(now()));
INSERT INTO albums
(id, title, artist, price, stockedSince)
VALUES (uuid(), 'Jeru', 'Gerry Mulligan', 17.99, toTimeStamp(now()));
INSERT INTO albums
(id, title, artist, price, stockedSince)
VALUES (uuid(), 'Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99, toTimeStamp(now()));