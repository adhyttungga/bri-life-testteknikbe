INSERT INTO product
(product_id, product_name, premium, active)
VALUES
('T10', 'Aurora', 1000, 1),
('L13', 'Davestera', 2000, 1);

INSERT INTO product_parameter
(product_id, parameter_name, parameter_value, active)
VALUES
('T10', 'MINUSIAMASUK', 1, 1),
('T10', 'MAXUSIAMASUK', 99, 1),
('l13', 'MINUSIAMASUK', 18, 1),
('l13', 'MAXUSIAMASUK', 50, 1);

INSERT INTO agent
(agent_id, agent_name, password, active)
VALUES
('BFA01', 'Agent Satu', 'P@ssagent1', 1);
