CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO users(id, username, encrypted_password, nickname) VALUES (uuid_generate_v4(), 'first', crypt('first', gen_salt('bf', 8)), 'first');
INSERT INTO users(id, username, encrypted_password, nickname) VALUES (uuid_generate_v4(), 'second', crypt('second', gen_salt('bf', 8)), 'second');
INSERT INTO users(id, username, encrypted_password, nickname) VALUES (uuid_generate_v4(), 'third', crypt('third', gen_salt('bf', 8)), 'third');
INSERT INTO posts(id, name, content) VALUES (uuid_generate_v4(), 'first', 'first');
INSERT INTO posts(id, name, content) VALUES (uuid_generate_v4(), 'second', 'second');
INSERT INTO posts(id, name, content) VALUES (uuid_generate_v4(), 'third', 'third');