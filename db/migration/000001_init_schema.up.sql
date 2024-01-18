CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE customer(
	_id uuid DEFAULT uuid_generate_v4(),
    external_id varchar(50) NOT NULL DEFAULT '',
    first_name varchar(100) NOT NULL DEFAULT '',
    middle_name varchar(100) NOT NULL DEFAULT '',
    last_name varchar(100) NOT NULL DEFAULT '',
    legal_name text NOT NULL,
    email text NOT NULL DEFAULT '',
    phone_number varchar(20),
	PRIMARY KEY (_id)
);

CREATE TABLE plan(
    _id uuid DEFAULT uuid_generate_v4(),
    external_id varchar(50) NOT NULL DEFAULT '',
    plan_name varchar(100) NOT NULL DEFAULT '',
    plan_code varchar(50) NOT NULL DEFAULT '',
    PRIMARY KEY (_id)
);