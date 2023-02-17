alter table t_user add first_state tinyint(1) default 0 comment 'Whether to log in for the first time';
alter table t_user add user_email varchar(50) comment 'user email';