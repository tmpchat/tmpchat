create table if not exists tmpchat.room (
  id bigint not null auto_increment primary key,
  uri varchar(128) not null,
  title varchar(512) not null,
  created_at datetime default current_timestamp not null,
  update_at datetime default current_timestamp on update current_timestamp not null,
  deleted_at datetime null
)
engine=InnoDB;
