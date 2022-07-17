drop database if exists catalyst_experience;
create database catalyst_experience;
use catalyst_experience;

create table user (
  id bigint not null auto_increment,
  invitation_token varchar(12) not null,
  username varchar(255) not null,
  password varchar(255) not null,
  role enum('admin', 'invitation') not null,
  created_at bigint not null,

  primary key (id),
  unique (username)
);

create table invitation (
  id bigint not null auto_increment,
  token varchar(12) not null,
  status enum('disabled', 'inactive', 'active') not null,
  created_at bigint not null,

  primary key (id),
  unique (token)
);
