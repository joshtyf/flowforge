import { ServiceRequest } from "@/types/service"
import { ServiceRequestForm } from "@/types/sevice-request-form"
import { RJSFSchema, UiSchema } from "@rjsf/utils"
import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const convertServiceRequestFormToRJSFSchema = (
  serviceRequest: ServiceRequestForm
) => {
  const schema: RJSFSchema = {
    type: "object",
  }

  const properties: { [key: string]: object } = {}
  const required: string[] = []

  for (const item in serviceRequest) {
    const component = serviceRequest[item]
    // To create required array
    if (component.required) {
      required.push(item)
    }

    switch (component.type) {
      case "input":
        properties[item] = {
          type: "string",
          title: component.title,
          description: component.description,
          minLength: component.minLength,
        }
        break
      case "select":
        properties[item] = {
          type: "string",
          title: component.title,
          description: component.description,
          enum: component.options,
        }
        break
      case "checkboxes":
        properties[item] = {
          type: "array",
          title: component.title,
          description: component.description,
          items: {
            enum: component.options,
          },
          uniqueItems: true,
        }
        break
      default:
        break
    }
  }

  schema.required = required
  schema.properties = properties

  return schema
}

export const generateUiSchema = (serviceRequest: ServiceRequest) => {
  const uiSchema: UiSchema = {}
  const form: ServiceRequestForm = serviceRequest.form

  for (const item in form) {
    const itemOptions = form[item]

    switch (itemOptions.type) {
      case "input":
        break
      case "select":
        uiSchema[item] = {
          "ui:placeholder": "Select an item...",
        }
        break
      case "checkboxes":
        uiSchema[item] = {
          "ui:widget": "checkboxes",
        }
        break

      default:
        break
    }
  }

  return uiSchema
}
