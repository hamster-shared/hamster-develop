alter table t_workflow
    add column  tool_type tinyint(1) comment 'tool type',
    add column  tool varchar(200) comment 'tool';

alter table t_report
    add column  tool_type tinyint(1) comment 'tool type',
    add column  issues    int        comment 'total issues',
    add column  meta_scan_overview mediumtext comment 'meta scan overview';

alter table t_report
    modify report_file mediumtext;

alter table t_template_detail
    add column gist_id varchar(50),
    add column default_file varchar(50);

alter table t_project
    add column gist_id varchar(50),
    add column default_file varchar(50);
update t_template_detail set gist_id='f4691149c7d856a8f3a2cd7da89d234e',default_file='ERC4907.sol' where id = 2;
update t_template_detail set gist_id='57891463242d5ed418fd68a8549452bd',default_file='NFT.sol' where id = 3;
update t_template_detail set gist_id='a85f5d6c113baeebfea41a81977e749b',default_file='Edition.sol' where id = 4;
update t_template_detail set gist_id='abffdb1cc106fb8111dc976a4af20bca',default_file='ERC721CommunityStream.sol' where id = 5;
update t_template_detail set gist_id='129f5632ae09d5dacad8573c92ebd8af',default_file='FunctionsConsumer.sol' where id = 42;
update t_template_detail set gist_id='1a32466fffd5dc50510fd058fd22f0b8',default_file='MyERC20.sol' where id = 8;
update t_template_detail set gist_id='9c091463f40db5c1114d0b59062caf99',default_file='IUniswapV2ERC20.sol' where id = 13;
update t_template_detail set gist_id='0f3ba749f2667798e0dc4d8403f333c2',default_file='token-staking.move' where id = 17;
update t_template_detail set gist_id='532efb49674fd409d37078d73f68a816',default_file='token-vesting.move' where id = 18;
update t_template_detail set gist_id='ec84f26a32b36b6317b724b2f4d6555e',default_file='nftborrowlend.move' where id = 19;
update t_template_detail set gist_id='d9905a6aedcce2cb804c5610f9ca760d',default_file='raffle.move' where id = 20;
update t_template_detail set gist_id='d4ada2f413e324269fa5d34a4ae58c4b',default_file='ERC721.cairo' where id = 24;
update t_template_detail set gist_id='900c42ac75ff2ed270044c472d42ba75',default_file='ERC1155.cairo' where id = 25;
update t_template_detail set gist_id='5b6bb95073550ca97289af42051f5e58',default_file='ERC20.cairo' where id = 26;
update t_template_detail set gist_id='016887603d86b5de0131351f23739733',default_file='todolist.move' where id = 27;
update t_template_detail set gist_id='8d18cb0269e53da291105ffb56894a62',default_file='auction.move' where id = 41;

# create table t_report_detail (
#                         id BIGINT primary key auto_increment,
#                         report_id  int comment 'report id',
#                         project_id int comment 'project id',
#                         workflow_id int comment 'workflow id',
#                         workflow_detail_id int comment 'workflow detail id',
#                         total_issues  int comment 'total issue',
#                         report_result mediumtext comment 'report result',
#                         report_issues mediumtext comment 'report issues',
#                         create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间'
# ) comment 'report detail';