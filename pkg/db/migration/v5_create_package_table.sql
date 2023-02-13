create table t_frontend_package (
                            id int primary key auto_increment,
                            project_id char(36)  comment 'project id',
                            workflow_id int comment 'workflow id',
                            workflow_detail_id int comment 'workflow detail id',
                            name varchar(100) not null comment 'package name',
                            version varchar(50) comment 'package version',
                            branch  varchar(100) comment 'build code info',
                            domain  varchar(100) null comment 'frontend deploy domains',
                            package_identity  varchar(100)  comment 'frontend package identity',
                            build_time timestamp null comment 'build time',
                            create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment 'create time'
) comment 'frontend package table';