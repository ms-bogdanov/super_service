-- +goose Up
create table if not exists books(
     id serial primary key,
     "title" varchar not null,
     author_id integer not null
);

-- +goose Down
drop table if exists books
