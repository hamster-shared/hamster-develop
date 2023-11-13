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
             91,
             5,
             'linea-json-rpc-demo',
             'Call the demo of Linea chain json rpc.',
             1,
             '0.0.1',
             1,
             0,
             'https://g.alpha.hamsternet.io/ipfs/QmPJTikU15M48FFMFKRuVZMsft7GBYUGtNKmgHjhXkmnvp',
             1,
             'https://g.alpha.hamsternet.io/ipfs/Qmbk4jKTmg3chDQacpihDSn8paepHS5gFizqGiwmMvLDTf'
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
           9,
           91,
           'linea-json-rpc-demo',
           '',
           'hamster-template',
           'https://github.com/hamster-template/linea_json_rpc_demo.git',
           'linea_json_rpc_demo',
           'main',
           '0.0.1',
           1,
           5,
           'https://g.alpha.hamsternet.io/ipfs/QmWPutdzGZsN5UHYemvbZ9gUVjBPKa2thTFXAZiHU1nVVx',
           'Call the demo of Linea chain json rpc.'
       );