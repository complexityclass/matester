drop table if exists auth;
drop table if exists friends;
drop table if exists user2hobbies;
drop table if exists users;
drop table if exists hobbies;

create table users
(
    user_id int not null auto_increment,
    login varchar(40) not null,
    first_name varchar(40) null,
    last_name varchar(40) null,
    birth_date date default null,
    gender enum('male', 'female', 'other') default null,
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

create table friends 
(
    id int not null auto_increment,
    fst int,
    snd int,
    primary key(id),
    foreign key (fst) references users(user_id) on delete cascade,
    foreign key (snd) references users(user_id) on delete cascade
);

create table hobbies
(
    id int not null auto_increment,
    name varchar(100) not null,
    primary key(id)
);

create table user2hobbies 
(
    id int not null auto_increment,
    user_id int,
    hobby_id int,
    foreign key(user_id) references users(user_id) on delete cascade,
    foreign key(hobby_id) references hobbies(id) on delete cascade,
    primary key(id)
);

show tables;