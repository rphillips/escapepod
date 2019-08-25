# Development

For local development the following dependencies are required:

* Nodejs 11+
* Golang 1.12+
* Docker and Docker-Compose (Postgresql or Sqlite3 database)

Escapepod has the following components a frontend and backend component. The internal sqlite3 database can be used or postgresql.

## Frontend

`./frontend-vue` contains the frontend client side code displayed in a webbrowser.

## Backend

`./` the main project directory is a standard golang project layout for the
backend service. This backend service contains the API and RSS/ATOM crawler.

## Bundled builds

Packaging bundles the built frontend and embeds it on the end of the binary
for easy distribution.

## Iterative Development

[tmuxp](https://github.com/tmux-python/tmuxp) is a nice session manager for
tmux. The following command will run a tmux session, start the database (with
docker-compose), start frontend dev server (localhost:8080), and start the
backend service:

```sh
tmuxp load .tmuxp.yaml
```

The backend service is run with [modd](https://github.com/cortesi/modd). modd
is a nice utility to watch for source changes. When it detects a golang
source change it will automatically rebuild the backend and restart the
service.
