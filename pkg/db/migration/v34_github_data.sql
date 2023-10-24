create table t_handler_failed_data (
                        id BIGINT primary key auto_increment,
                        installation_id int not null comment 'install id',
                        action      varchar(50) comment 'installation action',
                        create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment 'create time'
) comment 'github app installation handler failed';

create table t_git_repo (
    id BIGINT primary key auto_increment,
    user_id int     comment 'user id',
    repo_id         int      comment 'repo id',
    installation_id int not null comment 'install id',
    name    varchar(100) comment 'repo name',
    clone_url   varchar(200) comment 'repo clone url',
    ssh_url     varchar(200) comment 'repo ssh url',
    default_branch  varchar(50) comment 'default branch',
    create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment 'create time'
) comment 'github repo';