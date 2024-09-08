-- Создание таблицы "users"
CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "name" VARCHAR NOT NULL,
                         "lastname" VARCHAR NOT NULL,
                         "email" VARCHAR NOT NULL UNIQUE,
    "country" VARCHAR NOT NULL,
    "hashpass" VARCHAR NOT NULL,
    "isactive" BOOLEAN NOT NULL DEFAULT FALSE,
    "issuperuser" BOOLEAN NOT NULL DEFAULT FALSE,
    "isverified" BOOLEAN NOT NULL DEFAULT FALSE,
    "created_at" timestamptz not null default(now()),
    "updated_at" timestamptz not null default(now())
);

-- Создание таблицы "networkStandards"
CREATE TABLE "network_standards" (
                                    "id" bigserial PRIMARY KEY,
                                    "name" VARCHAR NOT NULL UNIQUE CHECK (name ~* '^[a-zA-Z0-9_]+$'),
    "code" VARCHAR NOT NULL UNIQUE CHECK (code ~* '^[a-zA-Z0-9_]+$')
);

-- Создание таблицы "projects"
CREATE TABLE "projects" (
                            "id" bigserial PRIMARY KEY,
                            "title" VARCHAR NOT NULL UNIQUE CHECK (title ~* '^[a-zA-Z0-9_]+$'),
    "description" VARCHAR NOT NULL,
    "token_title" VARCHAR NOT NULL,
    "image" VARCHAR NOT NULL UNIQUE,
    "amount" float NOT NULL,
    "cost_per_token" float NOT NULL,
    "created_at" timestamptz not null default(now())
);

-- Создание таблицы "transactions"
CREATE TABLE "transactions" (
                                "id" bigserial PRIMARY KEY,
                                "user_id" bigint NOT NULL,
                                "entry_id" bigint NOT NULL,
                                "tx_hash" varchar NOT NULL,
                                "network_id" bigint NOT NULL,
                                "amount" numeric NOT NULL,
                                "created_at" timestamptz not null default(now())
);

-- Создание таблицы "wallets"
CREATE TABLE "wallets" (
                           "id" bigserial PRIMARY KEY,
                           "address" VARCHAR NOT NULL CHECK (address ~* '^[a-zA-Z0-9_]+$'),
    "project_id" bigint,
    "user_id" bigint,
    "network_standard_id" bigint NOT NULL,
    "balance" numeric NOT NULL,
    "created_at" timestamptz not null default(now())
);

-- Создание таблицы "entries"
create table "entries" (
    "id" bigserial primary key,
    "entry_type" varchar not null,
    "project_id" bigint not null,
    "to_wallet_id" bigint not null,
    "from_wallet_id" bigint not null,
    "amount" numeric not null,
    "currency" varchar not null,
    "status" varchar not null,
    "created_at" timestamptz not null default(now())
);

-- Создание индексов для таблиц

-- Индексы для таблицы "users"
CREATE INDEX idx_users_email ON "users"(email);

-- Индексы для таблицы "networkStandards"
CREATE INDEX idx_networkStandards_name ON "network_standards"(name);

-- Индексы для таблицы "projects"
CREATE INDEX idx_cost_per_token ON "projects"(cost_per_token);

-- Индексы для таблицы "transactions"
CREATE INDEX idx_transactions_tx_hash ON "transactions"(tx_hash);
CREATE INDEX idx_entry_idt ON "transactions"(entry_id);

-- Индексы для таблицы "wallets"
CREATE INDEX idx_wallets_address ON "wallets"(address);
create index idx_user_id on wallets (user_id);-- для кошелька
create index idx_project_id on wallets (project_id);-- для кошелька

-- Индексы для таблицы "entries"
CREATE INDEX idx_entries_project_id ON entries (project_id);
CREATE INDEX idx_entries_to_wallet_id ON entries (to_wallet_id);
create index idx_entries_from_wallet_id on entries (from_wallet_id);-- для кошелька

-- Добавление внешних ключей

-- Внешние ключи для таблицы "transactions"
ALTER TABLE "transactions"
    ADD CONSTRAINT fk_user_id FOREIGN KEY ("user_id") REFERENCES "users"("id") ON UPDATE CASCADE ON DELETE CASCADE;
    --ADD CONSTRAINT fk_entry_id FOREIGN KEY ("ticket_project") REFERENCES "projects"("id") ON UPDATE CASCADE ON DELETE CASCADE;

-- Внешние ключи для таблицы "wallets"
ALTER TABLE "wallets"
    ADD CONSTRAINT fk_network_standard_id FOREIGN KEY ("network_standard_id") REFERENCES "network_standards"("id"),
    ADD CONSTRAINT fk_project_id FOREIGN KEY ("project_id") REFERENCES "projects"("id") ON UPDATE CASCADE ON DELETE CASCADE,
    ADD CONSTRAINT fk_user_id FOREIGN KEY ("user_id") REFERENCES "users"("id") ON UPDATE CASCADE ON DELETE CASCADE;

-- Внешние ключи для таблицы "entries"
ALTER TABLE "entries"
    ADD CONSTRAINT fk_from_wallet_id FOREIGN KEY ("from_wallet_id") REFERENCES "wallets"("id") ON UPDATE CASCADE ON DELETE CASCADE,
    ADD CONSTRAINT fk_to_wallet_id FOREIGN KEY ("to_wallet_id") REFERENCES "wallets"("id") ON UPDATE CASCADE ON DELETE CASCADE,
    ADD CONSTRAINT fk_project_id FOREIGN KEY ("project_id") REFERENCES "projects"("id") ON UPDATE CASCADE ON DELETE CASCADE;
