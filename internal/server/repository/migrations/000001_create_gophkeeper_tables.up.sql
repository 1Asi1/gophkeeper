DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "items";

create table users
(
    id       uuid primary key,
    email text unique not null,
    password bytea          not null
);

create table items
(
    id      uuid primary key,
    user_id uuid,
    type    text not null,
    data    bytea,
    meta    bytea,

    CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users (id) on delete cascade
);
