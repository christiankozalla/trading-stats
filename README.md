# inloopo Trading Stats

Traders need to track their progress over time to eventually sit down and assess their progress, reflect decision they made, sharpen their mindsets and iterate on their strategy.

This app aims to be a trader's daily companion and single source of truth for providing insights on their trading activity via data visualizations such as charts and tables.


## Development

A Procfile is used for starting up the development environment. Here is a list of tools that need to be installed on the development machine.

- [Hivemind](https://evilmartians.com/opensource/hivemind): Runs processes declared in a Procfile in parallel
- [modd](https://github.com/cortesi/modd): Watches for changes in Golang source code and rebuilds the PocketBase server upon changes.
- [Docker](https://docker.com): Spins up a Caddy server delegating requests from the main host to the Vite dev-server and the PocketBase REST API. With this setup both frontend dev-server and backend API are available on the same host during development, so it's running mostly like in production, without CORS issues etc.

```shell
# Start the Dev-Environment
hivemind
```

## Running in production

The app is currently running on a VPS behind a Caddy server reverse-proxy that routes all requests to PocketBase. PocketBase serves the Vue frontend app from its static directory `pb_public`, and acts as a REST API on all routes like /api/* . PocketBase is deployed as a single binary. A single process is managed by systemd. All data is currently stored in the `pb_data` directory such as the SQLite database or user-provided uploaded files. In the future, an S3 storage can be used for such file hosting.

### Hosting

#### Directory structure

```
    -- ~/projects/trading-stats
        -- production # contains the `its` binary
            -- pb_public # static assets (e.g. Vue app) being served by PocketBase
            -- pb_data  # SQLite database(s) and user file uploads
        - standard.log # logs from stdout of the application
        - errors.log # logs from stderr of the application
```

#### Systemd services

- /lib/systemd/system/trading-stats.service is running the application
- /lib/systemd/system/trading-stats-restart.service restarting the application when a the `its` binary is changed, watched by trading-stats.restart.path unit
    - a restart can be triggered manually with `touch ~/projects/trading-stats/production/its` or through the upload of a fresh executable

#### Logs

Follow the logs with e.g. `tail -f ~/projects/trading-stats/standard.log`