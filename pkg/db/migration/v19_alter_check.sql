alter table t_workflow
    add column  tool_type tinyint(1) comment 'tool type',
    add column  tool varchar(50) comment 'tool';



alter table t_workflow_detail
    add column  tool_type tinyint(1) comment 'tool type',
    add column  tool varchar(50) comment 'tool';

alter table t_report
    add column  tool_type tinyint(1) comment 'tool type';

create table t_report_detail (
                        id BIGINT primary key auto_increment,
                        report_id  int comment 'report id',
                        project_id int comment 'project id',
                        workflow_id int comment 'workflow id',
                        workflow_detail_id int comment 'workflow detail id',
                        total_issues  int comment 'total issue',
                        report_result mediumtext comment 'report result',
                        report_issues mediumtext comment 'report issues',
                        create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间'
) comment 'report detail';