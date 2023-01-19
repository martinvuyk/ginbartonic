# Edit this Readme and alter the folder as you wish.

# Happy Coding!

### Endpoints:

- base_url/hello/:name -> greeting GET Endpoint
- base_url/docs -> swagger docs (not yet implemented)

### Deployment:

- download the folder
- put a .env file inside with the following fields:
SQL_ENGINE=<br />
SQL_DATABASE=<br />
SQL_USER=<br />
SQL_PASSWORD=<br />
SQL_HOST=db<br />
SQL_PORT=<br />
SERVER_PORT=4599<br />
GIN_MODE=release|debug<br />

- in a terminal, write: "docker compose up"

#### Requirements:

-docker

#### Built with the GinBarTonic REST Framework (github.com/martinvuyk/ginbartonic)
