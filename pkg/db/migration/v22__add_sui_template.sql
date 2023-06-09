insert into t_template (
    id,
    template_type_id,
    name,
    description,
    audited,
    last_version,
    whether_display,
    language_type,
    logo
) values (
             43,
             1,
             'Defi',
             'Some Common Defi Use Cases for Sui Networks',
             1,
             '0.0.1',
             1,
             5,
             'https://explorer.sui.io/favicon32x32.png'
         ),
         (
          44,
          1,
          'Games',
          'Examples of toy games built on top of Sui!',
          1,
          '0.0.1',
          1,
          5,
          'https://explorer.sui.io/favicon32x32.png'
         ),
         (
          45,
          1,
          'FungibleTokens',
          'Some alternative token templates on the Sui network',
          1,
          '0.0.1',
          1,
          5,
          'https://explorer.sui.io/favicon32x32.png'
         );

INSERT INTO t_template_detail (
    id,
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
) values (
             43,
             43,
             'Defi',
             1,
             '',
             'Some Common Defi Use Cases for Sui Networks',
             'You could use the sui-defi contract to:
* FlashLoan: a flash loan is a loan that must be initiated and repaid during the same transaction. This implementation works for any currency type, and is a good illustration of the power of Move abilities and the "hot potato" design pattern.
* Escrow: an atomic swap leveraging an escrow agent that is trusted for liveness, but not safety (i.e., the agent cannot steal the goods being swapped).',
             '',
             '{
  "jsonrpc":"2.0",
  "result":{
    "data":{
      "objectId":"0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
      "version":"1",
      "digest":"HXXuyE7u6D6GxSRRgQzNRnnUufM3DRiexTtJnZq8Snxh",
      "content":{
        "dataType":"package",
        "disassembled":{
          "dev_pass":{
            "add_uses": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": [
                    "Drop"
                  ]
                }
              ],
              "parameters": [
                {
                  "TypeParameter": 0
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "dev_pass",
                      "name": "Subscription",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                },
                "U64"
              ],
              "return": []
            },
            "confirm_use": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": [
                    "Drop"
                  ]
                }
              ],
              "parameters": [
                {
                  "TypeParameter": 0
                },
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "dev_pass",
                    "name": "SingleUse",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                }
              ],
              "return": []
            },
            "destroy": {
              "visibility": "Public",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "dev_pass",
                    "name": "Subscription",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                }
              ],
              "return": []
            },
            "issue_subscription": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": [
                    "Drop"
                  ]
                }
              ],
              "parameters": [
                {
                  "TypeParameter": 0
                },
                "U64",
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": [
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "dev_pass",
                    "name": "Subscription",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                }
              ]
            },
            "transfer": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": [
                    "Drop"
                  ]
                }
              ],
              "parameters": [
                {
                  "TypeParameter": 0
                },
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "dev_pass",
                    "name": "Subscription",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                },
                "Address"
              ],
              "return": []
            },
            "use_pass": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "dev_pass",
                      "name": "Subscription",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                }
              ],
              "return": [
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "dev_pass",
                    "name": "SingleUse",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                }
              ]
            },
            "uses": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "dev_pass",
                      "name": "Subscription",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                }
              ],
              "return": [
                "U64"
              ]
            }
          },
          "escrow":{
            "create": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                },
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                }
              ],
              "parameters": [
                "Address",
                "Address",
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "object",
                    "name": "ID",
                    "typeArguments": []
                  }
                },
                {
                  "TypeParameter": 0
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": []
            },
            "return_to_sender": {
              "visibility": "Public",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                },
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                }
              ],
              "parameters": [
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "escrow",
                    "name": "EscrowedObj",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      },
                      {
                        "TypeParameter": 1
                      }
                    ]
                  }
                }
              ],
              "return": []
            },
            "swap": {
              "visibility": "Public",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                },
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                }
              ],
              "parameters": [
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "escrow",
                    "name": "EscrowedObj",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      },
                      {
                        "TypeParameter": 1
                      }
                    ]
                  }
                },
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "escrow",
                    "name": "EscrowedObj",
                    "typeArguments": [
                      {
                        "TypeParameter": 1
                      },
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                }
              ],
              "return": []
            }
          },
          "flash_lender":{
            "create": {
              "visibility": "Public",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                },
                "U64",
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": []
            },
            "deposit": {
              "visibility": "Public",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "FlashLender",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                },
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "AdminCap",
                      "typeArguments": []
                    }
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                }
              ],
              "return": []
            },
            "fee": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "FlashLender",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                }
              ],
              "return": [
                "U64"
              ]
            },
            "flash_lender_id": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "Receipt",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                }
              ],
              "return": [
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "object",
                    "name": "ID",
                    "typeArguments": []
                  }
                }
              ]
            },
            "loan": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "FlashLender",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                },
                "U64",
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": [
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                },
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "flash_lender",
                    "name": "Receipt",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                }
              ]
            },
            "max_loan": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "FlashLender",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                }
              ],
              "return": [
                "U64"
              ]
            },
            "new": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "balance",
                    "name": "Balance",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                },
                "U64",
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": [
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "flash_lender",
                    "name": "AdminCap",
                    "typeArguments": []
                  }
                }
              ]
            },
            "repay": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "FlashLender",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                },
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "flash_lender",
                    "name": "Receipt",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                }
              ],
              "return": []
            },
            "repay_amount": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "Receipt",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                }
              ],
              "return": [
                "U64"
              ]
            },
            "update_fee": {
              "visibility": "Public",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "FlashLender",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                },
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "AdminCap",
                      "typeArguments": []
                    }
                  }
                },
                "U64"
              ],
              "return": []
            },
            "withdraw": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "FlashLender",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        }
                      ]
                    }
                  }
                },
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "flash_lender",
                      "name": "AdminCap",
                      "typeArguments": []
                    }
                  }
                },
                "U64",
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": [
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 0
                      }
                    ]
                  }
                }
              ]
            }
          },
          "pool":{
            "add_liquidity": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "pool",
                      "name": "Pool",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2",
                          "module": "sui",
                          "name": "SUI",
                          "typeArguments": []
                        }
                      }
                    ]
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 1
                      }
                    ]
                  }
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": [
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                          "module": "pool",
                          "name": "LSP",
                          "typeArguments": [
                            {
                              "TypeParameter": 0
                            },
                            {
                              "TypeParameter": 1
                            }
                          ]
                        }
                      }
                    ]
                  }
                }
              ]
            },
            "add_liquidity_": {
              "visibility": "Private",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "pool",
                      "name": "Pool",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2",
                          "module": "sui",
                          "name": "SUI",
                          "typeArguments": []
                        }
                      }
                    ]
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 1
                      }
                    ]
                  }
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": []
            },
            "create_pool": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": [
                    "Drop"
                  ]
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "TypeParameter": 0
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 1
                      }
                    ]
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2",
                          "module": "sui",
                          "name": "SUI",
                          "typeArguments": []
                        }
                      }
                    ]
                  }
                },
                "U64",
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": [
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                          "module": "pool",
                          "name": "LSP",
                          "typeArguments": [
                            {
                              "TypeParameter": 0
                            },
                            {
                              "TypeParameter": 1
                            }
                          ]
                        }
                      }
                    ]
                  }
                }
              ]
            },
            "get_amounts": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "pool",
                      "name": "Pool",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                }
              ],
              "return": [
                "U64",
                "U64",
                "U64"
              ]
            },
            "get_input_price": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [],
              "parameters": [
                "U64",
                "U64",
                "U64",
                "U64"
              ],
              "return": [
                "U64"
              ]
            },
            "remove_liquidity": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "pool",
                      "name": "Pool",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                          "module": "pool",
                          "name": "LSP",
                          "typeArguments": [
                            {
                              "TypeParameter": 0
                            },
                            {
                              "TypeParameter": 1
                            }
                          ]
                        }
                      }
                    ]
                  }
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": [
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2",
                          "module": "sui",
                          "name": "SUI",
                          "typeArguments": []
                        }
                      }
                    ]
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 1
                      }
                    ]
                  }
                }
              ]
            },
            "remove_liquidity_": {
              "visibility": "Private",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "pool",
                      "name": "Pool",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                          "module": "pool",
                          "name": "LSP",
                          "typeArguments": [
                            {
                              "TypeParameter": 0
                            },
                            {
                              "TypeParameter": 1
                            }
                          ]
                        }
                      }
                    ]
                  }
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": []
            },
            "sui_price": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "pool",
                      "name": "Pool",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                "U64"
              ],
              "return": [
                "U64"
              ]
            },
            "swap_sui": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "pool",
                      "name": "Pool",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2",
                          "module": "sui",
                          "name": "SUI",
                          "typeArguments": []
                        }
                      }
                    ]
                  }
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": [
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 1
                      }
                    ]
                  }
                }
              ]
            },
            "swap_sui_": {
              "visibility": "Private",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "pool",
                      "name": "Pool",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2",
                          "module": "sui",
                          "name": "SUI",
                          "typeArguments": []
                        }
                      }
                    ]
                  }
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": []
            },
            "swap_token": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "pool",
                      "name": "Pool",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 1
                      }
                    ]
                  }
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": [
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2",
                          "module": "sui",
                          "name": "SUI",
                          "typeArguments": []
                        }
                      }
                    ]
                  }
                }
              ]
            },
            "swap_token_": {
              "visibility": "Private",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "pool",
                      "name": "Pool",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "coin",
                    "name": "Coin",
                    "typeArguments": [
                      {
                        "TypeParameter": 1
                      }
                    ]
                  }
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": []
            },
            "token_price": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "pool",
                      "name": "Pool",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                "U64"
              ],
              "return": [
                "U64"
              ]
            }
          },
          "shared_escrow":{
            "cancel": {
              "visibility": "Public",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                },
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                }
              ],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "shared_escrow",
                      "name": "EscrowedObj",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": []
            },
            "create": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                },
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                }
              ],
              "parameters": [
                "Address",
                {
                  "Struct": {
                    "address": "0x2",
                    "module": "object",
                    "name": "ID",
                    "typeArguments": []
                  }
                },
                {
                  "TypeParameter": 0
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": []
            },
            "exchange": {
              "visibility": "Public",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                },
                {
                  "abilities": [
                    "Store",
                    "Key"
                  ]
                }
              ],
              "parameters": [
                {
                  "TypeParameter": 1
                },
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "shared_escrow",
                      "name": "EscrowedObj",
                      "typeArguments": [
                        {
                          "TypeParameter": 0
                        },
                        {
                          "TypeParameter": 1
                        }
                      ]
                    }
                  }
                },
                {
                  "Reference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": []
            }
          },
          "smart_swapper":{
            "cross_pool_swap": {
              "visibility": "Private",
              "isEntry": true,
              "typeParameters": [],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "dev_pass",
                      "name": "Subscription",
                      "typeArguments": [
                        {
                          "Struct": {
                            "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                            "module": "some_amm",
                            "name": "DEVPASS",
                            "typeArguments": []
                          }
                        }
                      ]
                    }
                  }
                }
              ],
              "return": []
            }
          },
          "some_amm":{
            "dev_swap": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [
                {
                  "Struct": {
                    "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                    "module": "dev_pass",
                    "name": "SingleUse",
                    "typeArguments": [
                      {
                        "Struct": {
                          "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                          "module": "some_amm",
                          "name": "DEVPASS",
                          "typeArguments": []
                        }
                      }
                    ]
                  }
                }
              ],
              "return": [
                "Bool"
              ]
            },
            "purchase_pass": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2",
                      "module": "tx_context",
                      "name": "TxContext",
                      "typeArguments": []
                    }
                  }
                }
              ],
              "return": []
            },
            "swap": {
              "visibility": "Private",
              "isEntry": true,
              "typeParameters": [
                {
                  "abilities": []
                },
                {
                  "abilities": []
                }
              ],
              "parameters": [],
              "return": []
            },
            "topup_pass": {
              "visibility": "Public",
              "isEntry": false,
              "typeParameters": [],
              "parameters": [
                {
                  "MutableReference": {
                    "Struct": {
                      "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                      "module": "dev_pass",
                      "name": "Subscription",
                      "typeArguments": [
                        {
                          "Struct": {
                            "address": "0x2222f801a7bd5169b52bbc155ecd198a6fa174d284a8a97ef65d14d12427fcd7",
                            "module": "some_amm",
                            "name": "DEVPASS",
                            "typeArguments": []
                          }
                        }
                      ]
                    }
                  }
                }
              ],
              "return": []
            }
          }
        }
      }
    }
  },
  "id":"2"
}',
             '',
             'hamster-template',
             'https://github.com/hamster-template/sui-defi.git',
             'sui-defi',
             'main',
             '0.0.1',
             'https://api.github.com/repos/hamster-template/sui-defi/contents/sources?ref=main',
             '',
             ''
         ),
         (
            44,
            44,
            'Games',
            1,
            '',
            'Examples of toy games built on top of Sui!',
            'Examples of toy games built on top of Sui!
* TicTacToe: the pencil and paper classic, now on Sui, implemented using single-owner objects only.
* SharedTicTacToe: the pencil and paper classic, now on Sui, using shared objects.
* Hero: an adventure game where an intrepid hero slays vicious boars with a magic sword and heals himself with potions.
* SeaHero: a permissionless mod of the Hero game where the hero can slay sea monsters to earn RUM tokens.
* SeaHeroHelper: a permissionless mod of the economics of the Sea Hero game. A weak hero can request help from a stronger hero, who receives a share of the monster slaying reward.
* RockPaperScissors: a commit-reveal scheme in which players first submit their commitments and then reveal the data that led to these commitments.',
            '',
            '{
    "jsonrpc": "2.0",
    "result": {
        "data": {
            "objectId": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "version": "1",
            "digest": "FYmoxwLZEAbPU4LBcRF6CCpjSq8AqnWbyHiGx4v7fV3b",
            "content": {
                "dataType": "package",
                "disassembled": {
                    "drand_based_lottery":{
  "close": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_lottery",
            "name": "Game",
            "typeArguments": []
          }
        }
      },
      {
        "Vector": "U8"
      },
      {
        "Vector": "U8"
      }
    ],
    "return": []
  },
  "complete": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_lottery",
            "name": "Game",
            "typeArguments": []
          }
        }
      },
      {
        "Vector": "U8"
      },
      {
        "Vector": "U8"
      }
    ],
    "return": []
  },
  "create": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      "U64",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "delete_game_winner": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "drand_based_lottery",
          "name": "GameWinner",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "delete_ticket": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "drand_based_lottery",
          "name": "Ticket",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "get_game_winner_game_id": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_lottery",
            "name": "GameWinner",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Reference": {
          "Struct": {
            "address": "0x2",
            "module": "object",
            "name": "ID",
            "typeArguments": []
          }
        }
      }
    ]
  },
  "get_ticket_game_id": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_lottery",
            "name": "Ticket",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Reference": {
          "Struct": {
            "address": "0x2",
            "module": "object",
            "name": "ID",
            "typeArguments": []
          }
        }
      }
    ]
  },
  "participate": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_lottery",
            "name": "Game",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "redeem": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_lottery",
            "name": "Ticket",
            "typeArguments": []
          }
        }
      },
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_lottery",
            "name": "Game",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  }
},
"drand_based_scratch_card":{
  "buy_ticket": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x2",
                "module": "sui",
                "name": "SUI",
                "typeArguments": []
              }
            }
          ]
        }
      },
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_scratch_card",
            "name": "Game",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "create": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x2",
                "module": "sui",
                "name": "SUI",
                "typeArguments": []
              }
            }
          ]
        }
      },
      "U64",
      "U64",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "delete_ticket": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "drand_based_scratch_card",
          "name": "Ticket",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "end_of_game_round": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      "U64"
    ],
    "return": [
      "U64"
    ]
  },
  "evaluate": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "drand_based_scratch_card",
          "name": "Ticket",
          "typeArguments": []
        }
      },
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_scratch_card",
            "name": "Game",
            "typeArguments": []
          }
        }
      },
      {
        "Vector": "U8"
      },
      {
        "Vector": "U8"
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "get_game_base_drand_round": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_scratch_card",
            "name": "Game",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "get_game_base_epoch": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_scratch_card",
            "name": "Game",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "redeem": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_scratch_card",
            "name": "Reward",
            "typeArguments": []
          }
        }
      },
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_scratch_card",
            "name": "Game",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "take_reward": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "drand_based_scratch_card",
          "name": "Winner",
          "typeArguments": []
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "drand_based_scratch_card",
            "name": "Reward",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  }
},
"drand_lib":{
  "derive_randomness": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Vector": "U8"
      }
    ],
    "return": [
      {
        "Vector": "U8"
      }
    ]
  },
  "safe_selection": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      "U64",
      {
        "Reference": {
          "Vector": "U8"
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "verify_drand_signature": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Vector": "U8"
      },
      {
        "Vector": "U8"
      },
      "U64"
    ],
    "return": []
  },
  "verify_time_has_passed": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      "U64",
      {
        "Vector": "U8"
      },
      {
        "Vector": "U8"
      },
      "U64"
    ],
    "return": []
  }
},
"hero":{
  "acquire_hero": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "GameInfo",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x2",
                "module": "sui",
                "name": "SUI",
                "typeArguments": []
              }
            }
          ]
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "assert_hero_strength": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "Hero",
            "typeArguments": []
          }
        }
      },
      "U64"
    ],
    "return": []
  },
  "check_id": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "GameInfo",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "object",
          "name": "ID",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "create_hero": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "GameInfo",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "hero",
          "name": "Sword",
          "typeArguments": []
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "hero",
          "name": "Hero",
          "typeArguments": []
        }
      }
    ]
  },
  "create_sword": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "GameInfo",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x2",
                "module": "sui",
                "name": "SUI",
                "typeArguments": []
              }
            }
          ]
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "hero",
          "name": "Sword",
          "typeArguments": []
        }
      }
    ]
  },
  "equip_sword": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "Hero",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "hero",
          "name": "Sword",
          "typeArguments": []
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x1",
          "module": "option",
          "name": "Option",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
                "module": "hero",
                "name": "Sword",
                "typeArguments": []
              }
            }
          ]
        }
      }
    ]
  },
  "heal": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "Hero",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "hero",
          "name": "Potion",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "hero_strength": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "Hero",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "id": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "GameInfo",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x2",
          "module": "object",
          "name": "ID",
          "typeArguments": []
        }
      }
    ]
  },
  "new_game": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "remove_sword": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "Hero",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "hero",
          "name": "Sword",
          "typeArguments": []
        }
      }
    ]
  },
  "send_boar": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "GameInfo",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "GameAdmin",
            "typeArguments": []
          }
        }
      },
      "U64",
      "U64",
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "send_potion": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "GameInfo",
            "typeArguments": []
          }
        }
      },
      "U64",
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "GameAdmin",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "slay": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "GameInfo",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "Hero",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "hero",
          "name": "Boar",
          "typeArguments": []
        }
      },
      {
        "Reference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "sword_strength": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "Sword",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  }
},
"rock_paper_scissors":{
  "add_hash": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "rock_paper_scissors",
            "name": "Game",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "rock_paper_scissors",
          "name": "PlayerTurn",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "match_secret": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "rock_paper_scissors",
            "name": "Game",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "rock_paper_scissors",
          "name": "Secret",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "new_game": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      "Address",
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "paper": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [],
    "return": [
      "U8"
    ]
  },
  "player_turn": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      "Address",
      {
        "Vector": "U8"
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "reveal": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      "Address",
      {
        "Vector": "U8"
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "rock": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [],
    "return": [
      "U8"
    ]
  },
  "scissors": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [],
    "return": [
      "U8"
    ]
  },
  "select_winner": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "rock_paper_scissors",
          "name": "Game",
          "typeArguments": []
        }
      },
      {
        "Reference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "status": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "rock_paper_scissors",
            "name": "Game",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U8"
    ]
  }
},
"sea_hero":{
  "create_monster": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "sea_hero",
            "name": "SeaHeroAdmin",
            "typeArguments": []
          }
        }
      },
      "U64",
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "monster_reward": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "sea_hero",
            "name": "SeaMonster",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "slay": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "Hero",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "sea_hero",
          "name": "SeaMonster",
          "typeArguments": []
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x2",
          "module": "balance",
          "name": "Balance",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
                "module": "sea_hero",
                "name": "RUM",
                "typeArguments": []
              }
            }
          ]
        }
      }
    ]
  }
},
"sea_hero_helper":{
  "create": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "sea_hero",
          "name": "SeaMonster",
          "typeArguments": []
        }
      },
      "U64",
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "owner_reward": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "sea_hero_helper",
            "name": "HelpMeSlayThisMonster",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "return_to_owner": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "sea_hero_helper",
          "name": "HelpMeSlayThisMonster",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "slay": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "hero",
            "name": "Hero",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "sea_hero_helper",
          "name": "HelpMeSlayThisMonster",
          "typeArguments": []
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
                "module": "sea_hero",
                "name": "RUM",
                "typeArguments": []
              }
            }
          ]
        }
      }
    ]
  }
},
"shared_tic_tac_toe":{
  "create_game": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      "Address",
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "delete_game": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "shared_tic_tac_toe",
          "name": "TicTacToe",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "delete_trophy": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "shared_tic_tac_toe",
          "name": "Trophy",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "get_status": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "shared_tic_tac_toe",
            "name": "TicTacToe",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U8"
    ]
  },
  "place_mark": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "shared_tic_tac_toe",
            "name": "TicTacToe",
            "typeArguments": []
          }
        }
      },
      "U8",
      "U8",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  }
},
"tic_tac_toe":{
  "create_game": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      "Address",
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "delete_cap": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "tic_tac_toe",
          "name": "MarkMintCap",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "delete_game": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "tic_tac_toe",
          "name": "TicTacToe",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "delete_trophy": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "tic_tac_toe",
          "name": "Trophy",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "get_status": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "tic_tac_toe",
            "name": "TicTacToe",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U8"
    ]
  },
  "mark_col": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "tic_tac_toe",
            "name": "Mark",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "mark_player": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "tic_tac_toe",
            "name": "Mark",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Reference": "Address"
      }
    ]
  },
  "mark_row": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "tic_tac_toe",
            "name": "Mark",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "place_mark": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "tic_tac_toe",
            "name": "TicTacToe",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
          "module": "tic_tac_toe",
          "name": "Mark",
          "typeArguments": []
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "send_mark_to_game": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x98e436b86eb76c275eba3d18d9f96b5e7ec0a34979637e462aba554c87e35e60",
            "module": "tic_tac_toe",
            "name": "MarkMintCap",
            "typeArguments": []
          }
        }
      },
      "Address",
      "U64",
      "U64",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  }
}
                }
            }
        }
    },
    "id": "2"
}',
            '',
            'hamster-template',
            'https://github.com/hamster-template/sui_games.git',
            'sui_games',
            'main',
            '0.0.1',
            'https://api.github.com/repos/hamster-template/sui_games/contents/sources?ref=main',
            '',
            ''
         ),
         (
          45,
          45,
          'FungibleTokens',
          1,
          '',
          'Some alternative token templates on the Sui network',
          '* MANAGED: a token managed by a treasurer trusted for minting and burning. This is how (e.g.) a fiat-backed stablecoin or an in-game virtual currency would work.
* BASKET: a synthetic token backed by a basket of other assets. This how (e.g.) a Special Drawing Rights (SDR) like asset would work.
* REGULATED_COIN: a coin managed by a central authority which can freeze accounts
* PRIVATE_COIN: a coin which has the option of hiding the amount that you transact
* PRIVATE_BALANCE: a balance which has the option of hiding the amount that it stores.',
          '',
          '{
    "jsonrpc": "2.0",
    "result": {
        "data": {
            "objectId": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "version": "1",
            "digest": "37zRRP76UWoxfUtBnn5wwHHKkZaSVNEAVKVjpw1TsKe6",
            "content": {
                "dataType": "package",
                "disassembled": {
                   "abc":{
  "accept_transfer": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "abc",
            "name": "Registry",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "Struct": {
                  "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                  "module": "abc",
                  "name": "Abc",
                  "typeArguments": []
                }
              }
            ]
          }
        }
      },
      {
        "Struct": {
          "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
          "module": "abc",
          "name": "Transfer",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "ban": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "abc",
            "name": "AbcTreasuryCap",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "abc",
            "name": "Registry",
            "typeArguments": []
          }
        }
      },
      "Address"
    ],
    "return": []
  },
  "banned": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "abc",
            "name": "Registry",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Reference": {
          "Vector": "Address"
        }
      }
    ]
  },
  "burn": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "abc",
            "name": "AbcTreasuryCap",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "Struct": {
                  "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                  "module": "abc",
                  "name": "Abc",
                  "typeArguments": []
                }
              }
            ]
          }
        }
      },
      "U64"
    ],
    "return": []
  },
  "create": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "abc",
            "name": "AbcTreasuryCap",
            "typeArguments": []
          }
        }
      },
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "mint": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "abc",
            "name": "AbcTreasuryCap",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "Struct": {
                  "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                  "module": "abc",
                  "name": "Abc",
                  "typeArguments": []
                }
              }
            ]
          }
        }
      },
      "U64"
    ],
    "return": []
  },
  "put_back": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "abc",
            "name": "Registry",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "Struct": {
                  "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                  "module": "abc",
                  "name": "Abc",
                  "typeArguments": []
                }
              }
            ]
          }
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                "module": "abc",
                "name": "Abc",
                "typeArguments": []
              }
            }
          ]
        }
      },
      {
        "Reference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "swapped_amount": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "abc",
            "name": "Registry",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "take": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "abc",
            "name": "Registry",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "Struct": {
                  "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                  "module": "abc",
                  "name": "Abc",
                  "typeArguments": []
                }
              }
            ]
          }
        }
      },
      "U64",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "transfer": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "abc",
            "name": "Registry",
            "typeArguments": []
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "Struct": {
                  "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                  "module": "abc",
                  "name": "Abc",
                  "typeArguments": []
                }
              }
            ]
          }
        }
      },
      "U64",
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  }
},
"basket":{
  "burn": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "basket",
            "name": "Reserve",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                "module": "basket",
                "name": "BASKET",
                "typeArguments": []
              }
            }
          ]
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x2",
                "module": "sui",
                "name": "SUI",
                "typeArguments": []
              }
            }
          ]
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                "module": "managed",
                "name": "MANAGED",
                "typeArguments": []
              }
            }
          ]
        }
      }
    ]
  },
  "managed_supply": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "basket",
            "name": "Reserve",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "mint": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "basket",
            "name": "Reserve",
            "typeArguments": []
          }
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x2",
                "module": "sui",
                "name": "SUI",
                "typeArguments": []
              }
            }
          ]
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                "module": "managed",
                "name": "MANAGED",
                "typeArguments": []
              }
            }
          ]
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                "module": "basket",
                "name": "BASKET",
                "typeArguments": []
              }
            }
          ]
        }
      }
    ]
  },
  "sui_supply": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "basket",
            "name": "Reserve",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "total_supply": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "basket",
            "name": "Reserve",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  }
},
"managed":{
  "burn": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "coin",
            "name": "TreasuryCap",
            "typeArguments": [
              {
                "Struct": {
                  "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                  "module": "managed",
                  "name": "MANAGED",
                  "typeArguments": []
                }
              }
            ]
          }
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "Coin",
          "typeArguments": [
            {
              "Struct": {
                "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                "module": "managed",
                "name": "MANAGED",
                "typeArguments": []
              }
            }
          ]
        }
      }
    ],
    "return": []
  },
  "mint": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "coin",
            "name": "TreasuryCap",
            "typeArguments": [
              {
                "Struct": {
                  "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
                  "module": "managed",
                  "name": "MANAGED",
                  "typeArguments": []
                }
              }
            ]
          }
        }
      },
      "U64",
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  }
},
"regulated_coin":{
  "borrow": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": [
          "Drop"
        ]
      }
    ],
    "parameters": [
      {
        "TypeParameter": 0
      },
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      }
    ],
    "return": [
      {
        "Reference": {
          "Struct": {
            "address": "0x2",
            "module": "balance",
            "name": "Balance",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      }
    ]
  },
  "borrow_mut": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": [
          "Drop"
        ]
      }
    ],
    "parameters": [
      {
        "TypeParameter": 0
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      }
    ],
    "return": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "balance",
            "name": "Balance",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      }
    ]
  },
  "creator": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      }
    ],
    "return": [
      "Address"
    ]
  },
  "from_balance": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": [
          "Drop"
        ]
      }
    ],
    "parameters": [
      {
        "TypeParameter": 0
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "balance",
          "name": "Balance",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      },
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
          "module": "regulated_coin",
          "name": "RegulatedCoin",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      }
    ]
  },
  "into_balance": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": [
          "Drop"
        ]
      }
    ],
    "parameters": [
      {
        "TypeParameter": 0
      },
      {
        "Struct": {
          "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
          "module": "regulated_coin",
          "name": "RegulatedCoin",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x2",
          "module": "balance",
          "name": "Balance",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      }
    ]
  },
  "join": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": [
          "Drop"
        ]
      }
    ],
    "parameters": [
      {
        "TypeParameter": 0
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "Struct": {
          "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
          "module": "regulated_coin",
          "name": "RegulatedCoin",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      }
    ],
    "return": []
  },
  "split": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": [
          "Drop"
        ]
      }
    ],
    "parameters": [
      {
        "TypeParameter": 0
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      "Address",
      "U64",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
          "module": "regulated_coin",
          "name": "RegulatedCoin",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      }
    ]
  },
  "value": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "regulated_coin",
            "name": "RegulatedCoin",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      }
    ],
    "return": [
      "U64"
    ]
  },
  "zero": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": [
          "Drop"
        ]
      }
    ],
    "parameters": [
      {
        "TypeParameter": 0
      },
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
          "module": "regulated_coin",
          "name": "RegulatedCoin",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      }
    ]
  }
},
"treasury_lock":{
  "ban_mint_cap_id": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "LockAdminCap",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "TreasuryLock",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "object",
          "name": "ID",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "ban_mint_cap_id_": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "LockAdminCap",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "TreasuryLock",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "object",
          "name": "ID",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "create_and_transfer_mint_cap": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "LockAdminCap",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      "U64",
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "create_mint_cap": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "LockAdminCap",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      "U64",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
          "module": "treasury_lock",
          "name": "MintCap",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      }
    ]
  },
  "mint_and_transfer": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "TreasuryLock",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "MintCap",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      "U64",
      "Address",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "mint_balance": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "TreasuryLock",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "MintCap",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      "U64",
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x2",
          "module": "balance",
          "name": "Balance",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      }
    ]
  },
  "new_lock": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "TreasuryCap",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": [
      {
        "Struct": {
          "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
          "module": "treasury_lock",
          "name": "LockAdminCap",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      }
    ]
  },
  "new_lock_": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "Struct": {
          "address": "0x2",
          "module": "coin",
          "name": "TreasuryCap",
          "typeArguments": [
            {
              "TypeParameter": 0
            }
          ]
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "tx_context",
            "name": "TxContext",
            "typeArguments": []
          }
        }
      }
    ],
    "return": []
  },
  "treasury_cap_mut": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "LockAdminCap",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "TreasuryLock",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      }
    ],
    "return": [
      {
        "MutableReference": {
          "Struct": {
            "address": "0x2",
            "module": "coin",
            "name": "TreasuryCap",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      }
    ]
  },
  "unban_mint_cap_id": {
    "visibility": "Public",
    "isEntry": false,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "LockAdminCap",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "TreasuryLock",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "object",
          "name": "ID",
          "typeArguments": []
        }
      }
    ],
    "return": []
  },
  "unban_mint_cap_id_": {
    "visibility": "Public",
    "isEntry": true,
    "typeParameters": [
      {
        "abilities": []
      }
    ],
    "parameters": [
      {
        "Reference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "LockAdminCap",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "MutableReference": {
          "Struct": {
            "address": "0x59c7c7e8b6235d6ab05c0963dba7b6e143961ec78cece4408013c4d4e9fb49fb",
            "module": "treasury_lock",
            "name": "TreasuryLock",
            "typeArguments": [
              {
                "TypeParameter": 0
              }
            ]
          }
        }
      },
      {
        "Struct": {
          "address": "0x2",
          "module": "object",
          "name": "ID",
          "typeArguments": []
        }
      }
    ],
    "return": []
  }
}
                }
            }
        }
    },
    "id": "2"
}',
          '',
          'hamster-template',
          'https://github.com/hamster-template/sui_fungible_tokens.git',
          'sui_fungible_tokens',
          'main',
          '0.0.1',
          'https://api.github.com/repos/hamster-template/sui_fungible_tokens/contents/sources?ref=main',
          '',
          ''
         );
