CREATE TABLE samples
(
    id         SERIAL NOT NULL PRIMARY KEY,
    title      VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
