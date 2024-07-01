# Flowforge Backend Architecture

The Flowforge backend is a Go application that uses a Postgres and Mongo database.

The Go application is composed of:

1. API server
2. Step execution manager

The API server will handle all incoming requests from the frontend. It will interact with the databases to create the necessary resources and also trigger the step execution manager to execute the automation of service requests.

The Mongo database is used to store the service requests and pipeline information while the Postgres database is used to store the users, organisations, as well as step lifecycle events.

## Authentication

Flowforge uses JWT tokens to authenticate users. The JWT token is generated by Auth0 and is used to authenticate the API requests.

> Changes will need to be made if you want to use a different authentication provider.

## Seeding

All the relevant code and information about seeding can be found in the `database/seed` directory.

This is the command to run to seed the databases:

```bash
docker compose --profile be-seed -p flowforge up --build
```

### Seeding service requests and pipelines

Seeding the service requests and pipelines is done in the `mongo_seed.go` file.

### Seeding the user information

Seeding user information is done in the `postgres_seed.go` file. The seeded data should match the data in the identity provider. To create users in the identity provider during the seeding process, set the `create_users` environment variable flag to `true` (requires relevant permissions and tokens).

```bash
create_users=true  docker compose --profile be-seed -p flowforge up --build
```

The data for seeding the users, organisations and memberships are to be created as csv files in the `database/seed/data` directory.

> these files are gitignored and should not be committed to the repository.

`user.csv`: contains all the user information

- `user_id`: id of the user (don't include the identity provider prefix).
- `email`: email of the user.
- `name`: name of the user.
- `idp`: identity provider for the user.
- `password`: password of the user (this column can be left empty if you don't need to create the users in the identity provider).

> The data specified here should match the data in the identity provider.

`org.csv`: contains all the organisation information

- `name`: name of the organisation.
- `owner`: email address of the owner of the organisation.

`membership.csv`: contains all the membership information

- `user`: email address of the user.
- `org`: name of the organisation.
- `role`: role of the user in the organisation.

## Bruno

Flowforge uses [Bruno](https://www.usebruno.com) as the API test tool. All the files can be found at `flowforge_api_bruno`.

To simplify the process of running each API, there is a pre-script created that will fetch the JWT token from Auth0 to authenticate the API requests. To get this pre-script to work correctly, simply create a `.env` file in the `flowforge_api_bruno` directory with the following content:

- `AUTH0_USERNAME`: username for the Auth0 application
- `AUTH0_PASSWORD`: password for the Auth0 application
- `AUTH0_CLIENT_SECRET`: client secret for the Auth0 application