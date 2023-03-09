alter table t_project add deploy_type tinyint(1) comment 'frontend deploy type 1:ipfs;2:container ';

alter table t_template add deploy_type tinyint(1) comment 'frontend deploy type 1:ipfs;2:container ';

INSERT INTO t_template (
    template_type_id,
    name,
    description,
    audited,
    last_version,
    whether_display,
    image,
    logo,
    language_type
) VALUES (
        5,
        'Nuxt.js',
        'A Nuxt.app, bootstrapped with create-nuxt-app.',
        1,
        '1.0.0',
        1,
        'https://develop-images.api.hamsternet.io/nuxt.png',
        'https://nuxt.com/assets/design-kit/logo/icon-green.png',
        0
    ),
    (
        5,
        'Next.js',
        'A Next.js app and a Serverless Function API.',
        1,
        '1.0.0',
        1,
        'https://develop-images.api.hamsternet.io/next.png',
        'https://nextjs.org/static/favicon/favicon-32x32.png',
        0
    );
INSERT INTO t_frontend_template_detail (
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
VALUES
(
        (
            SELECT id
            FROM t_template
            WHERE name = 'Nuxt.js'
        ),
        'Nuxt.js',
        '',
        'hamster-template',
        'https://github.com/hamster-template/nuxtjs.git',
        'nuxtjs',
        'main',
        '1.0.0',
        1,
        3,
        '',
        '## Nuxt Example

Deploy your [Nuxt](https://nuxt.com/) project to hamster with zero configuration.

*Live Example: [https://nuxtjs-template.hamster.app](https://nuxtjs-template.hamster.app/)*

Look at the [Nuxt 3 documentation](https://v3.nuxtjs.org/) to learn more.

### Setup

Make sure to install the dependencies:

```
# yarn
yarn

# npm
npm install

# pnpm
pnpm install --shamefully-hoist
```

### Development Server

Start the development server on http://localhost:3000

```
npm run dev
```

### Production

Build the application for production:

```
npm run build
```

Locally preview production build:

```
npm run preview
```

Checkout the [deployment documentation](https://v3.nuxtjs.org/guide/deploy/presets) for more information.'
    ),
(
    (
        SELECT id
        FROM t_template
        WHERE name = 'Next.js'
    ),
    'Next.js',
    'examples',
    'hamster-template',
    'https://github.com/hamster-template/nextjs.git',
    'nextjs',
    'main',
    '1.0.0',
    1,
    4,
    '',
    '### Deploy your own

Deploy your own Next.js project with Hamster.

Demo: https://next-template.hamster.app

### Running Locally

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
```

Open [http://localhost:3000](http://localhost:3000/) with your browser to see the result.

You can start editing the page by modifying `pages/index.js`. The page auto-updates as you edit the file.

[API routes](https://nextjs.org/docs/api-routes/introduction) can be accessed on http://localhost:3000/api/hello. This endpoint can be edited in `pages/api/hello.js`.

The `pages/api` directory is mapped to `/api/*`. Files in this directory are treated as [API routes](https://nextjs.org/docs/api-routes/introduction) instead of React pages.

This project uses [`next/font`](https://nextjs.org/docs/basic-features/font-optimization) to automatically optimize and load Inter, a custom Google Font.

### Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions are welcome!'
);

create table t_container_deploy_param (
                           id int primary key auto_increment,
                           project_id char(36) comment 'project id',
                           workflow_id int comment 'workflow id',
                           container_port     int comment  'container port',
                           service_protocol   varchar(50) comment 'service protocol',
                           service_port       int       comment 'service port',
                           service_target_port    int    comment 'service target port',
                           create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
                           update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间',
                           delete_time timestamp NULL comment '删除时间'
) comment 'container deploy param table';
