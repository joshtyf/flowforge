import { JsonFormComponents } from "./json-form-components"
import { ServiceRequestForm } from "./service-request"

type PipelineStep = {
  step_name: string
  step_type: "API" | "WAIT_FOR_APPROVAL"
  next_step_name: string
  prev_step_name: string
  parameters: {
    method: "GET" | "POST" | "PATCH"
    url: "string"
  }
  is_terminal_step: boolean
}

type PipelineDetails = {
  id?: string
  version?: number
  first_step_name?: string
  steps?: PipelineStep[]
  created_on?: string
}

type Pipeline = PipelineDetails & {
  pipeline_name: string
  pipeline_description?: string
  form?: JsonFormComponents
}

export type { Pipeline, PipelineDetails, PipelineStep }
