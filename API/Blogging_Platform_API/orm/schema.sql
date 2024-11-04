CREATE TABLE blog (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) not null,
    content text,
    category VARCHAR(255),
    tags JSON,
    createdAt TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT NOW()
);
