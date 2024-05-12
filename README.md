# flowforge

Service request pipeline web application.

# Development Guide

## Using Docker

Run the following command to start both the frontend and backend:

```bash
docker compose --profile main -p flowforge up --build
```

Run the following command to seed the database wth sample data

```bash
CREATE_USERS=<boolean> USER_SEED_FILENAME=<filename> ORG_SEED_FILENAME=<filename> MEMBERSHIP_SEED_FILENAME=<filename> docker compose --profile be-seed -p flowforge up --build
```

Run the following command to start just the frontend:

```bash
docker compose --profile fe -p flowforge up --build
```

Run the following command to start just the backend:

```bash
docker compose --profile be -p flowforge up --build
```

Run the following command to delete container

```bash
docker compose -p flowforge down
```

## Setup Frontend Development

Read `frontend/README.md` for further instructions on how to setup the frontend development environment

## Setup Backend Development

Read `backend/README.md` for further instructions on how to setup the frontend development environment
