create table if not exists users
(
    id         serial primary key,
    username   varchar(255) unique not null,
    api_key    varchar(36) unique  not null,
    super_user boolean             not null
);
create table if not exists sessions
(
    key     varchar(36),
    user_id int not null,
    expires int not null
);

create table if not exists services
(
    id              serial primary key,
    date            int not null,
    odometer        int not null,
    data            json,
    user_id         int not null,
    vehicle_id      int not null,
    service_type_id int not null
);

create table if not exists service_types
(
    id         serial primary key,
    name       varchar(255) not null,
    freq_miles int,
    freq_days  int,
    questions  json,
    user_id    int          not null
);

create table if not exists vehicles
(
    id      serial primary key,
    name    varchar(255) not null,
    make    varchar(255) not null,
    model   varchar(255) not null,
    year    int          not null,
    user_id int          not null
);
