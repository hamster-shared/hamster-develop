create table aline.t_user_wallet
(
    id          int auto_increment
        primary key,
    address     varchar(64)                         null comment 'address',
    create_time timestamp default CURRENT_TIMESTAMP not null comment 'create time',
    user_id     int                                 null comment 'user_id'
)
    comment '用户钱包地址';
