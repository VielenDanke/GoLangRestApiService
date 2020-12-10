CREATE TABLE IF NOT EXISTS users (
    id varchar(128) not null primary key,
    username varchar(128) not null unique,
    encrypted_password varchar(512) not null,
    nickname varchar(128) not null unique
);

CREATE TABLE IF NOT EXISTS posts (
    id varchar(128) not null primary key,
    name varchar(256) not null,
    content varchar(256) not null
);