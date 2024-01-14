# SvEcho
Docker compose development environment for the following stack:

Echo/Go backend + Svelte frontend + Caddy to serve static assets and proxy and Postgres as database.

> This setup is really just meant to imitate local development environment with hot reloading while also allowing you to test production-like served assets / routes / proxying with Caddy. Though this setup would need to be tweaked to actually serve properly for production. Its mostly for my own use-case and interest of learning and automating my local dev environments but feel free to use or take what you want from it.

## Service Container Images
- golang:1.21-alpine (`backend`)
- caddy:2.7.6-alpine (`caddy`)
- postgres:16-alpine (`db`)
- node:20-alpine (`frontend`)

## Main Dependency Versions
- Go 1.21
- Echo v4
- Svelte 3.x
- Vite 4.x
- Postgres 16

## Instructions to run
Docker + Docker Compose is required. The containers take ~10s to ~20s to come up all the way. Source code folders `backend`/`frontend` are mounted into the containers so we can support hot reloads. If you save an erroring line the container can possibly exit so just `docker compose up -d` and itll restart the stopped ones once you fix your code.

From the root project directory, to get the containers built and up and running:

```
docker compose up
```

or to bring them up in background to get your terminal back
```
docker compose up -d
```

To view/follow the logs of all service containers
```
docker compose logs -f
```

### Notes
- Caddy is reachable on port 80, will serve Svelte static files out on `http://localhost`. This will serve production like assets which are only built on launch of the container or if you manually run the `npm run build` command against the running container.

- Go/Echo server is meant to be serving a REST API that Svelte would tap on for data. So all your routes would ideally be prefixed with `/api` since caddy proxy passes to Echo using that suffix. So Svelte would consume the API via `fetch('http://localhost/api/....')` or Axios requests, etc.

- Go/Echo dev server is directly reachable on port 8000 - `http://localhost:8000/api` - hot reload is enabled.
- Svelte (vite dev server) is directly reachable on port 3000 - `http://localhost:3000/` - hot reload is enabled.
