-- 000002_create_user_security_memory_table.up.sql

CREATE TABLE IF NOT EXISTS public.user_security_memory (
    id             uuid PRIMARY KEY,
    session_id     uuid,
    payload        jsonb,
    created_at     timestamptz,
    updated_at     timestamptz,
    quality_status varchar(50),
    ad_question_id varchar(10),
    odm_id         varchar(10),
    user_id        varchar(255),
    team_id        varchar(255),
    source_agent   varchar(50),
    record_type    varchar(50)
);