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
             92,
             5,
             'react-nextjs',
             'This is a Next.js + React project bootstrapped with create-next-app.',
             1,
             '18.2.0',
             0,
             0,
             'https://nextjs.org/static/favicon/favicon-32x32.png',
             1,
             'https://g.alpha.hamsternet.io/ipfs/QmfCkFjMvdnZJin9SwJARk6Ro2kkyLzRfLX5rKuM8coGtc'
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
           10,
           92,
           'react-nextjs',
           '',
           'hamster-template',
           'https://github.com/hamster-template/react-nextjs.git',
           'react-nextjs',
           'main',
           '18.2.0',
           1,
           5,
           'https://g.alpha.hamsternet.io/ipfs/QmPixY9GqBHbbCLcE4ttZZQw76taXbifRN6a2Kv81dk6bK',
           'This is a Next.js + React project bootstrapped with create-next-app.'
       );


update t_template
set last_version = '18.2.0'
where id = 19;

update t_frontend_template_detail
set version = '18.2.0'
where id = 2;