-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

create table "users"
(
    id             uuid default gen_random_uuid() not null primary key,
    name                  text                           not null,
    lastname              text                           not null,
    age                   integer                        not null,
    phone                 text                           not null,
    password_hash         text                        not null,
    level                 text                        not null,

    created_at     timestamp with time zone       not null default (now() at time zone 'utc')
);

create UNIQUE INDEX user_phone_idx ON "users" (phone);

create table "sessions"
(
    user_id                text not NULL,
    token                  text                           not null,
    created_at     timestamp with time zone       not null default (now() at time zone 'utc'),
    PRIMARY KEY (user_id, token)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

drop INDEX IF EXISTS user_phone_idx;
drop table "users";

drop table "sessions";

-- +goose StatementEnd
