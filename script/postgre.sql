create schema job authorization postgres;

create table job.users
(
    id       bigserial primary key,
    username text unique not null,
    balance  numeric     not null default 0 check ( balance >= 0)
);

create table job.transactions
(
    id        bigserial primary key,
    from_user bigint    references job.users (id) on delete set null,
    to_user   bigint    references job.users (id) on delete set null,
    created   timestamp not null default current_timestamp,
    sum       numeric   not null,
    comment   text
);

create or replace function job.update_user_balance() returns trigger
    language plpgsql
as
$update_user_balance$
begin
    update job.users
    set balance = balance - new.sum
    where id = new.from_user;

    update job.users
    set balance = balance + new.sum
    where id = new.to_user;

    return new;
end;
$update_user_balance$;

create trigger update_user_balance_trigger
    after insert
    on job.transactions
    for each row
execute procedure job.update_user_balance();
