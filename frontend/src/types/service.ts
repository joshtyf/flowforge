import { ServiceRequestForm } from "./sevice-request-form"

type ServiceRequest = {
  id?: number
  name: string
  description: string
  form: ServiceRequestForm
  // More to be added
}

type ServiceRequestWithSteps = ServiceRequest & {
  steps: object
}

export type { ServiceRequest, ServiceRequestWithSteps }
