drop table if exists t_request_log;
create table t_request_log (
                           id int primary key auto_increment,
                           url varchar(255) not null  comment 'url路由',
                           method varchar(10) not null comment 'method',
                           token varchar(255) not null comment 'token',
                           create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
                           update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间',
                           delete_time timestamp NULL comment '删除时间'
) comment '请求日志表';
