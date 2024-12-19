CREATE TABLE person(
    id varchar NOT NULL PRIMARY KEY,
    name varchar NOT NULL,
    age integer NULL,
    active boolean NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);