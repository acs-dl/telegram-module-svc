-- +migrate Up

create type telegram_access_levels_enum as enum ('owner', 'admin', 'member', 'self', 'left', 'banned');

create table if not exists responses (
    id uuid primary key,
    status text not null,
    error text,
    description text,
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
    created_at timestamp with time zone not null default current_timestamp
);

create index if not exists users_id_idx on users(id);
create index if not exists users_username_idx on users(username);
create index if not exists users_telegramid_idx on users(telegram_id);

create table if not exists chats (
    title text not null,
    id bigint not null,
    access_hash bigint,
    members_amount bigint not null,
    photo_name text,
    photo_link text,

    unique(id, access_hash)
);

create index if not exists chats_access_hash_idx on chats(access_hash);
create index if not exists chats_id_idx on chats(id);

create table if not exists links (
    id serial primary key,
    link text not null,
    unique(link)
);

create index if not exists links_link_idx on links(link);

create table if not exists permissions (
    request_id text not null,
    telegram_id bigint not null,
    link text not null,
    access_level telegram_access_levels_enum not null,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    submodule_id bigint not null,
    submodule_access_hash bigint,

    unique (telegram_id, submodule_id, submodule_access_hash),
    foreign key(telegram_id) references users(telegram_id) on delete cascade on update cascade,
    foreign key(link) references links(link) on delete cascade on update cascade
);

create index if not exists permissions_telegramid_idx on permissions(telegram_id);
create index if not exists permissions_link_idx on permissions(link);
create index if not exists permissions_submodule_id_idx on permissions(submodule_id);
create index if not exists permissions_submodule_access_hash_idx on permissions(submodule_access_hash);

-- +migrate Down

drop index if exists permissions_telegramid_idx;
drop index if exists permissions_link_idx;
drop index if exists permissions_submodule_id_idx;
drop index if exists permissions_submodule_access_hash_idx;

drop table if exists permissions;

drop index if exists links_link_idx;

drop table if exists chats;
drop index if exists chats_access_hash_idx;
drop index if exists chats_id_idx;

drop table if exists links;

drop index if exists users_id_idx;
drop index if exists users_username_idx;
drop index if exists users_telegramid_idx;

drop table if exists users;
drop table if exists responses;

drop type if exists telegram_access_levels_enum;