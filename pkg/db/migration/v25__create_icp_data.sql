create table t_icp_dfx_data (
    id int primary key auto_increment,
    project_id char(36)  comment 'project id',
    dfx_data   text,
    create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间'
);

create table t_chain_template_detail(
    id int primary key auto_increment,
    template_id int not null comment '模板 id',
    name varchar(100) not null comment '模板名称',
    author varchar(50) comment '模板仓库作者',
    repository_url varchar(200) comment '模板仓库地址',
    repository_name varchar(200) comment '模板仓库名称',
    show_url      varchar(200) comment 'show url',
    branch varchar(50),
    version varchar(50),
    audited boolean not null default 0 comment '是否审核通过',
    description text comment '模板 markdown 描述',
    create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间',
    delete_time timestamp NULL comment '删除时间'
);

insert into t_chain_template_detail (
    id,
    template_id,
    name,
    author,
    repository_url,
    repository_name,
    show_url,
    branch,
    version,
    audited,
    description
) values (
    1,
    48,
    'substrate-node-template',
    'hamster-template',
    'https://github.com/hamster-template/substrate-node-template.git',
    'substrate-node-template',
    'https://g.alpha.hamsternet.io/ipfs/QmPbUjgPNW1eBVxh1zVgF9F7porBWijYrAeMth9QDPwEXk',
    'main',
    '1.0.0',
    1,
    '# Substrate Node Template

A fresh [Substrate](https://substrate.io/) node, ready for hacking :rocket:

A standalone version of this template is available for each release of Polkadot in the [Substrate Developer Hub Parachain Template](https://github.com/substrate-developer-hub/substrate-parachain-template/) repository.
The parachain template is generated directly at each Polkadot release branch from the [Node Template in Substrate](https://github.com/paritytech/substrate/tree/master/bin/node-template) upstream

It is usually best to use the standalone version to start a new project.
All bugs, suggestions, and feature requests should be made upstream in the [Substrate](https://github.com/paritytech/substrate/tree/master/bin/node-template) repository.

## Getting Started

Depending on your operating system and Rust version, there might be additional packages required to compile this template.
Check the [installation](https://docs.substrate.io/install/) instructions for your platform for the most common dependencies.
Alternatively, you can use one of the [alternative installation](#alternative-installations) options.

### Build

Use the following command to build the node without launching it:

```sh
cargo build --release
```

### Embedded Docs

After you build the project, you can use the following command to explore its parameters and subcommands:

```sh
./target/release/node-template -h
```

You can generate and view the [Rust Docs](https://doc.rust-lang.org/cargo/commands/cargo-doc.html) for this template with this command:

```sh
cargo +nightly doc --open
```

### Single-Node Development Chain

The following command starts a single-node development chain that doesn''t persist state:

```sh
./target/release/node-template --dev
```

To purge the development chain''s state, run the following command:

```sh
./target/release/node-template purge-chain --dev
```

To start the development chain with detailed logging, run the following command:

```sh
RUST_BACKTRACE=1 ./target/release/node-template -ldebug --dev
```

Development chains:

- Maintain state in a `tmp` folder while the node is running.
- Use the **Alice** and **Bob** accounts as default validator authorities.
- Use the **Alice** account as the default `sudo` account.
- Are preconfigured with a genesis state (`/node/src/chain_spec.rs`) that includes several prefunded development accounts.

To persist chain state between runs, specify a base path by running a command similar to the following:

```sh
// Create a folder to use as the db base path
$ mkdir my-chain-state

// Use of that folder to store the chain state
$ ./target/release/node-template --dev --base-path ./my-chain-state/

// Check the folder structure created inside the base path after running the chain
$ ls ./my-chain-state
chains
$ ls ./my-chain-state/chains/
dev
$ ls ./my-chain-state/chains/dev
db keystore network
```

### Connect with Polkadot-JS Apps Front-End

After you start the node template locally, you can interact with it using the hosted version of the [Polkadot/Substrate Portal](https://polkadot.js.org/apps/#/explorer?rpc=ws://localhost:9944) front-end by connecting to the local node endpoint.
A hosted version is also available on [IPFS (redirect) here](https://dotapps.io/) or [IPNS (direct) here](ipns://dotapps.io/?rpc=ws%3A%2F%2F127.0.0.1%3A9944#/explorer).
You can also find the source code and instructions for hosting your own instance on the [polkadot-js/apps](https://github.com/polkadot-js/apps) repository.

### Multi-Node Local Testnet

If you want to see the multi-node consensus algorithm in action, see [Simulate a network](https://docs.substrate.io/tutorials/get-started/simulate-network/).

## Template Structure

A Substrate project such as this consists of a number of components that are spread across a few directories.

### Node

A blockchain node is an application that allows users to participate in a blockchain network.
Substrate-based blockchain nodes expose a number of capabilities:

- Networking: Substrate nodes use the [`libp2p`](https://libp2p.io/) networking stack to allow the
  nodes in the network to communicate with one another.
- Consensus: Blockchains must have a way to come to [consensus](https://docs.substrate.io/fundamentals/consensus/) on the state of the network.
  Substrate makes it possible to supply custom consensus engines and also ships with several consensus mechanisms that have been built on top of [Web3 Foundation research](https://research.web3.foundation/en/latest/polkadot/NPoS/index.html).
- RPC Server: A remote procedure call (RPC) server is used to interact with Substrate nodes.

There are several files in the `node` directory.
Take special note of the following:

- [`chain_spec.rs`](./node/src/chain_spec.rs): A [chain specification](https://docs.substrate.io/build/chain-spec/) is a source code file that defines a Substrate chain''s initial (genesis) state.
  Chain specifications are useful for development and testing, and critical when architecting the launch of a production chain.
  Take note of the `development_config` and `testnet_genesis` functions.
  These functions are used to define the genesis state for the local development chain configuration.
  These functions identify some [well-known accounts](https://docs.substrate.io/reference/command-line-tools/subkey/) and use them to configure the blockchain''s initial state.
- [`service.rs`](./node/src/service.rs): This file defines the node implementation.
  Take note of the libraries that this file imports and the names of the functions it invokes.
  In particular, there are references to consensus-related topics, such as the [block finalization and forks](https://docs.substrate.io/fundamentals/consensus/#finalization-and-forks) and other [consensus mechanisms](https://docs.substrate.io/fundamentals/consensus/#default-consensus-models) such as Aura for block authoring and GRANDPA for finality.

### Runtime

In Substrate, the terms "runtime" and "state transition function" are analogous.
Both terms refer to the core logic of the blockchain that is responsible for validating blocks and executing the state changes they define.
The Substrate project in this repository uses [FRAME](https://docs.substrate.io/fundamentals/runtime-development/#frame) to construct a blockchain runtime.
FRAME allows runtime developers to declare domain-specific logic in modules called "pallets".
At the heart of FRAME is a helpful [macro language](https://docs.substrate.io/reference/frame-macros/) that makes it easy to create pallets and flexibly compose them to create blockchains that can address [a variety of needs](https://substrate.io/ecosystem/projects/).

Review the [FRAME runtime implementation](./runtime/src/lib.rs) included in this template and note the following:

- This file configures several pallets to include in the runtime.
  Each pallet configuration is defined by a code block that begins with `impl $PALLET_NAME::Config for Runtime`.
- The pallets are composed into a single runtime by way of the [`construct_runtime!`](https://crates.parity.io/frame_support/macro.construct_runtime.html) macro, which is part of the core FRAME Support [system](https://docs.substrate.io/reference/frame-pallets/#system-pallets) library.

### Pallets

The runtime in this project is constructed using many FRAME pallets that ship with the [core Substrate repository](https://github.com/paritytech/substrate/tree/master/frame) and a template pallet that is [defined in the `pallets`](./pallets/template/src/lib.rs) directory.

A FRAME pallet is compromised of a number of blockchain primitives:

- Storage: FRAME defines a rich set of powerful [storage abstractions](https://docs.substrate.io/build/runtime-storage/) that makes it easy to use Substrate''s efficient key-value database to manage the evolving state of a blockchain.
- Dispatchables: FRAME pallets define special types of functions that can be invoked (dispatched) from outside of the runtime in order to update its state.
- Events: Substrate uses [events and errors](https://docs.substrate.io/build/events-and-errors/) to notify users of important changes in the runtime.
- Errors: When a dispatchable fails, it returns an error.
- Config: The `Config` configuration interface is used to define the types and parameters upon which a FRAME pallet depends.

## Alternative Installations

Instead of installing dependencies and building this source directly, consider the following alternatives.

### CI

#### Binary

Check the [CI release workflow](./.github/workflows/release.yml) to see how the binary is built on CI.
You can modify the compilation targets depending on your needs.

Allow GitHub actions in your forked repository to build the binary for you.

Push a tag. For example, `v0.1.1`. Based on [Semantic Versioning](https://semver.org/), the supported tag format is `v?MAJOR.MINOR.PATCH(-PRERELEASE)?(+BUILD_METADATA)?` (the leading "v", pre-release version, and build metadata are optional and the optional prefix is also supported).

After the pipeline is finished, you can download the binary from the releases page.

#### Container

Check the [CI release workflow](./.github/workflows/release.yml) to see how the Docker image is built on CI.

Add your `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN` secrets or other organization settings to your forked repository.
Change the `DOCKER_REPO` variable in the workflow to `[your DockerHub registry name]/[image name]`.

Push a tag.

After the image is built and pushed, you can pull it with `docker pull <DOCKER_REPO>:<tag>`.

### Nix

Install [nix](https://nixos.org/), and optionally [direnv](https://github.com/direnv/direnv) and [lorri](https://github.com/nix-community/lorri) for a fully plug-and-play experience for setting up the development environment.
To get all the correct dependencies, activate direnv `direnv allow` and lorri `lorri shell`.

### Docker

Please follow the [Substrate Docker instructions here](https://github.com/paritytech/substrate/blob/master/docker/README.md) to build the Docker container with the Substrate Node Template binary.
'
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
    image,
    deploy_type
) values (
    50,
    5,
    'icp_vue',
    'A Vue.js app, created with the Vue CLI.',
    1,
    '1.0.0',
    1,
    0,
    'https://vuejs.org/logo.svg',
    'https://g.alpha.hamsternet.io/ipfs/QmdhtgKNuQn2aqkdTn4DxRQidSLmmtggtpaBgh8vuyVjxd',
    3
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
) values (
    5,
    50,
    'icp_vue',
    '',
    'hamster-template',
    'https://github.com/hamster-template/icp_vue.git',
    'icp_vue',
    'main',
    '1.0.0',
    1,
    1,
    'https://g.alpha.hamsternet.io/ipfs/QmZMP5jmafwZDeE98Hu9kdT9wDuPXDs7ygpwFVzQRD7iqP',
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
);


create table t_user_icp (
    id int primary key auto_increment,
    fk_user_id int unique comment 'user id',
    identity_name varchar(50) unique comment 'identity name',
    account_id varchar(100) comment 'dfx ledger account-id',
    principal_id varchar(100)  comment 'dfx identity get-principal',
    wallet_id varchar(100)  comment 'dfx identity --network ic get-wallet',
    create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间',
    delete_time timestamp NULL comment '删除时间'
);

create table t_icp_canister (
    id int primary key auto_increment,
    project_id  char(36),
    canister_id varchar(50),
    canister_name varchar(50),
    cycles  decimal(10,2),
    status  tinyint,
    create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '更新时间'
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
    image,
    deploy_type
) values (
             53,
             5,
             'icp_react',
             'A client-side React appcreated with create-react-app.',
             1,
             '1.0.0',
             1,
             0,
             'https://reactjs.org/favicon.ico',
             'https://g.alpha.hamsternet.io/ipfs/QmUcXKszQxwT21dnqf67vFxzBgFwT8iWEsxExeGA1oFf6N',
             3
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
) values (
             6,
             53,
             'icp_react',
             '',
             'hamster-template',
             'https://github.com/hamster-template/icp_react.git',
             'icp_react',
             'main',
             '1.0.0',
             1,
             1,
             'https://g.alpha.hamsternet.io/ipfs/QmUcXKszQxwT21dnqf67vFxzBgFwT8iWEsxExeGA1oFf6N',
             '# Getting Started with Create React App

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

The page will reload when you make changes.\
You may also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.\
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can''t go back!**

If you aren''t satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you''re on your own.

You don''t have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn''t feel obligated to use this feature. However we understand that this tool wouldn''t be useful if you couldn''t customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).

### Code Splitting

This section has moved here: [https://facebook.github.io/create-react-app/docs/code-splitting](https://facebook.github.io/create-react-app/docs/code-splitting)

### Analyzing the Bundle Size

This section has moved here: [https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size](https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size)

### Making a Progressive Web App

This section has moved here: [https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app](https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app)

### Advanced Configuration

This section has moved here: [https://facebook.github.io/create-react-app/docs/advanced-configuration](https://facebook.github.io/create-react-app/docs/advanced-configuration)

### Deployment

This section has moved here: [https://facebook.github.io/create-react-app/docs/deployment](https://facebook.github.io/create-react-app/docs/deployment)

### `npm run build` fails to minify

This section has moved here: [https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify](https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify)
'
         );