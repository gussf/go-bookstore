\c bookstore

CREATE TABLE IF NOT EXISTS books (
    id SERIAL,
    title  varchar(100) NOT NULL,
    author varchar(100) NOT NULL,
    copies integer NOT NULL,
    price  integer NOT NULL,
    creation_date timestamp
);