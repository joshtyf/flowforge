import { Pipeline } from "./pipeline"
import { JsonFormComponents } from "./json-form-components"

type ServiceRequestForm = {
  user_id: string
  user_name: string
  name: string
  description: string
  form?: JsonFormComponents
}

enum ServiceRequestStatus {
  NOT_STARTED = "Not Started",
  PENDING = "Pending",
  RUNNING = "Running",
  SUCCESS = "Success",
  FAILURE = "Failure",
  CANCELLED = "Canceled",
}

type ServiceRequest = {
  id: string
  pipeline_id: string
  pipeline_version: string
  status: ServiceRequestStatus
  created_on: string
  last_updated: string
  remarks: string
  form_data: ServiceRequestForm
}

export type { ServiceRequestForm, ServiceRequest }

export { ServiceRequestStatus }
