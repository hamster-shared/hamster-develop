alter table t_contract_deploy
    add column  abi_info mediumtext comment 'abi信息';


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
             40,
             1,
             'to_do_list',
             'Smart contracts that move tasks from one list to another.',
             1,
             '1.0.0',
             0,
             2,
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
            27,
            (
                SELECT id
                FROM t_template
                WHERE name = 'to_do_list'
            ),
            'to_do_list',
            1,
            '',
            '',
            '',
            '',
            '{
  "address": "0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a",
  "name": "todolist",
  "friends": [],
  "exposed_functions": [
    {
      "name": "complete_task",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "u64"
      ],
      "return": []
    },
    {
      "name": "create_list",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer"
      ],
      "return": []
    },
    {
      "name": "create_task",
      "visibility": "public",
      "is_entry": true,
      "generic_type_params": [],
      "params": [
        "&signer",
        "0x1::string::String"
      ],
      "return": []
    }
  ],
  "structs": [
    {
      "name": "Task",
      "is_native": false,
      "abilities": [
        "copy",
        "drop",
        "store"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "task_id",
          "type": "u64"
        },
        {
          "name": "address",
          "type": "address"
        },
        {
          "name": "content",
          "type": "0x1::string::String"
        },
        {
          "name": "completed",
          "type": "bool"
        }
      ]
    },
    {
      "name": "TodoList",
      "is_native": false,
      "abilities": [
        "key"
      ],
      "generic_type_params": [],
      "fields": [
        {
          "name": "tasks",
          "type": "0x1::table::Table<u64, 0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a::todolist::Task>"
        },
        {
          "name": "set_task_event",
          "type": "0x1::event::EventHandle<0x8c6ae8acc5839cfe9bff5c94bdcc36c81dc62634853808ec1b0ce2b71653f54a::todolist::Task>"
        },
        {
          "name": "task_counter",
          "type": "u64"
        }
      ]
    }
  ]
}',
            '',
            'hamster-template',
            'https://github.com/hamster-template/to_do_list.git',
            'to_do_list',
            'main',
            '1.0.0',
            'https://raw.githubusercontent.com/hamster-template/to_do_list/main/sources/todolist.move',
            '',
            ''
         );