create extension if not exists citext;

CREATE TABLE IF NOT EXISTS(
                              id  bigserial PRIMARY KEY,
                              email  citext unique NOT NULL,
                              username varchar(255) unique NOT NULL,
    password bytea not null,
    created_at timestamp(0) with time zone NOT NULL default NOW()






    )