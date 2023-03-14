create database mock_mattzero;
\c mock_mattzero
create table domains (
    domain_id serial primary key,
    domain_name varchar(32)
);

