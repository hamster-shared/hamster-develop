ALTER TABLE t_template ADD COLUMN image TEXT COMMENT '模板图片';
DROP TABLE IF EXISTS t_frontend_template_type;
DROP TABLE IF EXISTS t_frontend_template;
DROP TABLE IF EXISTS t_frontend_template_detail;
create table t_frontend_template_detail (
    id int primary key auto_increment,
    create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间',
    delete_time timestamp NULL comment '删除时间',
    template_id int not null comment '模板 id',
    name varchar(100) not null comment '模板名称',
    examples varchar(255),
    author varchar(50) comment '模板仓库作者',
    repository_url varchar(200) comment '模板仓库地址',
    repository_name varchar(200) comment '模板仓库名称',
    show_url      varchar(200) comment 'show url',
    template_type tinyint(1) comment 'template type,1:vue;2:react',
    branch varchar(50),
    version varchar(50),
    audited boolean not null default 0 comment '是否审核通过',
    description text comment '模板 markdown 描述'
) comment '模板详情';
-- 向模板类型表插入一条数据，表示前端模板
DELETE FROM t_template_type
WHERE name = 'frontend';
insert into t_template_type (
        id,
        name,
        description,
        type
    )
values (
        5,
        'frontend',
        '',
        2
    );
-- 模板表插入数据
DELETE FROM t_template
WHERE template_type_id = 5;
INSERT INTO t_template (
        template_type_id,
        name,
        description,
        audited,
        last_version,
        whether_display,
        image,
        logo,
        language_type,
        deploy_type
    )
VALUES
#     (
#         5,
#         'Nuxt.js',
#         'A Nuxt.app, bootstrapped with create-nuxt-app.',
#         1,
#         '1.0.0',
#         1,
#         'https://develop-images.api.hamsternet.io/nuxt.png',
#         'https://nuxt.com/assets/design-kit/logo/icon-green.png'
#     ),
    (
        5,
        'Vue.js',
        'A Vue.js app, created with the Vue CLI.',
        1,
        '1.0.0',
        1,
        'https://develop-images.api.hamsternet.io/vue.png',
        'https://vuejs.org/logo.svg',
        0,
        1
    ),
    (
        5,
        'React.js',
        'A client-side React app created with create-react-app.',
        1,
        '1.0.0',
        1,
        'https://develop-images.api.hamsternet.io/react.png',
        'https://reactjs.org/favicon.ico',
        0,
        1
    );
-- 向前端模板详情表插入数据
DELETE FROM t_frontend_template_detail;
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
# (
#         (
#             SELECT id
#             FROM t_template
#             WHERE name = 'Nuxt.js'
#         ),
#         'Nuxt.js',
#         'examples',
#         'hamster-template',
#         'https://github.com/hamster-template/nuxtjs.git',
#         'nuxtjs',
#         'main',
#         '1.0.0',
#         1,
#         '## Nuxt Example
#
# Deploy your [Nuxt](https://nuxt.com/) project to hamster with zero configuration.
#
# *Live Example: [https://nuxtjs-template.hamster.app](https://nuxtjs-template.hamster.app/)*
#
# Look at the [Nuxt 3 documentation](https://v3.nuxtjs.org/) to learn more.
#
# ### Setup
#
# Make sure to install the dependencies:
#
# ```
# # yarn
# yarn
#
# # npm
# npm install
#
# # pnpm
# pnpm install --shamefully-hoist
# ```
#
# ### Development Server
#
# Start the development server on http://localhost:3000
#
# ```
# npm run dev
# ```
#
# ### Production
#
# Build the application for production:
#
# ```
# npm run build
# ```
#
# Locally preview production build:
#
# ```
# npm run preview
# ```
#
# Checkout the [deployment documentation](https://v3.nuxtjs.org/guide/deploy/presets) for more information.'
#     ),
    (
        (
            SELECT id
            FROM t_template
            WHERE name = 'Vue.js'
        ),
        'Vue.js',
        'examples',
        'hamster-template',
        'https://github.com/hamster-template/vuejs.git',
        'vuejs',
        'main',
        '1.0.0',
        1,
        1,
        'http://g.develop.hamsternet.io/ipfs/QmTaSjCsdopHeiDrGfdXJk8xxxf8tdkNK7SWazae2LtnWu',
        '### Deploy Your Own

Deploy your own Vue.js project with hamster.



*Live Example: [https://vue-template.hamster.app](https://vue-template.vercel.app/)*

### Running Locally

```
yarn install
```

Compile and hot-reload for development

```
yarn serve
```

Compile and minify for production

```
yarn build
```

Lint and fix files

```
yarn lint
```'
    ),
    (
        (
            SELECT id
            FROM t_template
            WHERE name = 'React.js'
        ),
        'React.js',
        'examples',
        'hamster-template',
        'https://github.com/hamster-template/reactjs.git',
        'reactjs',
        'main',
        '1.0.0',
        1,
        2,
        'http://g.develop.hamsternet.io/ipfs/QmRZMRqvCE1qLBzvC4HHvGSAyKcTEqYgj3Mq3fFwjBd1XB',
        '### Deploy Your Own

Deploy your own Create React App project with Vercel.



*Live Example: https://create-react-template.hamster.app/*

### Available Scripts

In the project directory, you can run:

#### `npm start`

Runs the app in the development mode. Open [http://localhost:3000](http://localhost:3000/) to view it in your browser.

The page will reload when you make changes. You may also see any lint errors in the console.

#### `npm test`

Launches the test runner in the interactive watch mode. See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

#### `npm run build`

Builds the app for production to the `build` folder.

It correctly bundles React in production mode and optimizes the build for the best performance. The build is minified and the filenames include the hashes.'
    );