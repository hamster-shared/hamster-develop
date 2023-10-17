alter table t_contract
    add code_info varchar(255)  comment '代码信息 commit_id | commit_date ｜ message';
alter table t_backend_package
    add code_info varchar(255)  comment '代码信息 commit_id | commit_date ｜ message';

alter table t_workflow_detail
    change code_info  code_branch varchar(50)  comment '代码分支';
alter table t_workflow_detail
    add code_info varchar(255)  comment '代码信息 commit_id | commit_date ｜ message';