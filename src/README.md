# Edit this Readme and alter the folder as you wish.

# Happy Coding!

### Endpoints:

- host:port/api/v1/swagger/index.html -> swagger docs
- host:3000/                          -> Grafana UI

### Development:

- download the folder
- put a .env file inside with the following fields:<br />
  SQL_ENGINE=postgresql<br />
  SQL_DATABASE=postgres<br />
  SQL_USER=user<br />
  SQL_PASSWORD=password<br />
  SQL_HOST=db<br />
  SQL_PORT=5432<br />
  SERVER_PORT=4599<br />
  MONITORING_PORT=4788<br />
  GIN_MODE=release|debug<br />
  GRAFANA_USER=admin<br />
  GRAFANA_PASS=devops123<br />
- in a terminal, run: `docker compose up --build`

### Deployment:

- create file k8s-secrets.yml and add postgres secret
- $ `kubectl apply k8s-secrets.yml`
- $ `kubectl apply k8s.yml`

#### Built with the [GinBarTonic REST Framework](github.com/martinvuyk/ginbartonic)
