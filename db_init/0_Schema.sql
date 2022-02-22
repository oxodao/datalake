CREATE TABLE datalake_user (
    id UUID PRIMARY KEY default gen_random_uuid(),
    name VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR NOT NULL
);

CREATE TABLE datalake_refreshtoken (
    id UUID PRIMARY KEY default gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES datalake_user(id),
    token VARCHAR(128) NOT NULL,
    expires_at TIMESTAMP NOT NULL DEFAULT NOW() + INTERVAL '6 months'
);

CREATE TABLE provider_authentication (
    user_id UUID NOT NULL REFERENCES datalake_user(id),
    module_name VARCHAR(50) UNIQUE NOT NULL,
    data JSONB NOT NULL,
    PRIMARY KEY (user_id, module_name)
);