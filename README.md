# Bicincident

A simple API to keep track of my bike incidents

## Install

### Set up credentials

Copy client_secrets file and then fill it with your secrets
```bash
cp app/client_secrets.json.dist app/client_secrets.json
```

### Usage

Spin up containers:
```bash
docker-compose up -d
```

Make an API call:

```bash
curl 0.0.0.0:8585/healthcheck
```