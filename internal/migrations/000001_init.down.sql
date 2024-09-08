-- Удаление внешних ключей из таблицы "entries"
ALTER TABLE "entries"
DROP CONSTRAINT IF EXISTS fk_from_wallet_id,
    DROP CONSTRAINT IF EXISTS fk_to_wallet_id;

-- Удаление внешних ключей из таблицы "wallets"
ALTER TABLE "wallets"
DROP CONSTRAINT IF EXISTS fk_network_standard_id,
    DROP CONSTRAINT IF EXISTS fk_project_id,
    DROP CONSTRAINT IF EXISTS fk_user_id;

-- Удаление внешних ключей из таблицы "transactions"
ALTER TABLE "transactions"
DROP CONSTRAINT IF EXISTS fk_user_id,
    DROP CONSTRAINT IF EXISTS fk_project_id;

-- Удаление индексов из таблицы "entries"
DROP INDEX IF EXISTS idx_entries_from_wallet_id;
DROP INDEX IF EXISTS idx_entries_to_wallet_id;

-- Удаление индексов из таблицы "wallets"
DROP INDEX IF EXISTS idx_wallets_address;

-- Удаление индексов из таблицы "transactions"
DROP INDEX IF EXISTS idx_transactions_tx_hash;
DROP INDEX IF EXISTS idx_tickets_project;

-- Удаление индексов из таблицы "projects"
DROP INDEX IF EXISTS idx_cost_per_token;

-- Удаление индексов из таблицы "networkStandards"
DROP INDEX IF EXISTS idx_networkStandards_name;

-- Удаление индексов из таблицы "users"
DROP INDEX IF EXISTS idx_users_email;

-- Удаление таблицы "entries"
DROP TABLE IF EXISTS "entries" CASCADE;

-- Удаление таблицы "wallets"
DROP TABLE IF EXISTS "wallets" CASCADE;

-- Удаление таблицы "transactions"
DROP TABLE IF EXISTS "transactions" CASCADE;

-- Удаление таблицы "projects"
DROP TABLE IF EXISTS "projects" CASCADE;

-- Удаление таблицы "networkStandards"
DROP TABLE IF EXISTS "networkStandards" CASCADE;

-- Удаление таблицы "users"
DROP TABLE IF EXISTS "users" CASCADE;
