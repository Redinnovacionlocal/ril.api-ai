-- 000003_add_tree_sub_agent_variants.up.sql

ALTER TABLE tree_sub_agent
    ADD COLUMN domain_prefix VARCHAR(100) NULL,
    ADD COLUMN variant_key   VARCHAR(100) NULL,
    ADD COLUMN variant_hint  VARCHAR(255) NULL;

CREATE UNIQUE INDEX idx_tree_sub_agent_domain_variant
    ON tree_sub_agent (domain_prefix, variant_key)
    WHERE domain_prefix IS NOT NULL AND variant_key IS NOT NULL;
