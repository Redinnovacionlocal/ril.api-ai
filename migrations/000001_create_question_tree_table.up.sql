-- 000001_create_tree_sub_agent_table.up.sql

CREATE TABLE tree_sub_agent (
    id_tree_sub_agent   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name                VARCHAR(255) NOT NULL,
    active              BOOLEAN NOT NULL DEFAULT false,
    excel_gcs_path      VARCHAR(255) NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);