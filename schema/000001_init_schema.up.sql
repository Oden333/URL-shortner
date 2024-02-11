CREATE TABLE url(
        id serial not null unique PRIMARY KEY,
        alias varchar(255) not null UNIQUE,
        url varchar(255) not null UNIQUE ) ;
CREATE INDEX idx_alias ON url(alias);