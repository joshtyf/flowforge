# FlowForge Backend Documentation

## Database Documentation

### Service request pipeline

Definitions of each service request pipeline will be stored as a JSON object in a MongoDB database.
The JSON object will contain the following key information [TO BE FURTHER UPDATED]:

- Service request pipeline ID
- Service request pipeline name
- Service request pipeline description
- Service request pipeline steps
  - Next step name
  - Previous step name
  - Type
  - Parameters
  - [other fields required by the step type]
  - Start (boolean to indicate the start of the pipeline; only one step can be defined as the start)
  - End (boolean to indicate the end of the pipeline)
- Created by
- Version (use versioning to keep track of changes made to the pipeline)

> Can refer to [AWS States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) to see how it's being implemented by AWS

An example:

```json
{
  "id": "1",
  "name": "Create EC2 instance",
  "description": "Create an EC2 instance",
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
  }
}
```

We use placeholders, denoted by `${}`, to indicate parameters that will be provided by the user during the submission of the service request. The placeholders will be replaced by the actual values when the service request is being processed.

### Service request

Every service request will also be stored as a JSON object in a MongoDB database. The JSON object will contain the following key information [TO BE FURTHER UPDATED]:

- Service request ID
- Service request pipeline ID
- Service request pipeline version
- Status (not started, pending, success, error, cancelled)
- Created timestamp
- Last updated timestamp
- Remarks (used to store any additional information like error messages)
- Form data (the form data submitted by the user)
  - User ID
  - User name
  - [other fields required by the service request pipeline]

### Service request log

In addition, every action made on a server request will be appended to a SQL database. This will function like a log of all the actions made. The table will contain the following information:

- Service request ID
- Step name
- Action (e.g., start, wait for approval, approve, reject, complete, raise error)
- Timestamp

> To think about: compaction of the records in the SQL database

> To think about: passing intermediate data between steps
