<div align="center">
    <img src="https://socialify.git.ci/hamster-shared/hamster-develop/image?description=1&descriptionEditable=One-stop%20Toolkit%20and%20Middleware%20Platform%20for%20Web3.0%20Developers&font=KoHo&logo=https%3A%2F%2Fhamsternet.io%2F_nuxt%2Flogo.668de5a2.png&owner=1&pattern=Floating%20Cogs&theme=Auto" width="640" height="320" alt="logo" />

# <a href="https://develop.alpha.hamsternet.io/projects">Hamster Develop</a>

[![Discord](https://badgen.net/badge/icon/discord?icon=discord&label)](https://discord.gg/qMWUvs7jkV)
[![Telegram](https://badgen.net/badge/icon/telegram?icon=telegram&label)](https://t.me/hamsternetio)
[![Twitter](https://badgen.net/badge/icon/twitter?icon=twitter&label)](https://twitter.com/Hamsternetio)

_One-stop Toolkit and Middleware Platform for Web3.0 Developers_

</div>

## Getting Started

Follow these steps to set up and run the Go project locally:

### Prerequisites

Ensure you have the following installed on your system:

- [Go](https://golang.org/dl/) (latest version)
- [Docker](https://www.docker.com/get-started) (latest version)
- [Docker Compose](https://docs.docker.com/compose/install/) (latest version)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/hamster-shared/hamster-develop.git
cd hamster-develop
```

### Setting Up the Database

1. Start the database container using Docker Compose:

```bash
docker-compose up -d
```

2. Import the SQL files from the `pkg/db/migration` directory into the database:

- You can use your preferred database management tool (e.g., [DBeaver](https://dbeaver.io/), [MySQL Workbench](https://www.mysql.com/products/workbench/), or [pgAdmin](https://www.pgadmin.org/)) to connect to the database and import the SQL files.

### Building the Project

1. Depending on your system platform, run one of the following commands to compile the project:

- For Linux:

```bash
make linux
```

- For Windows:

```bash
make windows
```

- For macOS:

```bash
make macos
```

### Running the Project

1. Execute the project using the following command:

```bash
./aline daemon
```

Now the Go project should be running locally, and you can start exploring its features and functionalities.

## About Hamster

Hamster is aiming to build the one-stop infrastructure developer toolkits for Web3.0. It defines itself as a development, operation and maintenance DevOps service platform, providing a set of development tools as well as O&M tools, empowering projects in Web3.0 to improve their coding and delivery speed, quality and efficiency, as well as product reliability & safety.

With Hamster, developers or project teams realize the development, verification and O&M stages of their blockchain projects in an automatic, standardized and tooled approach: from contract template of multiple chains, contract/frontend code build, security check, contract deployment to the contract operation and maintenance.

Together with its developer toolkits, Hamster offers the RPC service and decentralized computing power network service when users choose to deploy their contracts via Hamster.

At the same time, the contract security check part within the developer toolkits is offered separately to its to-C customers, who could check their contracts to avoid potential security risks.

## Contributors

This project exists thanks to all the people who contribute.

 <a href="https://github.com/hamster-shared/hamster-develop/contributors">
  <img src="https://contrib.rocks/image?repo=hamster-shared/hamster-develop" />
 </a>

## License

[MIT](LICENSE)
