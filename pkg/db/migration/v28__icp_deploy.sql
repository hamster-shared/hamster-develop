alter table t_contract_deploy
    add name varchar(50) null comment 'ICP部署， 罐名';


alter table t_contract
    add branch varchar(50);

alter table t_icp_canister
    add contract varchar(50);


create table t_backend_package
(
    id                 int auto_increment
        primary key,
    project_id         char(36)                           not null comment 'project id',
    workflow_id        int                               not  null comment 'workflow id',
    workflow_detail_id int                               not  null comment 'workflow detail id',
    name               varchar(100)                        not null comment 'package name',
    version            varchar(50)                        not null comment 'package version',
    build_time         timestamp                           null comment 'build time',
    abi_info            mediumtext                          null comment 'abi信息',
    create_time        timestamp default CURRENT_TIMESTAMP null comment 'create time',
    type                tinyint(1)                          null comment 'project frame type： see #consts.ProjectFrameType',
    status              tinyint(1)                      null comment 'status #consts.DeployStatus',
    branch             varchar(100)                        null comment 'build code info'
)
    comment 'backend deploy  package table';


create table t_contract_deploy
(
    id              int auto_increment
        primary key,
    contract_id     int                                 null comment '合约id',
    project_id      char(36)                            null comment '项目ID',
    version         varchar(50)                         null comment '部署版本',
    deploy_time     timestamp                           null comment '部署时间',
    network         varchar(50)                         null comment '部署网络',
    address         varchar(100)                        null comment '部署地址',
    create_time     timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    type            int       default 1                 not null comment '合约类型： 1:evm, 4:starknet',
    declare_tx_hash char(72)                            null comment 'declare 交易hash(starknet专用)',
    deploy_tx_hash  char(72)                            null comment '部署hash',
    status          int       default 2                 null comment ' 1: deploying, 2: success , 3: fail',
    abi_info        mediumtext                          null comment 'abi信息',
    name            varchar(50)                         null comment 'ICP部署， 罐名'
)
    comment '合约部署表';



create table t_backend_deploy
(
    id                 int auto_increment
        primary key,
    package_id         int                              not null comment 'backend package id',
    project_id         char(36)                            null comment 'project id',
    workflow_id        int                                 null comment 'workflow id',
    workflow_detail_id int                                 null comment 'workflow detail id',
    version            varchar(50)                        not null comment 'package version',
    deploy_time        timestamp default CURRENT_TIMESTAMP null comment 'deploy time',
    network               varchar(50)                        not null comment 'network',
    address               varchar(50)                        not null comment 'address',
    create_time        timestamp default CURRENT_TIMESTAMP null comment 'create time',
    type                tinyint(1)                          null comment 'project frame type： see #consts.ProjectFrameType',
    deploy_tx_hash  char(72)                            null comment '部署hash',
    status          int       default 2                 null comment ' 1: deploying, 2: success , 3: fail',
    abi_info            mediumtext                          null comment 'abi信息',
    name            varchar(50)                         null comment 'ICP部署， 罐名'
)
    comment 'backend deploy table';


alter table t_contract_deploy drop column name;

