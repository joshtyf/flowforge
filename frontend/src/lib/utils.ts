import { ServiceRequest, ServiceRequestForm } from "@/types/service-request"
import { JsonFormComponents } from "@/types/json-form-components"
import { RJSFSchema, UiSchema } from "@rjsf/utils"
import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const convertServiceRequestFormToRJSFSchema = (
  jsonFormComponents?: JsonFormComponents
) => {
  const schema: RJSFSchema = {
    type: "object",
  }

  const properties: { [key: string]: object } = {}
  const required: string[] = []

  for (const item in jsonFormComponents) {
    const component = jsonFormComponents[item]
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

export const generateUiSchema = (jsonFormComponents?: JsonFormComponents) => {
  const uiSchema: UiSchema = {}

  for (const item in jsonFormComponents) {
    const itemOptions = jsonFormComponents[item]

    switch (itemOptions.type) {
      case "input":
        break
      case "select":
        uiSchema[item] = {
          "ui:placeholder": itemOptions.placeholder ?? "Select an item...",
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

export function isJson(item: string) {
  let value = typeof item !== "string" ? JSON.stringify(item) : item
  try {
    value = JSON.parse(value)
  } catch (e) {
    return false
  }

  return typeof value === "object" && value !== null
}
