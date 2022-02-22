# Datalake

The goal of this app is to process all RGPD stuff you can get and have neat dashboards helping you to understand them !

## Setup

> This section is a WIP

Install the binary on your server and create either a `datalake.yml` or `/etc/datalake.yml` file with the following default content:

```yaml
database:
  hostname: localhost
  port: 5432
  username: datalake
  password: datalake
  database: datalake

web:
  url: http://localhost:8534
  listening_addr: 0.0.0.0
  port: 8534

# Modules
spotify:
  enabled: true
  client_id: YOUR_SPOTIFY_CLIENT_ID
  client_secret: YOUR_SPOTIFY_CLIENT_SECRET
```