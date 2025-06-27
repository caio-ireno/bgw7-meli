-- Charset do banco
SELECT DEFAULT_CHARACTER_SET_NAME
FROM information_schema.SCHEMATA
WHERE SCHEMA_NAME = 'empresa_internet';

-- Charset da tabela
SHOW TABLE STATUS FROM empresa_internet WHERE Name = 'customers';

-- Charset das colunas
SHOW FULL COLUMNS FROM customers;

SET NAMES 'utf8mb4';
select * from customers;

SHOW VARIABLES LIKE 'character_set%';

ALTER TABLE customers
    MODIFY name       VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    MODIFY last_name  VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    MODIFY province   VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    MODIFY city       VARCHAR(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;