import { Pipeline } from "./pipeline"
import { ServiceRequestForm } from "./sevice-request-form"

type ServiceRequest = {
  id?: number
  name: string
  description: string
  form: ServiceRequestForm
  pipeline?: Pipeline
}

type ServiceRequestWithSteps = ServiceRequest & {
  steps: object
}

export type { ServiceRequest, ServiceRequestWithSteps }
