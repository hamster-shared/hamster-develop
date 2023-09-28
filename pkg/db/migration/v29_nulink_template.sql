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
    deploy_type,
    image
) values (
             58,
             5,
             'nulink-web-agent-demo',
             'NuLink agent integration demo is a demo for integration nulink-web-agent project.',
             1,
             '0.0.1',
             1,
             0,
             'https://g.alpha.hamsternet.io/ipfs/QmSEZGr1tT7hU2jzFCx2HJQobqrpPmfsxKJr8XBXSGWq9D',
             1,
             'https://g.alpha.hamsternet.io/ipfs/QmUy8gcpgm3bbQdzGj7YhcgQrHPxgTFVmpGjPiDMf1R3bm'
         );
INSERT INTO t_frontend_template_detail (
    id,
    template_id,
    name,
    examples,
    author,
    repository_url,
    repository_name,
    branch,
    version,
    audited,
    template_type,
    show_url,
    description
)
VALUES (
        7,
        58,
        'nulink-web-agent-demo',
        '',
        'hamster-template',
        'https://github.com/hamster-template/nulink-web-agent-demo.git',
        'nulink-web-agent-demo',
        'main',
        '0.0.1',
        1,
        5,
        'https://g.alpha.hamsternet.io/ipfs/QmUy8gcpgm3bbQdzGj7YhcgQrHPxgTFVmpGjPiDMf1R3bm',
        'NuLink agent integration demo is a demo for integration nulink-web-agent project.'
       );


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
    deploy_type,
    image
) values (
             73,
             5,
             'nulink-agent-sdk-tutorial',
             'Demo interact with NuLink Agent SDK.',
             1,
             '0.0.1',
             1,
             0,
             'https://g.alpha.hamsternet.io/ipfs/QmSEZGr1tT7hU2jzFCx2HJQobqrpPmfsxKJr8XBXSGWq9D',
             1,
             'https://g.alpha.hamsternet.io/ipfs/QmZ1LB1gDmtwRMuxrJCq24c8cNjckY3tRzKdvAhw2v5uP2'
         );
INSERT INTO t_frontend_template_detail (
    id,
    template_id,
    name,
    examples,
    author,
    repository_url,
    repository_name,
    branch,
    version,
    audited,
    template_type,
    show_url,
    description
)
VALUES (
           8,
           73,
           'nulink-agent-sdk-tutorial',
           '',
           'hamster-template',
           'https://github.com/hamster-template/nulink-agent-sdk-tutorial.git',
           'nulink-agent-sdk-tutorial',
           'main',
           '0.0.1',
           1,
           5,
           'https://g.alpha.hamsternet.io/ipfs/QmbCbkLrjX7KgRScBFrbtRWq37ZarJXWDVAJYtntfZcmkq',
           'Demo interact with NuLink Agent SDK.'
       );