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
docker compose --profile be-seed -p flowforge up --build
```
You may specify the following arguments in front of the above command to seed your own files:
create_users -> specifies whether to create user in auth0, default is "false"
users -> specify user csv file name, default is "user.csv"
orgs -> specify org csv file name, default is "org.csv"
memberships -> specify membership csv file name, default is "membership.csv"

Example:
```bash
create_users=false users=user.csv orgs=org.csv memberships=membership.csv docker compose --profile be-seed -p flowforge up --build
```
**Note: files must be in postgres seed directory**

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
docker compose -p flowforge down -v --rmi local 
```

## Setup Frontend Development

Read `frontend/README.md` for further instructions on how to setup the frontend development environment

## Setup Backend Development

Read `backend/README.md` for further instructions on how to setup the frontend development environment
