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
             54,
             3,
             'calc',
             'This template demonstrates a basic calculator dapp. It uses an orthogonally persistent cell variable to store an arbitrary precision integer that represents the result of the most recent calculation.',
             1,
             '0.0.1',
             1,
             7,
             'https://g.alpha.hamsternet.io/ipfs/QmUiEafbupgSwDWWpG5Q8fXrHLBHhMBfJK1APzdcgWnjFT'
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
          54,
          54,
          'calc',
          1,
          '',
          'This template demonstrates a basic calculator dapp. It uses an orthogonally persistent cell variable to store an arbitrary precision integer that represents the result of the most recent calculation.',
          '',
          '',
          '',
          '',
          'hamster-template',
          'https://github.com/hamster-template/calc.git',
          'calc',
          'main',
          '0.0.1',
          'https://raw.githubusercontent.com/hamster-template/calc/main/src/Main.mo',
          'Overview',
          'This example demonstrates a basic calculator dapp. It uses an orthogonally persistent cell variable to store an arbitrary precision integer that represents the result of the most recent calculation.

The dapp provides an interface that exposes the following methods:

add: accepts input and performs addition.
sub: accepts input and performs subtraction.
mul: accepts input and performs multiplication.
div: accepts input, performs division, and returns an optional type to guard against division by zero.
clearall: clears the cell variable by setting its value to zero.
This is a Motoko example that does not currently have a Rust variant.

'
         );