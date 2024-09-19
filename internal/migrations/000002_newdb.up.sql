ALTER TABLE "network_standards"
    DROP CONSTRAINT IF EXISTS network_standards_name_check,
    DROP CONSTRAINT IF EXISTS network_standards_code_check;

ALTER TABLE "network_standards"
    ADD CONSTRAINT network_standards_code_check CHECK (code IN ('TRC20', 'TON', 'ERC20', 'BEP20', 'SOL'));

-- Обновляем столбец name для разрешения любых значений (если необходимо)
ALTER TABLE "network_standards"
    ALTER COLUMN "name" TYPE VARCHAR;

ALTER TABLE "projects"
    ALTER COLUMN "amount" TYPE BIGINT USING amount::bigint,
    ADD COLUMN "unlocked_tokens" BIGINT CHECK ("unlocked_tokens" <= "amount") DEFAULT 0;