-- down migration
ALTER TABLE "network_standards"
    DROP CONSTRAINT IF EXISTS network_standards_code_check;

ALTER TABLE "network_standards"
    ADD CONSTRAINT network_standards_name_check CHECK (name ~* '^[a-zA-Z0-9_]+$'),
    ADD CONSTRAINT network_standards_code_check CHECK (code ~* '^[a-zA-Z0-9_]+$');

ALTER TABLE "wallets" 
    ALTER COLUMN "balance" DROP NOT NULL;