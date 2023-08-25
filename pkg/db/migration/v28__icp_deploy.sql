alter table t_contract_deploy
    add name varchar(50) null comment 'ICP部署， 罐名';

alter table t_contract
    add branch varchar(50);

alter table t_icp_canister
    add contract varchar(50);

