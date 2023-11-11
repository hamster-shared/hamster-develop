ALTER TABLE `t_backend_deploy`
  MODIFY `deploy_tx_hash` CHAR(200);

ALTER TABLE `t_contract_deploy`
  MODIFY `deploy_tx_hash` CHAR(200);
