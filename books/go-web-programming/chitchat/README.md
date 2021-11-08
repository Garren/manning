```shell
$ psql postgresql://postgres@localhost:5432/chitchat
```
```psql
chitchat=# create table sessions (
chitchat(#   id         serial primary key,
chitchat(#   uuid       varchar(64) not null unique,
chitchat(#   email      varchar(255),
chitchat(#   user_id    integer references users(id),
chitchat(#   created_at timestamp not null
chitchat(# );
CREATE TABLE
chitchat=#
chitchat=# create table threads (
chitchat(#   id         serial primary key,
chitchat(#   uuid       varchar(64) not null unique,
chitchat(#   topic      text,
chitchat(#   user_id    integer references users(id),
chitchat(#   created_at timestamp not null
chitchat(# );
CREATE TABLE
chitchat=#
chitchat=# create table posts (
chitchat(#   id         serial primary key,
chitchat(#   uuid       varchar(64) not null unique,
chitchat(#   body       text,
chitchat(#   user_id    integer references users(id),
chitchat(#   thread_id  integer references threads(id),
chitchat(#   created_at timestamp not null
chitchat(# );
CREATE TABLE
chitchat=# \dt
          List of relations
 Schema |   Name   | Type  |  Owner
--------+----------+-------+----------
 public | posts    | table | postgres
 public | sessions | table | postgres
 public | threads  | table | postgres
 public | users    | table | postgres
(4 rows)
```

ngrok http --region=us --hostname=garren.ngrok.io 80
