CREATE TABLE blog (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) not null,
    content text not null,
    category VARCHAR(255) not null ,
    tags JSON,
    updatedAt timestamp NOT NULL DEFAULT NOW(),
    createdAt TIMESTAMP not null
);
