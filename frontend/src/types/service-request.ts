import { Pipeline } from "./pipeline"
import { JsonFormComponents } from "./json-form-components"

type ServiceRequestForm = {
  id?: string
  name: string
  description: string
  form: JsonFormComponents
}

type ServiceRequest = {
  id: string
  pipeline_id: string
  pipeline_version: string
  status: string
  created_on: string
  last_updated: string
  remarks: string
}

export type { ServiceRequestForm, ServiceRequest }
