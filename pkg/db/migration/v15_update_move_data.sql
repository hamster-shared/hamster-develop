update aline.t_template
    set last_version = '1.0.0'
where name = 'aptos-token-staking';

update aline.t_template
set last_version = '1.0.0'
where name = 'aptos-token-vesting';

update aline.t_template
set last_version = '1.0.0'
where name = 'nft-borrowing-lending-aptos';

update aline.t_template
set last_version = '1.0.0'
where name = 'aptos-raffle';

update t_template_detail
    set repository_url = 'https://github.com/hamster-template/aptos-token-staking.git',
        repository_name = 'aptos-token-staking',
        code_sources = 'https://raw.githubusercontent.com/hamster-template/aptos-token-staking/main/sources/token-staking.move',
        abi_info = '{
  "address": "0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a",
  "name": "tokenstaking",
  "friends": [],
  "exposed_functions": [
    {
      "name": "claim_reward",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [
        {
          "constraints": []
        }
      ],
      "params": [
        "&signer",
        "0x1::string::String",
        "0x1::string::String",
        "address"
      ],
      "return": []
    },
    {
      "name": "create_staking",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [
        {
          "constraints": []
        }
      ],
      "params": [
        "&signer",
        "u64",
        "0x1::string::String",
        "u64"
      ],
      "return": []
    },
    {
      "name": "creator_stop_staking",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "0x1::string::String"
      ],
      "return": []
    },
    {
      "name": "deposit_staking_rewards",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [
        {
          "constraints": []
        }
      ],
      "params": [
        "&signer",
        "0x1::string::String",
        "u64"
      ],
      "return": []
    },
    {
      "name": "stake_token",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "address",
        "0x1::string::String",
        "0x1::string::String",
        "u64",
        "u64"
      ],
      "return": []
    },
    {
      "name": "unstake_token",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [
        {
          "constraints": []
        }
      ],
      "params": [
        "&signer",
        "address",
        "0x1::string::String",
        "0x1::string::String",
        "u64"
      ],
      "return": []
    },
    {
      "name": "update_dpr",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "u64",
        "0x1::string::String"
      ],
      "return": []
    }
  ],
  "structs": [
    {
      "name": "MokshyaMoney",
      "is_native": false,
      "abilities": [],
      "generic_type_params": [],
      "fields": [
        {
          "name": "dummy_field",
          "type": "bool"
        }
      ]
    },
    {
      "name": "MokshyaReward",
      "is_native": false,
      "abilities": [
        "drop",
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "staker",
          "type": "address"
        },
        {
          "name": "token_name",
          "type": "0x1::string::String"
        },
        {
          "name": "collection",
          "type": "0x1::string::String"
        },
        {
          "name": "withdraw_amount",
          "type": "u64"
        },
        {
          "name": "treasury_cap",
          "type": "0x1::account::SignerCapability"
        },
        {
          "name": "start_time",
          "type": "u64"
        },
        {
          "name": "tokens",
          "type": "u64"
        }
      ]
    },
    {
      "name": "MokshyaStaking",
      "is_native": false,
      "abilities": [
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "collection",
          "type": "0x1::string::String"
        },
        {
          "name": "dpr",
          "type": "u64"
        },
        {
          "name": "state",
          "type": "bool"
        },
        {
          "name": "amount",
          "type": "u64"
        },
        {
          "name": "coin_type",
          "type": "address"
        },
        {
          "name": "treasury_cap",
          "type": "0x1::account::SignerCapability"
        }
      ]
    },
    {
      "name": "ResourceInfo",
      "is_native": false,
      "abilities": [
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "resource_map",
          "type": "0x1::simple_map::SimpleMap<0x1::string::String, address>"
        }
      ]
    }
  ]
}'
where id = 17;

update t_template_detail
set repository_url = 'https://github.com/hamster-template/aptos-token-vesting.git',
    repository_name = 'aptos-token-vesting',
    code_sources = 'https://raw.githubusercontent.com/hamster-template/aptos-token-vesting/main/sources/token-vesting.move',
    abi_info = '{
  "address": "0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a",
  "name": "vesting",
  "friends": [],
  "exposed_functions": [
    {
      "name": "create_vesting",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [
        {
          "constraints": []
        }
      ],
      "params": [
        "&signer",
        "address",
        "vector<u64>",
        "vector<u64>",
        "u64",
        "vector<u8>"
      ],
      "return": []
    },
    {
      "name": "release_fund",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [
        {
          "constraints": []
        }
      ],
      "params": [
        "&signer",
        "address",
        "vector<u8>"
      ],
      "return": []
    }
  ],
  "structs": [
    {
      "name": "VestingCap",
      "is_native": false,
      "abilities": [
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "vestingMap",
          "type": "0x1::simple_map::SimpleMap<vector<u8>, address>"
        }
      ]
    },
    {
      "name": "VestingSchedule",
      "is_native": false,
      "abilities": [
        "store",
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "sender",
          "type": "address"
        },
        {
          "name": "receiver",
          "type": "address"
        },
        {
          "name": "coin_type",
          "type": "address"
        },
        {
          "name": "release_times",
          "type": "vector<u64>"
        },
        {
          "name": "release_amounts",
          "type": "vector<u64>"
        },
        {
          "name": "total_amount",
          "type": "u64"
        },
        {
          "name": "resource_cap",
          "type": "0x1::account::SignerCapability"
        },
        {
          "name": "released_amount",
          "type": "u64"
        }
      ]
    }
  ]
}'
where id = 18;

update t_template_detail
set repository_url = 'https://github.com/hamster-template/nft-borrowing-lending-aptos.git',
    repository_name = 'nft-borrowing-lending-aptos',
    code_sources = 'https://raw.githubusercontent.com/hamster-template/nft-borrowing-lending-aptos/main/sources/nftborrowlend.move',
    abi_info = '{
  "address": "0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a",
  "name": "borrowlend",
  "friends": [],
  "exposed_functions": [
    {
      "name": "borrow_select",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "0x1::string::String",
        "0x1::string::String",
        "u64",
        "address"
      ],
      "return": []
    },
    {
      "name": "borrower_pay_loan",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "0x1::string::String",
        "0x1::string::String"
      ],
      "return": []
    },
    {
      "name": "initiate_create_pool",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "address",
        "0x1::string::String",
        "u64",
        "u64"
      ],
      "return": []
    },
    {
      "name": "lender_claim_nft",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "0x1::string::String",
        "0x1::string::String"
      ],
      "return": []
    },
    {
      "name": "lender_offer",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "0x1::string::String",
        "u64",
        "u64"
      ],
      "return": []
    },
    {
      "name": "lender_offer_cancel",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "0x1::string::String"
      ],
      "return": []
    },
    {
      "name": "update_pool",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "0x1::string::String",
        "u64",
        "u64",
        "bool"
      ],
      "return": []
    }
  ],
  "structs": [
    {
      "name": "Amount",
      "is_native": false,
      "abilities": [
        "drop",
        "store"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "offer_per_nft",
          "type": "u64"
        },
        {
          "name": "number_of_offers",
          "type": "u64"
        },
        {
          "name": "total_amount",
          "type": "u64"
        }
      ]
    },
    {
      "name": "Borrower",
      "is_native": false,
      "abilities": [
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "borrows",
          "type": "0x1::table::Table<0x1::string::String, 0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a::borrowlend::Loan>"
        }
      ]
    },
    {
      "name": "CollectionPool",
      "is_native": false,
      "abilities": [
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "collection",
          "type": "0x1::string::String"
        },
        {
          "name": "creator",
          "type": "address"
        },
        {
          "name": "dpr",
          "type": "u64"
        },
        {
          "name": "days",
          "type": "u64"
        },
        {
          "name": "state",
          "type": "bool"
        },
        {
          "name": "treasury_cap",
          "type": "0x1::account::SignerCapability"
        },
        {
          "name": "offer",
          "type": "0x1::table::Table<address, 0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a::borrowlend::Amount>"
        },
        {
          "name": "loans",
          "type": "0x1::table::Table<0x1::string::String, 0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a::borrowlend::Loan>"
        }
      ]
    },
    {
      "name": "Lender",
      "is_native": false,
      "abilities": [
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "offers",
          "type": "0x1::table::Table<0x1::string::String, 0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a::borrowlend::Amount>"
        },
        {
          "name": "lends",
          "type": "0x1::table::Table<0x1::string::String, 0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a::borrowlend::Loan>"
        }
      ]
    },
    {
      "name": "Loan",
      "is_native": false,
      "abilities": [
        "copy",
        "drop",
        "store"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "borrower",
          "type": "address"
        },
        {
          "name": "lender",
          "type": "address"
        },
        {
          "name": "collection_name",
          "type": "0x1::string::String"
        },
        {
          "name": "token_name",
          "type": "0x1::string::String"
        },
        {
          "name": "property_version",
          "type": "u64"
        },
        {
          "name": "start_time",
          "type": "u64"
        },
        {
          "name": "dpr",
          "type": "u64"
        },
        {
          "name": "amount",
          "type": "u64"
        },
        {
          "name": "days",
          "type": "u64"
        }
      ]
    },
    {
      "name": "PoolMap",
      "is_native": false,
      "abilities": [
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "pools",
          "type": "0x1::table::Table<0x1::string::String, address>"
        }
      ]
    }
  ]
}'
where id = 19;

update t_template_detail
set repository_url = 'https://github.com/hamster-template/aptos-raffle.git',
    repository_name = 'aptos-raffle',
    code_sources = 'https://raw.githubusercontent.com/hamster-template/aptos-raffle/main/sources/raffle.move',
    abi_info = '{
  "address": "0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a",
  "name": "raffle",
  "friends": [],
  "exposed_functions": [
    {
      "name": "check_claim_reward",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "0x1::string::String",
        "0x1::string::String"
      ],
      "return": []
    },
    {
      "name": "declare_winner",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [
        {
          "constraints": []
        }
      ],
      "params": [
        "&signer",
        "0x1::string::String",
        "0x1::string::String"
      ],
      "return": []
    },
    {
      "name": "play_raffle",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [
        {
          "constraints": []
        }
      ],
      "params": [
        "&signer",
        "0x1::string::String",
        "0x1::string::String",
        "u64"
      ],
      "return": []
    },
    {
      "name": "start_raffle",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [
        {
          "constraints": []
        }
      ],
      "params": [
        "&signer",
        "address",
        "0x1::string::String",
        "0x1::string::String",
        "u64",
        "u64",
        "u64",
        "u64",
        "u64"
      ],
      "return": []
    }
  ],
  "structs": [
    {
      "name": "GameMap",
      "is_native": false,
      "abilities": [
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "raffles",
          "type": "0x1::table::Table<0x1::string::String, address>"
        }
      ]
    },
    {
      "name": "Player",
      "is_native": false,
      "abilities": [
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "raffles",
          "type": "0x1::table::Table<0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a::raffle::RaffleTokenInfo, vector<u64>>"
        }
      ]
    },
    {
      "name": "Raffle",
      "is_native": false,
      "abilities": [
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "raffle_creator",
          "type": "address"
        },
        {
          "name": "token_info",
          "type": "0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a::raffle::RaffleTokenInfo"
        },
        {
          "name": "start_time",
          "type": "u64"
        },
        {
          "name": "end_time",
          "type": "u64"
        },
        {
          "name": "num_tickets",
          "type": "u64"
        },
        {
          "name": "ticket_prize",
          "type": "u64"
        },
        {
          "name": "winning_ticket",
          "type": "u64"
        },
        {
          "name": "coin_type",
          "type": "address"
        },
        {
          "name": "treasury_cap",
          "type": "0x1::account::SignerCapability"
        },
        {
          "name": "ticket_sold",
          "type": "u64"
        },
        {
          "name": "winner",
          "type": "address"
        }
      ]
    },
    {
      "name": "RaffleTokenInfo",
      "is_native": false,
      "abilities": [
        "copy",
        "drop",
        "store"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "token_creator",
          "type": "address"
        },
        {
          "name": "collection",
          "type": "0x1::string::String"
        },
        {
          "name": "token_name",
          "type": "0x1::string::String"
        },
        {
          "name": "property_version",
          "type": "u64"
        }
      ]
    }
  ]
}'
where id = 20;
