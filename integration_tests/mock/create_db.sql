create database mock_mattzero;
\c mock_mattzero
create table domains (
    domain_id serial primary key,
    domain_name varchar(32)
);
create table origins (
    origin_id serial primary key,
    origin_name varchar(32),
    connection_info jsonb
);
create table origin_instances (
    origin_instance_id serial primary key,
    origin_instance_name varchar(32),
    origin_id integer,
    domain_id integer,
    connection_values jsonb
);
