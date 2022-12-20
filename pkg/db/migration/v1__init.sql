create table t_project (
                           id int primary key auto_increment,
                           name varchar(100) not null  comment '项目名称',
                           user_id varchar(50) not null comment '拥有者',
                           type tinyint(1) not null default 1 comment '项目类型(1: Contract,2: Frontend, 3: Blockchain Node)',
                           repository_url varchar(200) comment '仓库地址',
                           frame_type tinyint(1) comment '项目框架（1. truffle,hardhat;2: !ink; 3: move ; 4: vue.js; 5: nuxt.js;6: next.js;7: vite;8:Angular)',
                           creator int(11) ,
                           delete_user int(11),
                           create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
                           update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间',
                           delete_time timestamp NULL comment '删除时间'
) comment '项目表';

create table t_user (
                        id int primary key auto_increment,
                        username varchar(50) not null comment '用户名',
                        token varchar(100) comment 'token',
                        create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间'
) comment '用户表';

create table t_workflow (
                            id int primary key auto_increment,
                            project_id int not null comment '项目id',
                            type tinyint(1) default 1 comment 'pipeline类型（1:checked,2: build)',
                            exec_file text ,
                            last_exec_id int,
                            create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
                            update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间',
                            delete_time timestamp NULL comment '删除时间'
) comment '工作流表';

create table t_workflow_detail (
                                   id int primary key auto_increment,
                                   workflow_id int not null comment '工作流id',
                                   exec_number int comment '执行ID',
                                   stage_info text comment '状态信息',
                                   trigger_user varchar(50) comment '触发用户',
                                   trigger_mode tinyint(1) default '1' comment '触发方式（1： Manual)',
                                   code_info varchar(50) comment '代码信息 (master| commit_id )',
                                   status tinyint(1) default 0 comment 'pipeline 状态（0:未执行,1:执行中,2:success,3:failed,4: cancel） ',
                                   start_time timestamp null comment '开始时间',
                                   end_time timestamp null comment '结束时间',
                                   create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
                                   update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间',
                                   delete_time timestamp NULL comment '删除时间'
) comment '工作流详情';

create table t_contract (
                            id int primary key auto_increment,
                            project_id int comment '项目id',
                            workflow_id int comment '工作流id',
                            workflow_detail_id int comment '工作流详情id',
                            name varchar(100) not null comment '合约名',
                            version varchar(50) comment '合约版本',
                            network varchar(50) comment '部署网络',
                            build_time timestamp null comment '构建时间',
                            abi_info text comment 'abi信息',
                            byte_code text comment '合约字节码',
                            create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间'
) comment '合约表';

create table t_contract_deploy (
                                   id int primary key auto_increment,
                                   contract_id int comment '合约id',
                                   project_id int comment '项目ID',
                                   version varchar(50) comment '部署版本',
                                   deploy_time timestamp null comment '部署时间',
                                   network varchar(50) comment '部署网络',
                                   address varchar(100) comment '部署地址',
                                   create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间'
) comment '合约部署表';

create table t_report (
                          id int primary key auto_increment,
                          project_id int comment '项目ID',
                          workflow_id int comment '工作流id',
                          workflow_detail_id int comment '工作流详情id',
                          name varchar(100) comment '报告名称',
                          type tinyint(1) default 1 comment '报告类型(1: check)',
                          check_tool varchar(50) comment '检查工具',
                          result varchar(100) comment '检查结果',
                          check_time timestamp null comment '检查时间',
                          report_file text comment '检查报告内容',
                          create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间'
) comment '构建物报告表';


create table t_template_type (
                                 id int primary key auto_increment,
                                 name varchar(100) not null comment '模板类别名称',
                                 description varchar(2000) comment '模板类别描述',
                                 type tinyint(1) comment '模板名称(1:solidity; 2: !ink; 3: move; 4: frontend) ',
                                 create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
                                 update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间',
                                 delete_time timestamp NULL comment '删除时间'
) comment '模板分类表';

create table t_template (
                            id int primary key auto_increment,
                            template_type_id int not null comment '模板类型id',
                            name varchar(100) not null comment '模板名',
                            description varchar(255) comment '模板类别描述',
                            audited tinyint(1) comment '是否审计 0: false, 1: true ',
                            last_version varchar(20) comment '模板版本',
                            logo    varchar(50)  comment '模板图标',
                            create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
                            update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间',
                            delete_time timestamp NULL comment '删除时间'
) comment '模板表';

create table t_template_detail (
                                   id int primary key auto_increment,
                                   template_id int not null comment '模板id',
                                   name varchar(100) not null comment '模板名称',
                                   audited tinyint(1) default 0 comment '是否审计(0: false,1: true)',
                                   extensions  varchar(200),
                                   description varchar(255) comment '模板描述',
                                   examples    varchar(255),
                                   resources   varchar(255),
                                   abi_info    text,
                                   byte_code   text,
                                   author      varchar(50) comment '模板仓库作者',
                                   repository_url  varchar(200)    comment '模板仓库地址',
                                   branch      varchar(50),
                                   version     varchar(50),
                                   code_sources    text,
                                   create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
                                   update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间',
                                   delete_time timestamp NULL comment '删除时间'

) comment '模板详情';
