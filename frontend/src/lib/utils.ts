import { ServiceRequest } from "@/types/service"
import { UiSchema } from "@rjsf/utils"
import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

// TODO: Refactor util based on own schema
export const generateUiSchema = (serviceRequest: ServiceRequest) => {
  const uiSchema: UiSchema = {}
  const properties = serviceRequest.form?.properties // Optional chaining

  if (properties) {
    for (const property in properties) {
      const prop = properties[property] // Direct access after check
      if (typeof prop === "object" && prop.type === "array") {
        uiSchema[property] = {
          "ui:widget": "checkboxes",
        }
      }
      if (typeof prop === "object" && prop.type === "string" && !!prop.enum) {
        uiSchema[property] = {
          "ui:placeholder": "Select an item...",
        }
      }
    }
  }
  return uiSchema
}
