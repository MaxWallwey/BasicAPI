-- Create a keyspace
CREATE KEYSPACE IF NOT EXISTS store WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

-- Create a table
CREATE TABLE IF NOT EXISTS store.users (
                                                   id uuid PRIMARY KEY,
                                                   name text,
                                                   email_address text,
                                                   last_updated_timestamp timestamp
);

-- Insert some data
INSERT INTO store.users
(id, name, email_address, last_updated_timestamp)
VALUES (uuid(), 'John', 'john@test.com',toTimeStamp(now()));
INSERT INTO store.users
(id, name, email_address, last_updated_timestamp)
VALUES (uuid(), 'Paul', 'paul@test.com',toTimeStamp(now()));