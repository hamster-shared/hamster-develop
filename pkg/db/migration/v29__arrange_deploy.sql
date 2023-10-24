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
);

create table t_contract_arrange_cache
(
    id                 int auto_increment                     primary key,
    project_id         char(36)                               not null comment 'project id',
    contract_id        char(36)                               not null comment 'contract id',
    contract_name      varchar(100)                           not null comment 'contract name',
    version            varchar(50)                            not null comment 'contract build version',
    original_arrange   text                                   not null comment 'original arrange',
    create_time        timestamp default CURRENT_TIMESTAMP    not null comment 'create time',
    update_time        timestamp default CURRENT_TIMESTAMP    not null comment 'update time'
);
alter table t_contract_arrange
    add arrange_contract_name text  comment 'contract name arranged';

create table t_chain_network
(
    id                 int auto_increment                     primary key,
    logo               varchar(100)                           not null comment 'logo',
    category           varchar(50)                            not null comment 'category',
    chain_id           varchar(50)                            not null comment 'chain id',
    chain_name         varchar(50)                            not null comment 'chain name',
    rpc_url            varchar(100)                           not null comment 'rpc url',
    symbol             varchar(50)                            not null comment 'symbol',
    block_explorer_url varchar(100)                               null comment 'block explorer url',
    decimals           int                                    not null comment 'decimals',
    create_time        timestamp default CURRENT_TIMESTAMP    not null comment 'create time',
    update_time        timestamp default CURRENT_TIMESTAMP    not null comment 'update time'
);
