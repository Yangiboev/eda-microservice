create table if not exists companies
(
    id             uuid primary key,
    name           varchar unique not null,
    telephone      varchar(30) not null,
    email          varchar not null,
    mfo            integer not null,
    inn            integer not null,
    account_number integer not null,
    address        text not null,
    description    text,
    created_at     timestamp default current_timestamp,
    updated_at     timestamp,
    deleted_at     timestamp
);
create unique index companies_u1 on companies (name);
create index companies_i1 on companies (deleted_at);

create table if not exists cities (
                                      id uuid primary key,
                                      name varchar(50) not null
);
create table if not exists branches
(
    id                  uuid primary key,
    company_id          uuid not null references companies(id),
    city_id             uuid not null references cities(id),
    name                varchar unique not null,
    telephone           integer not null,
    number_of_employees integer not null,
    address             text not null,
    description         text,
    created_at          timestamp default current_timestamp,
    updated_at          timestamp,
    deleted_at          timestamp
);
create unique index branches_u1 on branches(name);
create index branches_i1 on branches(deleted_at);


create table if not exists counter_agents
(
    id             uuid primary key,
    name           varchar unique not null,
    telephone      varchar(30) not null,
    address        text not null,
    description    text,
    created_at     timestamp default current_timestamp,
    updated_at     timestamp,
    deleted_at     timestamp
);
create unique index counter_agents_u1 on counter_agents(name);
create index counter_agents_i1 on counter_agents(deleted_at);

create table if not exists legal_counter_agents
(
    counter_agent_id uuid,
    inn              integer not null,
    mfo              integer,
    account_number   integer,
    primary key (counter_agent_id),
    constraint fk_counter_agent
        foreign key (counter_agent_id)
            references counter_agents(id)
            on delete cascade
);