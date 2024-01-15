# FlowForge Backend Documentation

## Database Documentation

Definitions of each service request pipeline will be stored as a JSON object in a MongoDB database.
The JSON object will contain the following key information [TO BE FURTHER UPDATED]:

- Service request pipeline ID
- Service request pipeline name
- Service request pipeline description
- Service request pipeline steps
  - Next step name
  - Previous step name
  - Type
  - Parameters (?) (to be further updated)
- Created by
- Version (use versioning to keep track of changes made to the pipeline)

> Can refer to [AWS States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html) to see how it's being implemented by AWS

Every service request will also be stored as a JSON object in a MongoDB database. The JSON object will contain the following key information [TO BE FURTHER UPDATED]:

- Service request ID
- Service request pipeline ID
- Service request pipeline version
- Completion status
- Last updated timestamp
- Completed by (could be automated or the service admin)
- Remarks (used to store any additional information like error messages)

In addition, every action made on a server request will be appended to a SQL database. This will function like a log of all the actions made. The table will contain the following information:

- Service request ID
- Step name
- Action (e.g., not started, started, pending approval, rejected, completed, error)
- Timestamp

> To think about: compaction of the records in the SQL database

> To think about: how to design the database to pause and wait for approval
