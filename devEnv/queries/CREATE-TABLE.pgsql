CREATE TABLE users (
    id SERIAL,
    login VARCHAR(64),
    pass VARCHAR(256),
    PRIMARY KEY (id)
);