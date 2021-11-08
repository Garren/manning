# Go Web Programming (code and exercises)
## by Sau Sheong Chang

https://github.com/sausheong/gwp

```shell
$ sudo - postgres
$ createdb chitchat
```
You can only access a single database per connection. there's no
"use <db>" TSQL analog

Run psql as superuser
`$ sudo -u postgres psql`

Create a user
`$ sudo -u postgres createuser <user>`

Create a db
`$ sudo -u postgres createdb <db>`

```shell
$ sudo -u postgres psql
psql=# alter user <username> with encrypted password '<password>';
psql=# grant all privileges on database <dbname> to <username> ;
```

```postgres
CREATE DATABASE yourdbname;
CREATE USER youruser WITH ENCRYPTED PASSWORD 'yourpass';
GRANT ALL PRIVILEGES ON DATABASE yourdbname TO youruser;
```

```
root@b121136c74c4:/# su - postgres
postgres@b121136c74c4:~$ psql
psql (10.6 (Debian 10.6-1.pgdg90+1))
Type "help" for help.

postgres=# \q
postgres@b121136c74c4:~$ psql postgres://postgres@localhost:5432/gwp
psql (10.6 (Debian 10.6-1.pgdg90+1))
Type "help" for help.

gwp=# create user gwp with encrypted password 'mypass';
CREATE ROLE
gwp=# grant all privileges on database gwp to gwp;
GRANT
```

```shell
root@b121136c74c4:/# psql postgres://gwp:mypass@localhost:5432/gwp
psql (10.6 (Debian 10.6-1.pgdg90+1))
Type "help" for help.

gwp=>
```

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
