update t_template set whether_display = 0 where id = 2;



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
    label_display
) values (
            81,
            6,
            'GreenfieldMarketplaceContracts',
            'Greenfield-data-marketplace is a marketplace protocol for safely and efficiently buying and selling data uploaded in Greenfield.',
            1,
            '0.0.1',
            1,
            1,
            'https://g.alpha.hamsternet.io/ipfs/QmdtdBXqnqimKELbXxT83QgSx5urygbCcYobRqGJSFGxu1',
            'Greenfield'
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
    author,
    repository_url,
    repository_name,
    branch,
    version,
    code_sources,
    title,
    title_description,
    how_use_description,
    label_display,
    abi_info,
    byte_code
)values (
            81,
            81,
            'GreenfieldMarketplaceContracts',
            1,
            '',
            'Greenfield-data-marketplace is a marketplace protocol for safely and efficiently buying and selling data uploaded in Greenfield.',
            '',
            '',
            'hamster-template',
            'https://github.com/hamster-template/greenfield-data-marketplace-contracts.git',
            'greenfield-data-marketplace-contracts',
            'main',
            '0.0.1',
            'https://raw.githubusercontent.com/hamster-template/greenfield-data-marketplace-contracts/main/contracts/deployer.sol',
            '',
            '',
            '',
            'Greenfield',
            '[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"_deployer","type":"address"},{"internalType":"uint8","name":"_nonce","type":"uint8"}],"name":"calcCreateAddress","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"address","name":"_implMarketplace","type":"address"},{"internalType":"address","name":"_owner","type":"address"},{"internalType":"address","name":"_fundWallet","type":"address"},{"internalType":"uint256","name":"_tax","type":"uint256"},{"internalType":"uint256","name":"_callbackGasLimit","type":"uint256"},{"internalType":"uint8","name":"_failureHandleStrategy","type":"uint8"}],"name":"deploy","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"deployed","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"implMarketplace","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"proxyAdmin","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"proxyMarketplace","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]',
            '0x608060405234801561001057600080fd5b5060408051606b60f91b6020808301829052602560fa1b602184018190523060601b6001600160601b03191660228501819052600160f81b603686015285516017818703018152603786018752805190840120600080546001600160a01b03199081166001600160a01b03938416178255605788019690965260588701939093526059860191909152600160f91b606d8601528551808603604e018152606e9095019586905284519490920193909320600180549093169116179055906100d69061015d565b604051809103906000f0801580156100f2573d6000803e3d6000fd5b506000549091506001600160a01b038083169116146101575760405162461bcd60e51b815260206004820152601a60248201527f696e76616c69642070726f787941646d696e2061646472657373000000000000604482015260640160405180910390fd5b5061016a565b6107118061170683390190565b61158d806101796000396000f3fe60806040523480156200001157600080fd5b50600436106200006a5760003560e01c80630a59d9ee146200006f5780633e47158c14620000a057806357663d4a14620000b45780638b212a6114620000c8578063f905c15a14620000e1578063fb69ee171462000107575b600080fd5b60015462000083906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b60005462000083906001600160a01b031681565b60025462000083906001600160a01b031681565b620000df620000d936600462000647565b62000181565b005b600254620000f690600160a01b900460ff1681565b604051901515815260200162000097565b6200008362000118366004620006bc565b604051606b60f91b6020820152602560fa1b60218201526bffffffffffffffffffffffff19606084901b16602282015260f882901b6001600160f81b031916603682015260009060370160408051601f1981840301815291905280516020909101209392505050565b600254600160a01b900460ff1615620001d55760405162461bcd60e51b81526020600482015260116024820152701bdb9b1e481b9bdd0819195c1b1bde5959607a1b60448201526064015b60405180910390fd5b6002805460ff60a01b1916600160a01b1790556001600160a01b038516620002305760405162461bcd60e51b815260206004820152600d60248201526c34b73b30b634b21037bbb732b960991b6044820152606401620001cc565b6001600160a01b0384166200027d5760405162461bcd60e51b81526020600482015260126024820152711a5b9d985b1a5908199d5b9915d85b1b195d60721b6044820152606401620001cc565b6103e8831115620002bf5760405162461bcd60e51b815260206004820152600b60248201526a0d2dcecc2d8d2c840e8c2f60ab1b6044820152606401620001cc565b60008211620003115760405162461bcd60e51b815260206004820152601860248201527f696e76616c69642063616c6c6261636b4761734c696d697400000000000000006044820152606401620001cc565b6001600160a01b0386163b6200036a5760405162461bcd60e51b815260206004820152601760248201527f696e76616c696420696d706c4d61726b6574706c6163650000000000000000006044820152606401620001cc565b600280546001600160a01b0319166001600160a01b0388811691821790925560008054604051919316906200039f9062000609565b6001600160a01b03928316815291166020820152606060408201819052600090820152608001604051809103906000f080158015620003e2573d6000803e3d6000fd5b506001549091506001600160a01b03808316911614620004455760405162461bcd60e51b815260206004820181905260248201527f696e76616c69642070726f78794d61726b6574706c61636520616464726573736044820152606401620001cc565b60005460405163f2fde38b60e01b81526001600160a01b0388811660048301529091169063f2fde38b90602401600060405180830381600087803b1580156200048d57600080fd5b505af1158015620004a2573d6000803e3d6000fd5b50505050856001600160a01b031660008054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000502573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620005289190620006f6565b6001600160a01b031614620005805760405162461bcd60e51b815260206004820152601860248201527f696e76616c69642070726f787941646d696e206f776e657200000000000000006044820152606401620001cc565b600154604051635b1987cd60e01b81526001600160a01b0388811660048301528781166024830152604482018790526064820186905260ff8516608483015290911690635b1987cd9060a401600060405180830381600087803b158015620005e757600080fd5b505af1158015620005fc573d6000803e3d6000fd5b5050505050505050505050565b610e3a806200071e83390190565b6001600160a01b03811681146200062d57600080fd5b50565b803560ff811681146200064257600080fd5b919050565b60008060008060008060c087890312156200066157600080fd5b86356200066e8162000617565b95506020870135620006808162000617565b94506040870135620006928162000617565b93506060870135925060808701359150620006b060a0880162000630565b90509295509295509295565b60008060408385031215620006d057600080fd5b8235620006dd8162000617565b9150620006ed6020840162000630565b90509250929050565b6000602082840312156200070957600080fd5b8151620007168162000617565b939250505056fe608060405260405162000e3a38038062000e3a833981016040819052620000269162000424565b828162000036828260006200004d565b50620000449050826200007f565b50505062000557565b6200005883620000f1565b600082511180620000665750805b156200007a5762000078838362000133565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f620000c160008051602062000df3833981519152546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1620000ee8162000162565b50565b620000fc8162000200565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606200015b838360405180606001604052806027815260200162000e136027913962000297565b9392505050565b6001600160a01b038116620001cd5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b8060008051602062000df38339815191525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6200026f5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401620001c4565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc620001df565b6060600080856001600160a01b031685604051620002b6919062000504565b600060405180830381855af49150503d8060008114620002f3576040519150601f19603f3d011682016040523d82523d6000602084013e620002f8565b606091505b5090925090506200030c8683838762000316565b9695505050505050565b606083156200038a57825160000362000382576001600160a01b0385163b620003825760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401620001c4565b508162000396565b6200039683836200039e565b949350505050565b815115620003af5781518083602001fd5b8060405162461bcd60e51b8152600401620001c4919062000522565b80516001600160a01b0381168114620003e357600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b838110156200041b57818101518382015260200162000401565b50506000910152565b6000806000606084860312156200043a57600080fd5b6200044584620003cb565b92506200045560208501620003cb565b60408501519092506001600160401b03808211156200047357600080fd5b818601915086601f8301126200048857600080fd5b8151818111156200049d576200049d620003e8565b604051601f8201601f19908116603f01168101908382118183101715620004c857620004c8620003e8565b81604052828152896020848701011115620004e257600080fd5b620004f5836020830160208801620003fe565b80955050505050509250925092565b6000825162000518818460208701620003fe565b9190910192915050565b602081526000825180602084015262000543816040850160208701620003fe565b601f01601f19169190910160400192915050565b61088c80620005676000396000f3fe60806040523661001357610011610017565b005b6100115b61001f610169565b6001600160a01b0316330361015f5760606001600160e01b0319600035166364d3180d60e11b810161005a5761005361019c565b9150610157565b63587086bd60e11b6001600160e01b031982160161007a576100536101f3565b63070d7c6960e41b6001600160e01b031982160161009a57610053610239565b621eb96f60e61b6001600160e01b03198216016100b95761005361026a565b63a39f25e560e01b6001600160e01b03198216016100d9576100536102aa565b60405162461bcd60e51b815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f78792074617267606482015261195d60f21b608482015260a4015b60405180910390fd5b815160208301f35b6101676102be565b565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b546001600160a01b0316919050565b60606101a66102ce565b60006101b53660048184610683565b8101906101c291906106c9565b90506101df816040518060200160405280600081525060006102d9565b505060408051602081019091526000815290565b60606000806102053660048184610683565b81019061021291906106fa565b91509150610222828260016102d9565b604051806020016040528060008152509250505090565b60606102436102ce565b60006102523660048184610683565b81019061025f91906106c9565b90506101df81610305565b60606102746102ce565b600061027e610169565b604080516001600160a01b03831660208201529192500160405160208183030381529060405291505090565b60606102b46102ce565b600061027e61035c565b6101676102c961035c565b61036b565b341561016757600080fd5b6102e28361038f565b6000825111806102ef5750805b15610300576102fe83836103cf565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61032e610169565b604080516001600160a01b03928316815291841660208301520160405180910390a1610359816103fb565b50565b60006103666104a4565b905090565b3660008037600080366000845af43d6000803e80801561038a573d6000f35b3d6000fd5b610398816104cc565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606103f4838360405180606001604052806027815260200161083060279139610560565b9392505050565b6001600160a01b0381166104605760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b606482015260840161014e565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80546001600160a01b0319166001600160a01b039290921691909117905550565b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61018d565b6001600160a01b0381163b6105395760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b606482015260840161014e565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc610483565b6060600080856001600160a01b03168560405161057d91906107e0565b600060405180830381855af49150503d80600081146105b8576040519150601f19603f3d011682016040523d82523d6000602084013e6105bd565b606091505b50915091506105ce868383876105d8565b9695505050505050565b60608315610647578251600003610640576001600160a01b0385163b6106405760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161014e565b5081610651565b6106518383610659565b949350505050565b8151156106695781518083602001fd5b8060405162461bcd60e51b815260040161014e91906107fc565b6000808585111561069357600080fd5b838611156106a057600080fd5b5050820193919092039150565b80356001600160a01b03811681146106c457600080fd5b919050565b6000602082840312156106db57600080fd5b6103f4826106ad565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561070d57600080fd5b610716836106ad565b9150602083013567ffffffffffffffff8082111561073357600080fd5b818501915085601f83011261074757600080fd5b813581811115610759576107596106e4565b604051601f8201601f19908116603f01168101908382118183101715610781576107816106e4565b8160405282815288602084870101111561079a57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60005b838110156107d75781810151838201526020016107bf565b50506000910152565b600082516107f28184602087016107bc565b9190910192915050565b602081526000825180602084015261081b8160408501602087016107bc565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212206cc9b3cf497bc227c075d3fbcf9f3adf44fb92e03c1bb28ec037a9f069fce2a464736f6c63430008150033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122050a25a7af7b4f77a84a3460d01d3cd54838ba0a215edd6e5ef97a0b7f2ebc94b64736f6c63430008150033608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6106938061007e6000396000f3fe60806040526004361061007b5760003560e01c80639623609d1161004e5780639623609d1461011157806399a88ec414610124578063f2fde38b14610144578063f3b7dead1461016457600080fd5b8063204e1c7a14610080578063715018a6146100bc5780637eff275e146100d35780638da5cb5b146100f3575b600080fd5b34801561008c57600080fd5b506100a061009b366004610499565b610184565b6040516001600160a01b03909116815260200160405180910390f35b3480156100c857600080fd5b506100d1610215565b005b3480156100df57600080fd5b506100d16100ee3660046104bd565b610229565b3480156100ff57600080fd5b506000546001600160a01b03166100a0565b6100d161011f36600461050c565b610291565b34801561013057600080fd5b506100d161013f3660046104bd565b610300565b34801561015057600080fd5b506100d161015f366004610499565b610336565b34801561017057600080fd5b506100a061017f366004610499565b6103b4565b6000806000836001600160a01b03166040516101aa90635c60da1b60e01b815260040190565b600060405180830381855afa9150503d80600081146101e5576040519150601f19603f3d011682016040523d82523d6000602084013e6101ea565b606091505b5091509150816101f957600080fd5b8080602001905181019061020d91906105e2565b949350505050565b61021d6103da565b6102276000610434565b565b6102316103da565b6040516308f2839760e41b81526001600160a01b038281166004830152831690638f283970906024015b600060405180830381600087803b15801561027557600080fd5b505af1158015610289573d6000803e3d6000fd5b505050505050565b6102996103da565b60405163278f794360e11b81526001600160a01b03841690634f1ef2869034906102c990869086906004016105ff565b6000604051808303818588803b1580156102e257600080fd5b505af11580156102f6573d6000803e3d6000fd5b5050505050505050565b6103086103da565b604051631b2ce7f360e11b81526001600160a01b038281166004830152831690633659cfe69060240161025b565b61033e6103da565b6001600160a01b0381166103a85760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b6103b181610434565b50565b6000806000836001600160a01b03166040516101aa906303e1469160e61b815260040190565b6000546001600160a01b031633146102275760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161039f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146103b157600080fd5b6000602082840312156104ab57600080fd5b81356104b681610484565b9392505050565b600080604083850312156104d057600080fd5b82356104db81610484565b915060208301356104eb81610484565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60008060006060848603121561052157600080fd5b833561052c81610484565b9250602084013561053c81610484565b9150604084013567ffffffffffffffff8082111561055957600080fd5b818601915086601f83011261056d57600080fd5b81358181111561057f5761057f6104f6565b604051601f8201601f19908116603f011681019083821181831017156105a7576105a76104f6565b816040528281528960208487010111156105c057600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b6000602082840312156105f457600080fd5b81516104b681610484565b60018060a01b038316815260006020604081840152835180604085015260005b8181101561063b5785810183015185820160600152820161061f565b506000606082860101526060601f19601f83011685010192505050939250505056fea2646970667358221220eccd60d800c977cbe79b39c73c7c436dc6b87ecf4f38b2c54ea78cccdd04b9ae64736f6c63430008150033'
        );