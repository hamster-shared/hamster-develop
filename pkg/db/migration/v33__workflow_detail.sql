alter table t_contract_arrange_execute
    collate = utf8mb4_unicode_ci;

alter table t_contract_arrange_execute
    modify project_id char(36) collate utf8mb4_unicode_ci not null comment 'project id';



create view v_workflow_detail as (
    select t.id, t.project_id, 'workflow' as engine, t.create_time,t.type
    from t_workflow_detail t
    union all
    select t.id,t.project_id, 'arrange_execute' as engine, t.create_time,3 as type
    from t_contract_arrange_execute t
);
