create table user(user_id int not null auto_increment, login varchar(40) not null, primary key (user_id));

create table auth (id int not null auto_increment, user_id int, pass_hash char(32) not null, pass_salt char(8) not null, primary key(id), foreign key (user_id) references user(user_id) on delete cascade);

insert into auth(user_id, pass_hash, pass_salt) values (1, 'cb7b71b180c1746a142a0d93b9d063de', '12345678');