create table t_contract_arrange
(
    id                 int auto_increment                     primary key,
    project_id         char(36)                               not null comment 'project id',
    version            varchar(50)                            not  null comment 'contract build version',
    original_arrange   text                                   not  null comment 'original arrange',
    create_time        timestamp default CURRENT_TIMESTAMP    not null comment 'create time',
    update_time        timestamp default CURRENT_TIMESTAMP    not null comment 'update time'
);

create table t_contract_arrange_execute
(
    id                      int auto_increment                     primary key,
    project_id              char(36)                               not null comment 'project id',
    version                 varchar(50)                            not null comment 'contract build version',
    fk_arrange_id           int                                    not null comment 'contract build version',
    network                 varchar(50)                            not null comment 'network',
    arrange_process_data    text                                       null comment 'Process data for choreographed execution',
    create_time             timestamp default CURRENT_TIMESTAMP    not null comment 'create time',
    update_time             timestamp default CURRENT_TIMESTAMP    not null comment 'update time'
);

create table t_contract_arrange_cache
(
    id                 int auto_increment                     primary key,
    project_id         char(36)                               not null comment 'project id',
    contract_id        char(36)                               not null comment 'contract id',
    contract_name      varchar(100)                           not null comment 'contract name',
    version            varchar(50)                            not null comment 'contract build version',
    original_arrange   text                                   not null comment 'original arrange',
    create_time        timestamp default CURRENT_TIMESTAMP    not null comment 'create time',
    update_time        timestamp default CURRENT_TIMESTAMP    not null comment 'update time'
);
alter table t_contract_arrange
    add arrange_contract_name text  comment 'contract name arranged';

create table t_chain_network
(
    id                 int auto_increment                     primary key,
    logo               varchar(100)                           not null comment 'logo',
    category           varchar(50)                            not null comment 'category',
    chain_id           varchar(50)                            not null comment 'chain id',
    chain_name         varchar(50)                            not null comment 'chain name',
    rpc_url            varchar(100)                           not null comment 'rpc url',
    symbol             varchar(50)                            not null comment 'symbol',
    block_explorer_url varchar(100)                               null comment 'block explorer url',
    decimals           int                                    not null comment 'decimals',
    create_time        timestamp default CURRENT_TIMESTAMP    not null comment 'create time',
    update_time        timestamp default CURRENT_TIMESTAMP    not null comment 'update time'
);
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (10, 'Arbitrum', '0xa4b1', 'Arbitrum One', 'https://arb1.arbitrum.io/rpc', 'ETH', 'https://arbiscan.io', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmVFE5fEFDDwjjXBoTSJtxPxwdMKp9ygjLUvgK9CXNbRaV');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (11, 'Arbitrum', '0xa4ba', 'Arbitrum Nova', 'https://nova.arbitrum.io/rpc', 'ETH', 'https://nova-explorer.arbitrum.io', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmXU6ErzpJ1kHkjbyqLvNsWbbKRtvRT4tmpWoRvFbtUBoj');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (12, 'Arbitrum', '0x66eed', 'Arbitrum Goerli', 'https://goerli-rollup.arbitrum.io/rpc', 'AGOR', 'https://goerli-rollup-explorer.arbitrum.io', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmVFE5fEFDDwjjXBoTSJtxPxwdMKp9ygjLUvgK9CXNbRaV');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (22, 'Base', '0x14a33', 'Base Goerli Testnet', 'https://goerli.base.org', 'ETH', 'https://goerli.basescan.org', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmPmq82MdQYcBGY5BqTBiVhZt6bzq1Dm1ZUaex64hByyzk');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (21, 'Base', '0x2105', 'Base', 'https://mainnet.base.org', 'ETH', 'https://basescan.org', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmPmq82MdQYcBGY5BqTBiVhZt6bzq1Dm1ZUaex64hByyzk');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (29, 'BNB Smart Chain', '0xcc', 'opBNB Mainnet', 'https://opbnb-mainnet-rpc.bnbchain.org', 'BNB', 'http://opbnbscan.com', 18, '2023-10-30 10:42:37', '2023-10-30 10:42:37', 'https://g.alpha.hamsternet.io/ipfs/QmVMLy4XxA6gcFmr6kLFCLJMhVkfBZqGxLp5dUDoy39CUp');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (28, 'BNB Smart Chain', '0x15eb', 'opBNB Testnet', 'https://opbnb-testnet-rpc.bnbchain.org', 'tBNB', 'http://testnet.opbnbscan.com', 18, '2023-10-30 10:41:53', '2023-10-30 10:41:53', 'https://g.alpha.hamsternet.io/ipfs/QmVMLy4XxA6gcFmr6kLFCLJMhVkfBZqGxLp5dUDoy39CUp');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (27, 'BNB Smart Chain', '0x15e0', 'Greenfield Mekong Testnet', 'https://gnfd-testnet-fullnode-tendermint-us.bnbchain.org', 'tBNB', 'https://greenfieldscan.com', 18, '2023-10-30 10:40:50', '2023-10-30 10:40:50', 'https://g.alpha.hamsternet.io/ipfs/QmVMLy4XxA6gcFmr6kLFCLJMhVkfBZqGxLp5dUDoy39CUp');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (26, 'BNB Smart Chain', '0x3F9', 'Greenfield Mainnet', 'https://greenfield-chain.bnbchain.org', 'BNB', 'https://greenfieldscan.com', 18, '2023-10-30 10:24:14', '2023-10-30 10:24:14', 'https://g.alpha.hamsternet.io/ipfs/QmVMLy4XxA6gcFmr6kLFCLJMhVkfBZqGxLp5dUDoy39CUp');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (8, 'BNB Smart Chain', '0x38', 'BNB Smart Chain Mainnet', 'https://bsc-dataseed.binance.org/', 'BNB', 'https://bscscan.com', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmVMLy4XxA6gcFmr6kLFCLJMhVkfBZqGxLp5dUDoy39CUp');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (9, 'BNB Smart Chain', '0x61', 'BNB Smart Chain Testnet', 'https://bsc-testnet.publicnode.com', 'tBNB', 'https://testnet.bscscan.com', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmVMLy4XxA6gcFmr6kLFCLJMhVkfBZqGxLp5dUDoy39CUp');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (24, 'Conflux', '0x47', 'Conflux eSpace (Testnet)', 'https://evmtestnet.confluxrpc.com', 'CFX', 'https://evmtestnet.confluxscan.net', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmWXVeJc737QsfytApbdDZ1mb3ZN63nhmRyywEN7oSMpDJ');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (23, 'Conflux', '0x406', 'Conflux eSpace', 'https://evm.confluxrpc.com', 'CFX', 'https://evm.confluxscan.net', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmWXVeJc737QsfytApbdDZ1mb3ZN63nhmRyywEN7oSMpDJ');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (3, 'Ethereum', '0x5', 'Goerli', 'https://rpc.ankr.com/eth_goerli', 'ETH', 'https://goerli.etherscan.io', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmaXCw62BcjT74vgCcGNdd9mbm3KojW1KPLBGQ3JwRjkBd');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (5, 'Ethereum', '0x501', 'Hamster', 'https://rpc-moonbeam.hamster.newtouch.com', 'HM', null, 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmXgBwkCLx6oTwkGgpEJFd43pqKjUPsNx7xRrq8NDPd1dC');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (4, 'Ethereum', '0xaa36a7', 'Sepolia', 'https://ethereum-sepolia.publicnode.com', 'ETH', 'https://sepolia.etherscan.io', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmaXCw62BcjT74vgCcGNdd9mbm3KojW1KPLBGQ3JwRjkBd');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (1, 'Ethereum', '0x1', 'Ethereum Mainnet', 'https://rpc.ankr.com/eth', 'ETH', 'https://etherscan.io', 18, '2023-10-23 17:58:27', '2023-10-23 17:58:27', 'https://g.alpha.hamsternet.io/ipfs/QmaXCw62BcjT74vgCcGNdd9mbm3KojW1KPLBGQ3JwRjkBd');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (16, 'Filecoin', '0x4cb2f', 'Filecoin - Calibration testnet', 'https://api.calibration.node.glif.io/rpc/v1', 'tFIL', 'https://calibration.filscan.io', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmSVbCae4P2M5dFe8UrYG6nSdDRBhjvFcKhYd9zfD7Gpqx');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (15, 'Filecoin', '0x13a', 'Filecoin - Mainnet', 'https://api.node.glif.io', 'FIL', 'https://filfox.info/en', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmSVbCae4P2M5dFe8UrYG6nSdDRBhjvFcKhYd9zfD7Gpqx');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (14, 'IRIShub', '0x4130', 'IRIShub Testnet', 'https://evmrpc.nyancat.irisnet.org', 'ERIS', 'https://nyancat.iobscan.io', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmQapkNxQiffmr5H9JZs4JcnuXd8Y1zaCdQkNYHRwgMeQz');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (13, 'IRIShub', '0x1a20', 'IRIShub', 'https://evmrpc.irishub-1.irisnet.org', 'ERIS', 'https://irishub.iobscan.io', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmQapkNxQiffmr5H9JZs4JcnuXd8Y1zaCdQkNYHRwgMeQz');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (19, 'Linea', '0xe708', 'Linea', 'https://rpc.linea.build', 'ETH', 'https://lineascan.build', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmPJTikU15M48FFMFKRuVZMsft7GBYUGtNKmgHjhXkmnvp');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (20, 'Linea', '0xe704', 'Linea Testnet', 'https://rpc.goerli.linea.build', 'ETH', 'https://goerli.lineascan.build', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmPJTikU15M48FFMFKRuVZMsft7GBYUGtNKmgHjhXkmnvp');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (7, 'Polygon', '0x13881', 'Mumbai', 'https://rpc-mumbai.maticvigil.com', 'MATIC', 'https://mumbai.polygonscan.com', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmZQbRqxkw1gMeZ7beviU4jBmpT6xvxnRra7nnsnG6fNmJ');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (6, 'Polygon', '0x89', 'Polygon Mainnet', 'https://polygon-rpc.com/', 'MATIC', 'https://polygonscan.com', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmZQbRqxkw1gMeZ7beviU4jBmpT6xvxnRra7nnsnG6fNmJ');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (17, 'Scroll', '0x82751', 'Scroll Alpha Testnet', 'https://alpha-rpc.scroll.io/l2', 'ETH', 'https://alpha-blockscout.scroll.io', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmYWo3mGghmD8wUNwDdb85Z9Uysxmata55xVGU53rfgxJa');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (18, 'Scroll', '0x8274f', 'Scroll Sepolia Testnet', 'https://sepolia-rpc.scroll.io', 'ETH', 'https://sepolia-blockscout.scroll.io', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmYWo3mGghmD8wUNwDdb85Z9Uysxmata55xVGU53rfgxJa');
INSERT INTO t_chain_network (id, category, chain_id, chain_name, rpc_url, symbol, block_explorer_url, decimals, create_time, update_time, logo) VALUES (25, 'ZetaChain', '0x1b59', 'ZetaChain Athens-3 Testnet', 'https://zetachain-athens-evm.blockpi.network/v1/rpc/public', 'aZETA', 'https://zetachain-athens-3.blockscout.com/', 18, '2023-10-23 18:57:48', '2023-10-23 18:57:48', 'https://g.alpha.hamsternet.io/ipfs/QmTQkRfZeDdWEdLPeeFFEvGfr116UbcPxrKnzyazTCkbYW');
