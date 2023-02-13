-- 1:NFTS;2:DeFi;3:GameFi;4:DAOS    template_type_id
-- 1:EVM;2:Aptos;3:Ton;4:StarkWare  language_type
insert into t_template (
    template_type_id,
    name,
    description,
    audited,
    last_version,
    whether_display,
    language_type,
    logo
)
values (
        1,
        'Multiwrap',
        'Bundle multiple ERC721/ERC1155/ERC20 tokens into a single ERC721.',
        1,
        '1.0.0',
        0,
        1,
        ''
      ),
      (
       1,
       'Pack',
       'Pack multiple tokens into ERC1155 NFTs that act as randomized loot boxes',
       1,
       '1.0.0',
       0,
       1,
       ''
      ),
      (
       2,
       'rTokens',
       'rTokens are the tokens exchangeable for their underlying ones (e.g. rDai -> Dai).',
       1,
       '1.0.0',
       0,
       1,
       ''
      ),
      (
       2,
       'Quadratic Funding',
       'Funding round that uses quadratic matching (capital-constrained liberal radicalism)',
       1,
       '1.0.0',
       0,
       1,
       ''
      ),
      (
       2,
       'Uniswapper',
       'A component for swapping erc20s on Uniswap (plus tokenlists + local forks of mainnet!)',
       1,
       '1.0.0',
       0,
       1,
       ''
      ),
      (
       3,
       'Defi',
       'Create Contract for Defi.',
       1,
       '1.0.0',
       0,
       1,
       'https://static.devops.hamsternet.io/ipfs/QmT2gTnstKi95me3J42JjfxdgSR6BVmswkZbuhpLeWMGLZ'
      ),
      (
       4,
       'Vote',
       'Create and vote on proposals',
       1,
       '1.0.0',
       0,
       1,
       ''
      ),
      (
       4,
       'Split',
       'Distribute funds among multiple recipients',
       1,
       '1.0.0',
       0,
       1,
       ''
      ),
      (
       1,
       'aptos-token-staking',
       'The smart contract provides staking for Tokens and NFTs. The creators can decide the APR and way to distribute the gains.',
       1,
       '',
       1,
       2,
       'https://static.devops.hamsternet.io/ipfs/QmVPZy4GesehCxeLHijbnt1wSq1DJNWEAy244Dx3ZTHS8H'
      ),
      (
       1,
       'aptos-token-vesting',
       'Token Vesting Smart Contract for Aptos Blockchain.',
       1,
       '',
       1,
       2,
       'https://static.devops.hamsternet.io/ipfs/QmeVFcaVQYdpWZmx4UDdPKAiCRsh784vKnf57biuH3JFU5'
      ),
      (
       1,
       'nft-borrowing-lending-aptos',
       'Aptos Code for NFT borrow and lend',
       1,
       '',
       1,
       2,
       'https://static.devops.hamsternet.io/ipfs/QmaH6MQ4QmszZs37ZncRV8BaYacn3HPeWugiCkJ8XJrrpU'
      ),
      (
       2,
       'aptos-raffle',
       'Raffle in Aptos blockchain by Mokshya Protocol',
       1,
       '',
       1,
       2,
       'https://static.devops.hamsternet.io/ipfs/QmddQVuHLwvLS7qDnvLhUkKJZJC7d7ApcJMm5xsVWgjRwa'
      ),
      (
       1,
       'Non-Fungible tokens',
       'Basic implementation of smart contracts for NFT tokens and NFT collections in accordance with the Standard.',
       1,
       '',
       1,
       3,
       'https://static.devops.hamsternet.io/ipfs/QmPTqxymxKEGw5X9oNVGc38qqY24X4fGoDCdrKdfLU8LU1'
      ),
      (
       1,
       'Fungible tokens',
       'Basic implementation of smart contracts for Jetton wallet and Jetton minter in accordance with the Standard.',
       1,
       '',
       1,
       3,
       'https://static.devops.hamsternet.io/ipfs/QmPeStbFnayXcCDpBLfzdZVucTKDfrFhSybMcafvJu4t1u'
      ),
      (
       1,
       'TON DNS Smart Contracts',
       'Smart contracts of ".ton" zone',
       1,
       '',
       1,
       3,
       'https://static.devops.hamsternet.io/ipfs/QmUEn2zsMBW8xojsmaWrCRbi1K4x2Koz2K3hypACqDWT99'
      ),
      (
       1,
       'ERC721',
       'The ERC721 token standard is a specification for non-fungible tokens, or more colloquially: NFTs.',
       1,
       '',
       1,
       4,
       'https://static.devops.hamsternet.io/ipfs/QmUitnWFT9M7Kjmh9z4su4LGFfQsyqdCbGQ2vVYceoCDrj'
      ),
      (
       1,
       'ERC1155',
       'The ERC1155 multi token standard is a specification for fungibility-agnostic token contracts.',
       1,
       '',
       1,
       4,
       'https://static.devops.hamsternet.io/ipfs/QmYCLdgTEmLwVuecuzzXxC9aUhsPAW2cLd6TDi9QBEsQCn'
      ),
      (
       2,
       'ERC20',
       'The ERC20 preset offers a quick and easy setup for deploying a basic ERC20 token.',
       1,
       '',
       1,
       4,
       'https://static.devops.hamsternet.io/ipfs/QmNpaEzye5ipuD3vAqfCw7y95XgkBRMVaXc7SeyBvaCQhz'
      );
INSERT INTO t_template_detail (
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
)
VALUES(
          (
              SELECT id
              FROM t_template
              WHERE name = 'Multiwrap'
          ),
          'Multiwrap',
          '1',
          '',
          'Bundle multiple ERC721/ERC1155/ERC20 tokens into a single ERC721.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'Pack'
          ),
          'Pack',
          '1',
          '',
          'Pack multiple tokens into ERC1155 NFTs that act as randomized loot boxes.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'rTokens'
          ),
          'Pack',
          '1',
          '',
          'rTokens are the tokens exchangeable for their underlying ones (e.g. rDai -> Dai).',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'Quadratic Funding'
          ),
          'Quadratic Funding',
          '1',
          '',
          'Funding round that uses quadratic matching (capital-constrained liberal radicalism).',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'Uniswapper'
          ),
          'Uniswapper',
          '1',
          '',
          'A component for swapping erc20s on Uniswap (plus tokenlists + local forks of mainnet!).',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'Defi'
          ),
          'Defi',
          '1',
          '',
          'Create Contract for Defi.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'Vote'
          ),
          'Vote',
          '1',
          '',
          'Create and vote on proposals.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'Split'
          ),
          'Split',
          '1',
          '',
          'Distribute funds among multiple recipients.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'aptos-token-staking'
          ),
          'aptos-token-staking',
          '1',
          '',
          'The smart contract provides staking for Tokens and NFTs. The creators can decide the APR and way to distribute the gains.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'aptos-token-vesting'
          ),
          'aptos-token-vesting',
          '1',
          '',
          'Token Vesting Smart Contract for Aptos Blockchain.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'nft-borrowing-lending-aptos'
          ),
          'nft-borrowing-lending-aptos',
          '1',
          '',
          'Aptos Code for NFT borrow and lend.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'aptos-raffle'
          ),
          'aptos-raffle',
          '1',
          '',
          'Raffle in Aptos blockchain by Mokshya Protocol.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'Non-Fungible tokens'
          ),
          'Non-Fungible tokens',
          '1',
          '',
          'Basic implementation of smart contracts for NFT tokens and NFT collections in accordance with the Standard.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'Fungible tokens'
          ),
          'Fungible tokens',
          '1',
          '',
          'Basic implementation of smart contracts for Jetton wallet and Jetton minter in accordance with the Standard.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'TON DNS Smart Contracts'
          ),
          'TON DNS Smart Contracts',
          '1',
          '',
          'Smart contracts of ".ton" zone.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'ERC721'
          ),
          'ERC721',
          '1',
          '',
          'The ERC721 token standard is a specification for non-fungible tokens, or more colloquially: NFTs.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'ERC1155'
          ),
          'ERC1155',
          '1',
          '',
          'The ERC1155 multi token standard is a specification for fungibility-agnostic token contracts.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      ),
      (
          (
              SELECT id
              FROM t_template
              WHERE name = 'ERC20'
          ),
          'ERC20',
          '1',
          '',
          'The ERC20 preset offers a quick and easy setup for deploying a basic ERC20 token.',
          '',
          '',
          '',
          '',
          'hamster-template',
          '',
          '',
          'main',
          '1.0.0',
          '',
          '',
          ''
      );