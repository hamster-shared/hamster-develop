-- 留存删掉的数据 start
-- INSERT INTO aline.t_template (id, template_type_id, name, description, audited, last_version, logo, create_time, update_time, delete_time, whether_display, image, language_type, deploy_type) VALUES (18, 1, 'Multiwrap', 'Bundle multiple ERC721/ERC1155/ERC20 tokens into a single ERC721.', 1, '1.0.0', null, '2023-02-13 16:26:07', '2023-02-13 16:26:07', null, 0, null, 1, null);
-- INSERT INTO aline.t_template (id, template_type_id, name, description, audited, last_version, logo, create_time, update_time, delete_time, whether_display, image, language_type, deploy_type) VALUES (19, 1, 'Pack', 'Pack multiple tokens into ERC1155 NFTs that act as randomized loot boxes', 1, '1.0.0', null, '2023-02-13 16:26:07', '2023-02-13 16:26:07', null, 0, null, 1, null);
-- INSERT INTO aline.t_template (id, template_type_id, name, description, audited, last_version, logo, create_time, update_time, delete_time, whether_display, image, language_type, deploy_type) VALUES (20, 2, 'rTokens', 'rTokens are the tokens exchangeable for their underlying ones (e.g. rDai -> Dai).', 1, '1.0.0', null, '2023-02-13 16:26:07', '2023-02-13 16:26:07', null, 0, null, 1, null);
-- INSERT INTO aline.t_template (id, template_type_id, name, description, audited, last_version, logo, create_time, update_time, delete_time, whether_display, image, language_type, deploy_type) VALUES (21, 2, 'Quadratic Funding', 'Funding round that uses quadratic matching (capital-constrained liberal radicalism)', 1, '1.0.0', null, '2023-02-13 16:26:07', '2023-02-13 16:26:07', null, 0, null, 1, null);
-- INSERT INTO aline.t_template (id, template_type_id, name, description, audited, last_version, logo, create_time, update_time, delete_time, whether_display, image, language_type, deploy_type) VALUES (23, 3, 'Defi', 'Create Contract for Defi.', 1, '1.0.0', null, '2023-02-13 16:26:07', '2023-02-13 16:26:07', null, 0, null, 1, null);
-- INSERT INTO aline.t_template (id, template_type_id, name, description, audited, last_version, logo, create_time, update_time, delete_time, whether_display, image, language_type, deploy_type) VALUES (24, 4, 'Vote', 'Create and vote on proposals', 1, '1.0.0', null, '2023-02-13 16:26:07', '2023-02-13 16:26:07', null, 0, null, 1, null);
-- INSERT INTO aline.t_template (id, template_type_id, name, description, audited, last_version, logo, create_time, update_time, delete_time, whether_display, image, language_type, deploy_type) VALUES (25, 4, 'Split', 'Distribute funds among multiple recipients', 1, '1.0.0', null, '2023-02-13 16:26:07', '2023-02-13 16:26:07', null, 0, null, 1, null);
--
-- INSERT INTO aline.t_template_detail (id, template_id, name, audited, extensions, description, examples, resources, abi_info, byte_code, author, repository_url, repository_name, branch, version, code_sources, create_time, update_time, delete_time, title, title_description) VALUES (10, 18, 'Multiwrap', 1, '', 'Bundle multiple ERC721/ERC1155/ERC20 tokens into a single ERC721.', '', '', '', '', 'hamster-template', '', '', 'main', '1.0.0', '', '2023-02-13 16:28:27', '2023-02-13 16:28:27', null, '', '');
-- INSERT INTO aline.t_template_detail (id, template_id, name, audited, extensions, description, examples, resources, abi_info, byte_code, author, repository_url, repository_name, branch, version, code_sources, create_time, update_time, delete_time, title, title_description) VALUES (11, 19, 'Pack', 1, '', 'Pack multiple tokens into ERC1155 NFTs that act as randomized loot boxes.', '', '', '', '', 'hamster-template', '', '', 'main', '1.0.0', '', '2023-02-13 16:28:27', '2023-02-13 16:28:27', null, '', '');
-- INSERT INTO aline.t_template_detail (id, template_id, name, audited, extensions, description, examples, resources, abi_info, byte_code, author, repository_url, repository_name, branch, version, code_sources, create_time, update_time, delete_time, title, title_description) VALUES (12, 20, 'Pack', 1, '', 'rTokens are the tokens exchangeable for their underlying ones (e.g. rDai -> Dai).', '', '', '', '', 'hamster-template', '', '', 'main', '1.0.0', '', '2023-02-13 16:28:27', '2023-02-13 16:28:27', null, '', '');
-- INSERT INTO aline.t_template_detail (id, template_id, name, audited, extensions, description, examples, resources, abi_info, byte_code, author, repository_url, repository_name, branch, version, code_sources, create_time, update_time, delete_time, title, title_description) VALUES (13, 21, 'Quadratic Funding', 1, '', 'Funding round that uses quadratic matching (capital-constrained liberal radicalism).', '', '', '', '', 'hamster-template', '', '', 'main', '1.0.0', '', '2023-02-13 16:28:27', '2023-02-13 16:28:27', null, '', '');
-- INSERT INTO aline.t_template_detail (id, template_id, name, audited, extensions, description, examples, resources, abi_info, byte_code, author, repository_url, repository_name, branch, version, code_sources, create_time, update_time, delete_time, title, title_description) VALUES (15, 23, 'Defi', 1, '', 'Create Contract for Defi.', '', '', '', '', 'hamster-template', '', '', 'main', '1.0.0', '', '2023-02-13 16:28:27', '2023-02-13 16:28:27', null, '', '');
-- INSERT INTO aline.t_template_detail (id, template_id, name, audited, extensions, description, examples, resources, abi_info, byte_code, author, repository_url, repository_name, branch, version, code_sources, create_time, update_time, delete_time, title, title_description) VALUES (16, 24, 'Vote', 1, '', 'Create and vote on proposals.', '', '', '', '', 'hamster-template', '', '', 'main', '1.0.0', '', '2023-02-13 16:28:27', '2023-02-13 16:28:27', null, '', '');
-- INSERT INTO aline.t_template_detail (id, template_id, name, audited, extensions, description, examples, resources, abi_info, byte_code, author, repository_url, repository_name, branch, version, code_sources, create_time, update_time, delete_time, title, title_description) VALUES (17, 25, 'Split', 1, '', 'Distribute funds among multiple recipients.', '', '', '', '', 'hamster-template', '', '', 'main', '1.0.0', '', '2023-02-13 16:28:27', '2023-02-13 16:28:27', null, '', '');
------ end

DELETE from aline.t_template WHERE id IN (18,19,20,21,23,24,25);
DELETE from aline.t_template_detail WHERE id IN (10,11,12,13,15,16,17);
