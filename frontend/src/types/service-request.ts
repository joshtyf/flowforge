import { StepStatus } from "./pipeline"
import { JsonFormComponents } from "./json-form-components"

type ServiceRequestForm = object

enum ServiceRequestStatus {
  NOT_STARTED = "Not Started",
  RUNNING = "Running",
  PENDING = "Pending",
  FAILED = "Failed",
  CANCELLED = "Cancelled",
  COMPLETED = "Completed",
}

// type ServiceRequestStep = {
//   name: string
//   type?: string
//   next?: string
//   start?: boolean
//   end?: boolean
//   status: ServiceRequestStatus
// }

type ServiceRequestStep = {
  name: string
  status: StepStatus
  updated_at?: string
  updated_by?: string
  next_step_name: string
}

type ServiceRequestSteps = {
  [key: string]: ServiceRequestStep
}

type ServiceRequest = {
  id: string
  user_id: string
  pipeline_id: string
  pipeline_name: string
  pipeline_version: string
  status: ServiceRequestStatus
  created_on: string
  // TODO: Make field mandatory once accounts are tag to service request
  created_by?: string
  last_updated: string
  remarks: string
  form_data: ServiceRequestForm
  first_step_name: string
  steps?: ServiceRequestSteps
  // TODO: To refactor in future when service request details is implemented
  pipeline?: { form: JsonFormComponents }
}

type ServiceRequestLogs = {
  step_name: string
  logs: string[]
  end_of_logs: boolean
  next_offset: number
}

export type {
  ServiceRequestForm,
  ServiceRequestStep,
  ServiceRequestSteps,
  ServiceRequest,
  ServiceRequestLogs,
}

export { ServiceRequestStatus }
