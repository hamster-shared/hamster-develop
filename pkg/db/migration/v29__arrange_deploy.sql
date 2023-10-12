create table t_contract_arrange
(
    id                 int auto_increment                     primary key,
    project_id         char(36)                               not null comment 'project id',
    version            varchar(50)                            not  null comment 'contract build version',
    original_arrange   text                                   not  null comment 'original arrange',
    create_time        timestamp default CURRENT_TIMESTAMP    not null comment 'create time',
    update_time        timestamp default CURRENT_TIMESTAMP    not null comment 'update time'
);

create table t_contract_arrange_execute
(
    id                      int auto_increment                     primary key,
    project_id              char(36)                               not null comment 'project id',
    version                 varchar(50)                            not null comment 'contract build version',
    fk_arrange_id           int                                    not null comment 'contract build version',
    network                 varchar(50)                            not null comment 'network',
    arrange_process_data    text                                       null comment 'Process data for choreographed execution',
    create_time             timestamp default CURRENT_TIMESTAMP    not null comment 'create time',
    update_time             timestamp default CURRENT_TIMESTAMP    not null comment 'update time'
)