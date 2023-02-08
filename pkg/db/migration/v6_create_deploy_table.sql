create table t_frontend_deploy (
                            id int primary key auto_increment,
                            project_id char(36)  comment 'project id',
                            workflow_id int comment 'workflow id',
                            workflow_detail_id int comment 'workflow detail id',
                            package_id        int  comment 'frontend package id',
                            name varchar(100) not null comment 'package name',
                            version varchar(50) comment 'package version',
                            branch  varchar(100) comment 'build code info',
                            domain  varchar(100) null comment 'frontend deploy domains',
                            deploy_info varchar(100) comment 'deploy info example CID',
                            create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment 'create time',
                            deploy_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment 'deploy time',
                            delete_time timestamp NULL comment 'delete_time'
) comment 'frontend deploy table';