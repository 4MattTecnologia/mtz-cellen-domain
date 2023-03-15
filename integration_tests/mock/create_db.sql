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
create table stakeholders (
    stakeholder_id serial primary key,
    stakeholder_name varchar(32),
    domain_ids integer[]
);
create table agreements (
    agreement_id serial primary key,
    agreement_name varchar(32),
    num_mtz_users integer,
    num_monitored_users integer,
    page_limit integer
);
create table modules (
    module_id serial primary key,
    module_name varchar(32)
);
create table profiles (
    profile_id serial primary key,
    profile_name varchar(32),
    security jsonb
);
create table mtz_users (
    user_id serial primary key,
    user_name varchar(32),
    password varchar(32),
    domain_id integer,
    stakeholder_id integer,
    profile_id integer,
    start_date date,
    end_date date,
    private_key bytea,
    public_key bytea
);
