create table invites (
    id integer primary key not null,
    name varchar
);
create table guests (
    id integer primary key not null,
    name varchar not null,
    invite_id bigint,
    foreign key (invite_id) references invites (id)
);