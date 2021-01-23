drop table if exists auth;
drop table if exists users;

create table users
(
    user_id int not null auto_increment,
    login varchar(40) not null,
    first_name varchar(40) null,
    last_name varchar(40) null,
    birth_date date null,
    job_title varchar(100) null,
    city varchar(50) null,
    primary key (user_id)
);

create table auth
(
    id int not null auto_increment,
    user_id int,
    token varchar(100) not null,
    primary key(id),
    foreign key (user_id) references users(user_id) on delete cascade
);

show tables;