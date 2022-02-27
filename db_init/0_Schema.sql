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
    account_username VARCHAR(128) UNIQUE NOT NULL,
    data JSONB NOT NULL,
    PRIMARY KEY (user_id, module_name)
);

-- MODULES
--- SPOTIFY
CREATE TABLE module_spotify_artist (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    spotify_id VARCHAR(128) UNIQUE DEFAULT NULL, -- @TODO: Unique only on spotify_id
    display_name VARCHAR(128) UNIQUE NOT NULL    -- Remember to check if there are multiple artist going by the same name
);

CREATE TABLE module_spotify_track (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    spotify_id VARCHAR(128) DEFAULT NULL,
    artist_id UUID NOT NULL REFERENCES module_spotify_artist(id),
    display_name VARCHAR(256) UNIQUE,
    UNIQUE (spotify_id, artist_id)
);

CREATE TABLE module_spotify_played_track (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    track_id UUID NOT NULL REFERENCES module_spotify_track(id),
    played_at TIMESTAMP NULL,
    spotify_played_at VARCHAR(30),
    duration_played INTEGER NULL,
    UNIQUE (track_id, spotify_played_at)
);

--- Twitch