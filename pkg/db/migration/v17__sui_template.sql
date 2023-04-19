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
             41,
             1,
             'sui-nfts',
             ' Issuing the first ten natural numbers as collectible NFT''s.',
             1,
             '0.0.1',
             1,
             5,
             ''
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
             41,
             41,
             'sui-nfts',
             1,
             '',
             '',
             '',
             '',
             '',
             '',
             'hamster-template',
             'https://github.com/hamster-template/sui-nft.git',
             'sui-nft',
             'main',
             '0.0.1',
             'https://raw.githubusercontent.com/hamster-template/sui-nft/main/sources/devnet_nft.move',
             '',
             ''
         );
