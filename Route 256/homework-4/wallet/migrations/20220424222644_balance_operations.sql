-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table account_operations
(
    id             uuid default gen_random_uuid() not null
                        constraint account_operations_pk primary key,
    account_id     text                           not null,
    amount         integer                        not null,
    operation_id   text                           not null,
    operation_type integer                        not null,

    created_at     timestamp with time zone       not null default (now() at time zone 'utc')
);

create unique index account_operations_operation_id_uindex on account_operations (operation_id);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop index account_operations_operation_id_uindex;
drop table account_operations;
-- +goose StatementEnd
