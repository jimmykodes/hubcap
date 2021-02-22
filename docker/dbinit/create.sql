create database if not exists vehicles;

use vehicles;

create table if not exists users
(
    id         int auto_increment primary key,
    email      varchar(255) unique not null,
    api_key    varchar(36) unique  not null,
    super_user tinyint(1)          not null
);

create table if not exists services
(
    id              int auto_increment primary key,
    date            int not null,
    odometer        int not null,
    data            text,
    user_id         int not null,
    vehicle_id      int not null,
    service_type_id int not null
);

create table if not exists service_types
(
    id         int auto_increment primary key,
    name       varchar(255) not null,
    freq_miles int,
    freq_days  int,
    user_id    int          not null
);

create table if not exists vehicles
(
    id      int auto_increment primary key,
    name    varchar(255) not null,
    make    varchar(255) not null,
    mode    varchar(255) not null,
    year    int          not null,
    user_id int          not null
);
