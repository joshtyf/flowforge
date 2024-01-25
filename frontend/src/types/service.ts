import { RJSFSchema } from "@rjsf/utils"

type ServiceRequest = {
  id?: number
  name: string
  description: string
  form: RJSFSchema
  // More to be added
}

export type { ServiceRequest }
