CREATE TABLE buys (
    id SERIAL NOT NULL PRIMARY KEY,
    idUser INTEGER NOT NULL REFERENCES users(id),
    createdat TIMESTAMP NOT NULL
);