import { Pipeline } from "./pipeline"
import { JsonFormComponents } from "./json-form-components"

type ServiceRequestForm = object

enum ServiceRequestStatus {
  NOT_STARTED = "Not Started",
  PENDING = "Pending",
  REJECTED = "Rejected",
  RUNNING = "Running",
  SUCCESS = "Success",
  FAILURE = "Failure",
  CANCELLED = "Canceled",
  COMPLETED = "Completed",
}

type ServiceRequestStep = {
  name: string
  type?: string
  next?: string
  start?: boolean
  end?: boolean
  status: ServiceRequestStatus
}

type ServiceRequest = {
  id: string
  pipeline_id: string
  pipeline_name: string
  pipeline_description?: string
  pipeline_version: string
  status: ServiceRequestStatus
  created_on: string
  // TODO: Make field mandatory once accounts are tag to service request
  created_by?: string
  last_updated: string
  remarks: string
  form: JsonFormComponents
  form_data: ServiceRequestForm
  steps?: ServiceRequestStep[]
  currentStep?: ServiceRequestStep
}

export type { ServiceRequestForm, ServiceRequestStep, ServiceRequest }

export { ServiceRequestStatus }
