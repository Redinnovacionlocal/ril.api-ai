-- 000003_add_tree_sub_agent_variants.down.sql

DROP INDEX IF EXISTS idx_tree_sub_agent_domain_variant;

ALTER TABLE tree_sub_agent
    DROP COLUMN IF EXISTS domain_prefix,
    DROP COLUMN IF EXISTS variant_key,
    DROP COLUMN IF EXISTS variant_hint;
