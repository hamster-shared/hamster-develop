alter table t_contract
    add type int(1) default '1' not null comment '合约类型： 1:evm, 4:starknet';

alter table t_contract
    add status int(1) default 2 null comment ' 1: deploying, 2: success , 3: fail';

alter table t_contract_deploy
    add type int(1) default '1' not null comment '合约类型： 1:evm, 4:starknet';

alter table t_contract_deploy
    add declare_tx_hash char(72) null comment 'declare 交易hash(starknet专用)';

alter table t_contract_deploy
    add deploy_tx_hash char(72) null comment '部署hash';

alter table t_contract_deploy
    add status int(1) default 2 null comment ' 1: deploying, 2: success , 3: fail';

alter table t_contract
    modify abi_info mediumtext null comment 'abi信息';

