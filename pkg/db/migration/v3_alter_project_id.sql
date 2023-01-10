alter table t_project
    modify id char(36)  comment '主键ID';
alter table t_workflow
    modify project_id char(36) not null comment '项目id';

alter table t_workflow_detail
    modify project_id char(36) null;

alter table t_contract
    modify project_id char(36) null comment '项目id';

alter table t_report
    modify project_id char(36) null comment '项目ID';