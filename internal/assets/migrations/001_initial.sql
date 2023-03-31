-- +migrate Up

create type telegram_access_levels_enum as enum ('owner', 'admin', 'member', 'self', 'left', 'banned');

create table if not exists responses (
    id uuid primary key,
    status text not null,
    error text,
    payload jsonb,
    created_at timestamp without time zone not null default current_timestamp
);

create table if not exists users (
    id bigint unique,
    username text unique,
    telegram_id bigint primary key,
    access_hash bigint not null,
    first_name text not null,
    last_name text not null,
    phone text unique,
    updated_at timestamp with time zone not null default current_timestamp,
    created_at timestamp with time zone default current_timestamp
);

create index if not exists users_id_idx on users(id);
create index if not exists users_username_idx on users(username);
create index if not exists users_telegramid_idx on users(telegram_id);

create table if not exists links (
    id serial primary key,
    link text not null,
    unique(link)
);
insert into links (link) values ('HELP TG API');
insert into links (link) values ('WE vs. ACS');
insert into links (link) values ('Messenger Internal');
insert into links (link) values ('DL / Make TokenE even better');

create index if not exists links_link_idx on links(link);

create table if not exists permissions (
    request_id text not null,
    telegram_id int not null,
    link text not null,
    access_level telegram_access_levels_enum not null,
    created_at timestamp without time zone not null,
    updated_at timestamp without time zone not null default current_timestamp,

    unique (telegram_id, link),
    foreign key(telegram_id) references users(telegram_id) on delete cascade on update cascade,
    foreign key(link) references links(link) on delete cascade on update cascade
);

create index if not exists permissions_telegramid_idx on permissions(telegram_id, link);
create index if not exists permissions_link_idx on permissions(telegram_id, link);

-- +migrate Down

drop index if exists permissions_telegramid_idx;
drop index if exists permissions_link_idx;

drop table if exists permissions;

drop index if exists links_link_idx;

drop table if exists links;

drop index if exists users_id_idx;
drop index if exists users_username_idx;
drop index if exists users_telegramid_idx;

drop table if exists users;
drop table if exists responses;

drop type if exists telegram_access_levels_enum;