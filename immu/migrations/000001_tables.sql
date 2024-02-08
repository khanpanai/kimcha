create table if not exists projects (
    id uuid,
    name varchar[256],
    primary key(id)
);

create table if not exists secret_group (
    id uuid,
    name varchar[256],
    prefix varchar[64],
    project_id uuid,
    primary key(id)
);

create table if not exists access_key (
    id uuid,
    project_id uuid,
    note varchar[2048],
    key varchar[24],
    mask varchar[40],
    signature blob,
    expires timestamp,
    primary key(id)
);