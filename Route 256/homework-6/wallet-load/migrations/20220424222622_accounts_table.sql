-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE OR REPLACE FUNCTION set_updated_at_column() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now() at time zone 'utc';
    RETURN NEW;
END;
$$ language 'plpgsql';

create table accounts
(
    id          uuid default gen_random_uuid()   not null
                    constraint accounts_pk primary key,
    user_id uuid not null,
    account_id  text                            not null,
    amount      integer                         not null,
    description text,

    created_at  timestamp with time zone        not null default (now() at time zone 'utc'),
    updated_at  timestamp with time zone        not null default (now() at time zone 'utc')
);
create trigger update_account_updated_at before update on accounts for each row execute procedure set_updated_at_column();
create unique index accounts_account_id_uindex on accounts (user_id, account_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table accounts;
drop function set_updated_at_column;
-- +goose StatementEnd
