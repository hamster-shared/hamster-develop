update t_template_detail
    set repository_url = 'https://github.com/hamster-template/aptos-token-staking.git',
        repository_name = 'aptos-token-staking',
        code_sources = 'https://raw.githubusercontent.com/hamster-template/aptos-token-staking/main/sources/token-staking.move'
where id = 17;

update t_template_detail
set repository_url = 'https://github.com/hamster-template/aptos-token-vesting.git',
    repository_name = 'aptos-token-vesting',
    code_sources = 'https://raw.githubusercontent.com/hamster-template/aptos-token-vesting/main/sources/token-vesting.move'
where id = 18;

update t_template_detail
set repository_url = 'https://github.com/hamster-template/nft-borrowing-lending-aptos.git',
    repository_name = 'nft-borrowing-lending-aptos',
    code_sources = 'https://raw.githubusercontent.com/hamster-template/nft-borrowing-lending-aptos/main/sources/nftborrowlend.move'
where id = 19;

update t_template_detail
set repository_url = 'https://github.com/hamster-template/aptos-raffle.git',
    repository_name = 'aptos-raffle',
    code_sources = 'https://raw.githubusercontent.com/hamster-template/aptos-raffle/main/sources/raffle.move'
where id = 20;
