alter table t_contract
add column aptos_mv text
after byte_code;

-- 为 t_project 表添加 params 字段
alter table t_project
add column params text
after deploy_type;