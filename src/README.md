# Edit this Readme and alter the folder as you wish.

# Happy Coding!

### Endpoints:

- host:port/api/v1/swagger/index.html -> swagger docs
- host:3000/                          -> Grafana UI

### Development:

- download the folder
- put a .env file inside with the following fields:<br />
  SQL_ENGINE=postgresql<br />
  SQL_DATABASE=<br />
  SQL_USER=<br />
  SQL_PASSWORD=<br />
  SQL_HOST=db<br />
  SQL_PORT=5432<br />
  SERVER_PORT=4599<br />
  MONITORING_PORT=4788<br />
  GIN_MODE=release|debug<br />
  GRAFANA_USER=admin<br />
  GRAFANA_PASS=devops123<br />
- in a terminal, write: "docker compose up --build"

#### Requirements:

-docker

#### Built with the [GinBarTonic REST Framework](github.com/martinvuyk/ginbartonic)
