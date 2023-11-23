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
    private     bool    comment 'private or public',
    language    varchar(50),
    update_at   timestamp NULL DEFAULT CURRENT_TIMESTAMP comment 'update time',
    create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment 'create time'
) comment 'github repo';


create table t_git_app_install (
    id BIGINT primary key auto_increment,
    install_user_id bigint comment 'install user id',
    name    varchar(100) comment 'install name',
    repository_selection  varchar(20) comment 'selected,all',
    install_id bigint    comment 'install id',
    user_id  bigint    comment 'user id',
    avatar_url varchar(100) comment 'avatar url',
    create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment 'create time'
) comment 'github app install user';