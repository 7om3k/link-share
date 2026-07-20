# Link Share
Keep your links in one place and share them with your friends.

## Architecture

- Microservices architecture on back-end.
  - Split by domain (DDD).
  - Every service has its own storage (with different access rights and roles in future).
  - All services are not available from outside of docker network.
- Single access point by web-ui service
  - Progressive enhancement on front-end with web components.
  - Declarative code inside components.
  - Declarative CSS.

## Development Setup

### Required env parameters setup

on Mac or Linux:
```
export LINK_SHARE_DB_PASSWORD=mysecretpassword
```

on Windows
```
C:\Users\you\data-access> set LINK_SHARE_DB_PASSWORD=mysecretpassword
```

### Run the app
In order to run the app for development simply run the command bellow.

```
docker compose -f compose-dev.yaml up --build -d
```

Or without building images again:

```
docker compose -f compose-dev.yaml up -d
```

To stop containers without volumes cleanup (initial database scrips will not be run next time you run database containers) run:

```
docker compose -f compose-dev.yaml down
```

To stop containers along with volumes cleanup, for example when you want to run initial database scrips next time the database containers run use:

```
docker compose -f compose-dev.yaml down -v
```

## Production Run

### Required env parameters setup and security
In order to run the app in production use strong password for LINK_SHARE_DB_PASSWORD

### Run the app

```
docker compose -f compose-prod.yaml up --build -d
```

To stop application

```
docker compose -f compose-prod.yaml down
```
