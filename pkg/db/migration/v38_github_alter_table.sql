alter table t_handler_failed_data
    add column app_id BIGINT  comment 'app id';

alter table t_git_app_install
    add column app_id BIGINT  comment 'app id';