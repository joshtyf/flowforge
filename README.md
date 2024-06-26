<p align="center">
  <img src="./docs/assets/flowforge.png" alt="Flowforge's logo" width="100" height="120"/>
</p>
<h1 align="center">Flowforge</h1>

Flowforge is a service request management system that allows teams with little developer resources to easily create service pipelines and automate service request fulfilment.

## Quickstart

Using Docker:

Define the following environment variables in a `.env` file in the root directory:

- `POSTGRES_PASSWORD`: The password for the postgres user
- `MONGO_ROOT_PASSWORD`: The password for the root user in the mongo database
- `MONGO_URI`: The URI for the mongo database
- `POSTGRES_URI`: The URI for the postgres database
- `ENV`: The environment the application is running in. Set to `dev` for development
- `AUTH0_DOMAIN`: The domain for the Auth0 application
- `AUTH0_AUDIENCE`: The audience for the Auth0 application
- `MANAGEMENT_API_SECRET`: The secret for the Auth0 management API
- `MANAGEMENT_API_CLIENT`: The client id for the Auth0 management API
- `MANAGEMENT_API_AUDIENCE`: The audience for the Auth0 management API

Then run the following command to start both the frontend and backend:

```bash
docker compose --profile main -p flowforge up --build
```

## Why Flowforge?

A traditional service request system will require a developer to customise a service request form and also create an automation script that will fulfil the request. In a team with few developers, and many different kinds of service requests to deal with, this can be a bottleneck.

Flowforge aims to solve this problem by providing a simple way for developers to define small and reusable steps that can accept dynamic parameters. Non-developers can then craft service pipelines using these steps and map the required step parameters to static values defined by the pipeline creator. These static values can be further made dynamic by using placeholders in the form of `${placeholder_name}` which will be replaced by the final requester when submitting the service request via a form. Apart from the steps which are defined in code, the pipelines and forms are only defined using JSON schema which are curated by non-developers.

This way, developers can focus on creating the steps and maintaining the system, while non-developers can easily create service pipelines and automate service request fulfilment.

## How does Flowforge work?

Service pipelines are defined in JSON schema and consists of a series of execution steps that are already pre-defined in code. Out of the box, Flowforge provides two step types: `WAIT_FOR_APPROVAL` and `API`. The `WAIT_FOR_APPROVAL` step will pause the service request until an admin approves it. The `API` step will accept a URL, headers, query parameters, request body, and method to make an API request. These parameters are statically configured when defining the pipeline. They can also be made dynamic by using placeholders in the form of `${placeholder_name}`.

**Example**

Suppose a pipeline creator wants to create a service pipeline that will fetch data from an API based on the requester's id, the creator can use the `API` step and map the `url` parameter to the following value: `https://myorgdomain.com/api/data/${requester_id}`. The `requester_id` parameter is a dynamic value that will be provided by the final requester when submitting the service request.

This is the JSON schema for the pipeline:

```json
{
  "version": 1,
  "first_step_name": "Make API Call",
  "steps": [
    {
      "step_name": "Make API Call",
      "step_type": "API",
      "next_step_name": "",
      "prev_step_name": "",
      "parameters": {
        "method": "GET",
        "url": "https://myorgdomain.com/api/data/${requester_id}",
        "data": {},
        "headers": {}
      },
      "is_terminal_step": true
    }
  ]
}
```

The pipeline creator will also define a form, in JSON schema, that will collect the requester's id. This form will be used to create a service request. When the service request is submitted, the requester will provide their id and the pipeline will be executed with the `API` step's `url` parameter set to `https://myorgdomain.com/api/data/${requester_id}`, with the placeholder `${requester_id}` being substituted with the submitted requester's id.

This is the form schema:

```json
{
  "fields": [
    {
      "name": "requester_id",
      "title": "RequesterId",
      "description": "Id to be used as the query parameter of the API call",
      "type": "input",
      "required": true,
      "min_length": 1,
      "placeholder": "Enter text..."
    }
  ]
}
```

**Video Demos**

[`./docs/assets/Login to Create Service to Create SR to Approve SR.mp4`](./docs/assets/Login%20to%20Create%20Service%20to%20Create%20SR%20to%20Approve%20SR.mp4)

https://github.com/joshtyf/flowforge/assets/51166055/832586d3-c148-438c-8a09-d2b3318290e3

[`./docs/assets/Reject SR.mp4`](./docs/assets/Reject%20SR.mp4)

https://github.com/joshtyf/flowforge/assets/51166055/45fd31b9-76b2-4fa3-a97f-b5fcabb08156

## Features

- **Reusable Steps**: Service pipeline steps are pre-defined in code and can accept dynamic parameter values
- **Dynamic pipelines**: Non-developers can easily create service pipelines with a variety of pre-defined steps and map the required parameters to static/dynamic values
- **JSON schema**: Pipelines and forms are defined using JSON schema.
- **Approval Workflow**: Pause service requests until an admin approves them
- **Service Request History**: View the history of all service requests and their statuses
- **Detailed Logging**: View detailed logs of each step in the service request
- **Organisation based access control**: Grant access to service requests based on the user's organisation membership

## Organization and Membership

As mentioned under Features, Flowforge provides organizational based access control for pipelines and service requests. As such, users will need to be in an organization to access respective features in Flowforge. Users can either create an organization in `/organization` page, or be added into an existing organization.

Memberships in the organization comes with these 3 roles in descending order of authority: **Owner**, **Admin** and **Member**. Each role comes with access to the following features:

### Owner

- Change organization name
- Transfer ownership of organization

### Admin

- Add member to organization as `Member` or `Admin`
- Remove member from organization
- Promote member from `Member` to `Admin`
- Demote member from `Admin` to `Member`
- Access to _Create Service_ feature
- Access to _Admin Service Request Dashboard_ page

### Member

- View members in organization and their respective roles
- Leave organization
- Access to _Service Catalog_ page
- Access to _Your Service Request Dashboard_ page

> Do note that roles with higher authority levels will also have access to those at lower levels.

## Development

For the development guide, please refer to the respective frontend and backend development guides in [docs/](./docs/).
