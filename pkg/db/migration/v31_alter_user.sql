alter table t_user
    add login_type tinyint  comment 'login type';

alter table t_user_wallet
    add first_state tinyint  comment 'first login';