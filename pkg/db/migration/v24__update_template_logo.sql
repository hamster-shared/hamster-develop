update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmUiEafbupgSwDWWpG5Q8fXrHLBHhMBfJK1APzdcgWnjFT' where id = 2;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmYd8ysx1kWJR7PGzqrSzPmhCtcjEiwLJMKCAprrWk4jPw' where id = 3;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmWhDyrKC3855hAzGZz1HDyez6UM6NRW8SQsyJRm3no46F' where id = 4;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmWG1UsNyeQbzh2Edbm6exdaJfqj9EXgCRiApP1Xkyj24p' where id = 5;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmZTXorRr8HmKbm1xEAWLGXTfpjfwHRfvMWNtBXg4Xg2sd' where id = 8;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmVPZy4GesehCxeLHijbnt1wSq1DJNWEAy244Dx3ZTHS8H' where id = 28;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmeVFcaVQYdpWZmx4UDdPKAiCRsh784vKnf57biuH3JFU5' where id = 29;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmaH6MQ4QmszZs37ZncRV8BaYacn3HPeWugiCkJ8XJrrpU' where id = 30;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmddQVuHLwvLS7qDnvLhUkKJZJC7d7ApcJMm5xsVWgjRwa' where id = 31;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmPTqxymxKEGw5X9oNVGc38qqY24X4fGoDCdrKdfLU8LU1' where id = 32;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmPeStbFnayXcCDpBLfzdZVucTKDfrFhSybMcafvJu4t1u' where id = 33;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmUEn2zsMBW8xojsmaWrCRbi1K4x2Koz2K3hypACqDWT99' where id = 34;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmUitnWFT9M7Kjmh9z4su4LGFfQsyqdCbGQ2vVYceoCDrj' where id = 35;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmYCLdgTEmLwVuecuzzXxC9aUhsPAW2cLd6TDi9QBEsQCn' where id = 36;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmNpaEzye5ipuD3vAqfCw7y95XgkBRMVaXc7SeyBvaCQhz' where id = 37;
update t_template set logo='https://g.alpha.hamsternet.io/ipfs/QmWG1UsNyeQbzh2Edbm6exdaJfqj9EXgCRiApP1Xkyj24p' where id = 42;

update t_frontend_template_detail set show_url='https://g.alpha.hamsternet.io/ipfs/QmZMP5jmafwZDeE98Hu9kdT9wDuPXDs7ygpwFVzQRD7iqP' where id = 1;
update t_frontend_template_detail set show_url='https://g.alpha.hamsternet.io/ipfs/QmRZMRqvCE1qLBzvC4HHvGSAyKcTEqYgj3Mq3fFwjBd1XB' where id = 2;


insert into t_template_type (
                id,
                name,
                description,
                type
)values (
         7,
         'polkadot',
         '',
         3
        );
alter table t_template
    modify description text;
insert into t_template (
    id,
    template_type_id,
    name,
    description,
    audited,
    last_version,
    whether_display,
    language_type,
    logo,
    image,
    deploy_type
) values (
    48,
    7,
    'substrate-node-template',
    'A standalone version of this template is available for each release of Polkadot in the Substrate Developer Hub Parachain Template repository. The parachain template is generated directly at each Polkadot release branch from the Node Template in Substrate upstream.',
     1,
    '0.0.1',
    1,
    1,
    'https://g.alpha.hamsternet.io/ipfs/QmNpaEzye5ipuD3vAqfCw7y95XgkBRMVaXc7SeyBvaCQhz',
    'https://g.alpha.hamsternet.io/ipfs/QmPbUjgPNW1eBVxh1zVgF9F7porBWijYrAeMth9QDPwEXk',
    2
         );

INSERT INTO t_template_detail (
    id,
    template_id,
    name,
    audited,
    extensions,
    description,
    examples,
    resources,
    abi_info,
    byte_code,
    author,
    repository_url,
    repository_name,
    branch,
    version,
    code_sources,
    title,
    title_description
) values (
            48,
            48,
            'substrate-node-template',
            1,
            'A standalone version of this template is available for each release of Polkadot in the Substrate Developer Hub Parachain Template repository. The parachain template is generated directly at each Polkadot release branch from the Node Template in Substrate upstream.',
            '',
            '',
            '',
            '',
            '',
            'hamster-template',
            'https://github.com/hamster-template/substrate-node-template.git',
            'substrate-node-template',
            'main',
            '0.0.1',
            '',
            '',
            ''
         );