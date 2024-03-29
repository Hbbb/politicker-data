-- Workaround for CREATE IF NOT EXISTS
-- https://stackoverflow.com/questions/18389124/simulate-create-database-if-not-exists-for-postgresql
SELECT 'CREATE DATABASE politicker'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'politicker')\gexec

CREATE TABLE IF NOT EXISTS senate_members (
    id SERIAL,
    title TEXT,
    short_title VARCHAR(255),
    api_uri TEXT,
    first_name VARCHAR(255),
    middle_name VARCHAR(255),
    last_name VARCHAR(255),
    suffix VARCHAR(255),
    date_of_birth DATE,
    gender VARCHAR(255),
    party VARCHAR(255),
    leadership_role TEXT,
    twitter_account TEXT,
    facebook_account TEXT,
    youtube_account TEXT,
    govtrack_id VARCHAR(255),
    cspan_id VARCHAR(255),
    votesmart_id VARCHAR(255),
    icpsr_id VARCHAR(255),
    crp_id VARCHAR(255),
    google_entity_id VARCHAR(255),
    fec_candidate_id VARCHAR(255),
    url TEXT,
    rss_url TEXT,
    contact_form TEXT,
    in_office BOOLEAN,
    cook_pvi VARCHAR(255),
    dw_nominate FLOAT4,
    ideal_point VARCHAR(255),
    seniority VARCHAR(255),
    next_election VARCHAR(255),
    total_votes INTEGER,
    missed_votes INTEGER,
    total_present INTEGER,
    last_updated TIMESTAMP,
    ocd_id VARCHAR(255),
    office TEXT,
    phone TEXT,
    fax TEXT,
    state VARCHAR(255),
    senate_class VARCHAR(255),
    state_rank VARCHAR(255),
    lis_id VARCHAR(255),
    missed_votes_pct FLOAT4,
    votes_with_party_pct FLOAT4,
    votes_against_party_pct FLOAT4,

    PRIMARY KEY (id)
);


