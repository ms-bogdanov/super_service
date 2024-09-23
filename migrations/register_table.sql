-- +goose Up
create table if not exists register
(
    id          bigserial primary key,
    user_id     bigint      not null references users (id),
    book_id     bigint      not null references books (id),
    took_at     timestamptz not null default now(),
    return_at   timestamptz
);

-- +goose Down
drop table if exists register;