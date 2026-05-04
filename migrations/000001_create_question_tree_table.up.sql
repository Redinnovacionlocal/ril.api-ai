-- 000001_create_tree_sub_agent_table.up.sql

CREATE TABLE tree_sub_agent (
    id_tree_sub_agent   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name                VARCHAR(255) NOT NULL,
    active              BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE tree_questions (
    id                    UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    id_tree_sub_agent     UUID NOT NULL REFERENCES tree_sub_agent(id_tree_sub_agent),
    number                INT NOT NULL,
    dimension             VARCHAR(255) NOT NULL,
    question              TEXT NOT NULL,
    options               TEXT,
    supporting_document   TEXT,
    minimum_cert_criteria VARCHAR(100) NOT NULL,
    valid_formats         TEXT[],
    good_practices        TEXT[],
    alert_signals         TEXT[],
    agent_help            TEXT[],
    tags                  TEXT[],
    odm_ids               TEXT[],
    created_at            TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT uq_question_number UNIQUE (id_tree_sub_agent, number)
);