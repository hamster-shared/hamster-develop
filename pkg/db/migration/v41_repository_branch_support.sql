alter table t_backend_package
    add commit_id varchar(50) null comment '代码提交id';

alter table t_backend_package
    add commit_info varchar(100) null comment '代码提交信息';

alter table t_frontend_package
    add commit_id varchar(50) null comment '代码提交id';

alter table t_frontend_package
    add commit_info varchar(100) null comment '代码提交信息';

alter table t_workflow_detail
    add branch varchar(50) null comment '分支信息';

alter table t_workflow_detail
    add commit_id varchar(50) null comment '代码提交id';

alter table t_workflow_detail
    add commit_info varchar(100) null comment '代码提交信息';

alter table t_contract
    add commit_id varchar(50) null comment '代码提交id';
alter table t_contract
    add commit_info varchar(100) null comment '代码提交信息';

alter table t_frontend_deploy
    add commit_id varchar(50) null comment '代码提交id';

alter table t_frontend_deploy
    add commit_info varchar(100) null comment '代码提交信息';


alter table t_contract_deploy
    add branch varchar(50) null comment '分支信息';

alter table t_contract_deploy
    add commit_id varchar(50) null comment '代码提交id';

alter table t_contract_deploy
    add commit_info varchar(100) null comment '代码提交信息';

alter table t_backend_deploy
    add branch varchar(50) null comment '分支信息';

alter table t_backend_deploy
    add commit_id varchar(50) null comment '代码提交id';

alter table t_backend_deploy
    add commit_info varchar(100) null comment '代码提交信息';
