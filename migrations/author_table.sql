-- +goose Up
create table if not exists author(
     id serial primary key,
     "name" varchar not null,
     books varchar not null
);

-- +goose Down
drop table if exists author

