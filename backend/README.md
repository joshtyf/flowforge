# FlowForge Backend Documentation

## Docker

Use docker profiles to manage the different environments.

For example, to seed the database, run the following command:

```bash
docker compose --profile be-seed -p flowforge up --build
```
You may specify the following arguments in front of the above command to seed your own files:<br>
create_users -> specifies whether to create user in auth0, default is "false"<br>
users -> specify user csv file name, default is "user.csv"<br>
orgs -> specify org csv file name, default is "org.csv"<br>
memberships -> specify membership csv file name, default is "membership.csv"

Example:
```bash
create_users=false users=user.csv orgs=org.csv memberships=membership.csv docker compose --profile be-seed -p flowforge up --build
```
**Note: files must be in postgres seed directory**

To run the backend server (no seeding), run the following command:

```bash
docker compose --profile main -p flowforge up --build
```

> This will also start the frontend

## Database Documentation

### Service request pipeline

Definitions of each service request pipeline will be stored as a JSON object in a MongoDB database.
The JSON object will contain the following key information [TO BE FURTHER UPDATED]:

- Service request pipeline ID
- Created by User ID
- Organization ID that pipeline belongs to
- Service request pipeline name
- Service request pipeline description
- Name of the first step
- Service request pipeline steps
  - Next step name
  - Previous step name
  - Type
  - Parameters
  - [other fields required by the step type]
  - Start (boolean to indicate the start of the pipeline; only one step can be defined as the start)
  - End (boolean to indicate the end of the pipeline)
- Created on
- Version (use versioning to keep track of changes made to the pipeline)

> Can refer to [AWS States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) to see how it's being implemented by AWS

An example:

```json
{
  "id": "1",
  "user_id": "testid",
  "org_id": 1,
  "name": "Create EC2 instance",
  "description": "Create an EC2 instance",
  "first_step_name": "Approval",
  "steps": {
    "Approval": {
      "name": "Approval",
      "type": "approval",
      "next": "Create EC2 instance",
      "start": true
    },
    "Create EC2 instance": {
      "name": "Create EC2 instance",
      "type": "GET",
      "parameters": {
        "url": "https://api.example.com/create_ec2_instance",
        "body": {
          "instance_type": "${instance_type}",
          "instance_name": "${instance_name}"
        }
      },
      "end": true
    }
  },
  "created_on": "2024-06-08T06:30:06.137+0000"
}
```

We use placeholders, denoted by `${}`, to indicate parameters that will be provided by the user during the submission of the service request. The placeholders will be replaced by the actual values when the service request is being processed.

#### Step types

To kickstart the development, we will only support the following step types:

1. Approval
2. API call (GET, POST)

### Service request

Every service request will also be stored as a JSON object in a MongoDB database. The JSON object will contain the following key information [TO BE FURTHER UPDATED]:

- Service request ID
- Created by User ID
- Organization ID service request belongs to
- Service request pipeline ID
- Service request pipeline name
- Service request pipeline version
- Status (not started, running, pending, failed, completed, cancelled)
- Created timestamp
- Last updated timestamp
- Remarks (used to store any additional information like error messages)
- Form data (the form data submitted by the user)
  - User ID
  - User name
  - [other fields required by the service request pipeline]

### Service request log

In addition, every action made on a service request will be appended to a SQL database. This will function like a log of all the actions made. The table will contain the following information:

- Service request event ID
- Service request event Type
- Service request ID
- Step name
- Step Type
- Created By user ID
- Create At timestamp

> To think about: compaction of the records in the SQL database

> To think about: passing intermediate data between steps

### User
Every user will be stored in a SQL database without their authorization credentials. For authorization, auth0 is utilized to generate and validate jwt tokens.

The table will contain the following information:

- User Id (using the user_id generated in auth0)
- Name
- Email address (corrosponding email address used when creating a user account in auth0)
- Identity provider (for now only accounts created via auth0 have been integrated. Logins via SSO has not been integrated)
- Created on
- Deleted (for soft deletion)

### Organization
We have internally implemented the concept of collaborating through organizations instead of using auth0 to make the application more modular.

The table will contain the following information:

- Organization ID
- Organization name
- Organization owner
- Created on
- Deleted (for soft deletion)

### Membership
The implementation of organization comes with memberships. This includes developing middleware to validate user permissions.

The table will contain the following information:

- User ID
- Organization ID
- Role (owner, admin, member)
- Joined on
- Deleted (for soft deletion)

### Connecting to MongoDB

Using Homebrew (macOS):

```bash
brew install mongosh
```

Start the mongo db server in Docker.

Connect to the mongo db server:

```bash
mongosh --host 127.0.0.1 -u root
```

Follow the instructions in the official [documentation](https://www.mongodb.com/docs/mongodb-shell/crud/) on how to perform CRUD operations on the database.

## Deployment
To deploy the Flowforge Backend API, you can use a service like Railway. Follow the instructions in the official [documention](https://docs.railway.app/) on how to deploy with Dockerfiles.

### Environment Variables
To deploy Flowforge, the following environment variables will be required:

- POSTGRES_PASSWORD
- POSTGRES_URI
- MONGO_ROOT_PASSWORD
- MONGO_URI
- AUTH0_DOMAIN
- AUTH0_AUDIENCE
- MANAGEMENT_API_SECRET
- MANAGEMENT_API_CLIENT
- MANAGEMENT_API_AUDIENCE

The last five environment variables are required only if auth0 is for authentication and seeding. However, if a different authentication service provider is used, to remember to change the code accordingly to fit the new service provider.

## Authentication
For any queries regarding authentication via auth0, please refer to the official [documention](https://auth0.com/docs).