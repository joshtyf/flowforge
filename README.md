<p align="center">
  <img src="./docs/assets/flowforge.png" alt="Flowforge's logo" width="100" height="120"/>
</p>
<h1 align="center">Flowforge</h1>

Flowforge is a service request management system that allows teams with little developer resources to easily create service pipelines and automate service request fulfilment.

## Quickstart

Using Docker:

Define the following environment variables in a `.env` file in the root directory:

- `POSTGRES_PASSWORD`: The password for the postgres user.
- `MONGO_ROOT_PASSWORD`: The password for the root user in the mongo database.
- `MONGO_URI`: The URI for the mongo database.
- `POSTGRES_URI`: The URI for the postgres database.
- `ENV`: The environment the application is running in. Set to `dev` for development.
- `AUTH0_DOMAIN`: The domain for the Auth0 application.
- `AUTH0_AUDIENCE`: The audience for the Auth0 application.
- `MANAGEMENT_API_SECRET`: The secret for the Auth0 management API.
- `MANAGEMENT_API_CLIENT`: The client id for the Auth0 management API.
- `MANAGEMENT_API_AUDIENCE`: The audience for the Auth0 management API.

Ensure that you have provided environment variables for the frontend application in a `.env.local` file in the frontend directory. Check [docs/FRONTEND_DEV_GUIDE.md](./docs/FRONTEND_DEV_GUIDE.md) for more info.

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

[`docs/assets/1. Login, New User, New Org.mp4`](./docs/assets/1.%20Login,%20New%20User,%20New%20Org.mp4)

https://github.com/joshtyf/flowforge/assets/51166055/1f964ee7-98b3-4614-9277-3c17e56d6f6b

[`/docs/assets/2. Create Service, Create Request, View Service Request Info.mp4`](./docs/assets/2.%20Create%20Service,%20Create%20Request,%20View%20Service%20Request%20Info.mp4)

https://github.com/joshtyf/flowforge/assets/51166055/2f8b2827-d7c9-418b-bc36-b31f5505ce85

[`docs/assets/3. Start Service Request, View Logs.mp4`](./docs/assets/3.%20Start%20Service%20Request,%20View%20Logs.mp4)

https://github.com/joshtyf/flowforge/assets/51166055/7d5fcef6-f635-4d2c-ab0e-9091c9f269d2

[`docs/assets/4. Approve Service Request.mp4`](./docs/assets/4.%20Approve%20Service%20Request.mp4)

https://github.com/joshtyf/flowforge/assets/51166055/368e4d1e-ddeb-4749-847d-ac4b1de01951

[`docs/assets/5. Reject Service Request.mp4`](./docs/assets/5.%20Reject%20Service%20Request.mp4)

https://github.com/joshtyf/flowforge/assets/51166055/82dcb17e-726b-4161-bc69-e70e3da37c12

[`docs/assets/6. Cancel Service Request.mp4`](./docs/assets/6.%20Cancel%20Service%20Request.mp4)

https://github.com/joshtyf/flowforge/assets/51166055/1489e83f-5832-4593-b09f-2b981eab38a6

## Features

- **Reusable Steps**: Service pipeline steps are pre-defined in code and can accept dynamic parameter values
- **Dynamic pipelines**: Non-developers can easily create service pipelines with a variety of pre-defined steps and map the required parameters to static/dynamic values
- **JSON schema**: Pipelines and forms are defined using JSON schema.
- **Approval Workflow**: Pause service requests until an admin approves them
- **Service Request History**: View the history of all service requests and their statuses
- **Detailed Logging**: View detailed logs of each step in the service request
- **Organisation based access control**: Grant access to service requests based on the user's organisation membership

## Organization and Membership

Flowforge provides organizational based access control for pipelines and service requests. Users can create or join organizations to access features. There are three roles: Owner, Admin, and Member. Each role has different levels of authority and access to features. Members can access the service catalog and request for services. Admins can additionally manage members, create services, and access the admin dashboard to approve or reject requests. Owners have full control over the organization and can manage admins and members.

## Development

For the development guide, please refer to the respective frontend and backend development guides in [docs/](./docs/).
