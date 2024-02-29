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
        uiSchema[item] = {
          "ui:placeholder": itemOptions.placeholder,
        }
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

export function formatDateString(date: Date) {
  const options: Intl.DateTimeFormatOptions = {
    day: "numeric",
    month: "short",
    year: "numeric",
  }

  return date.toLocaleDateString("en-UK", options)
}

export function formatTimeDifference(date: Date) {
  if (!date || isNaN(date.getTime())) {
    return "Invalid date"
  }

  const now = new Date()

  const diff = now.getTime() - date.getTime()
  if (diff < 0) {
    return "Date is ahead of current time"
  }

  if (diff < 3600000) {
    const minutes = Math.floor(diff / 60000)
    return `${minutes} min${minutes > 1 ? "s" : ""} ago`
  } else if (diff < 86400000) {
    const hours = Math.floor(diff / 3600000) // Convert to hours
    return `${hours} hr${hours > 1 ? "s" : ""} ago`
  } else if (diff < 31536000000) {
    const days = Math.floor(diff / 86400000) // Convert to days
    return `${days} day${days > 1 ? "s" : ""} ago`
  } else {
    const years = Math.floor(diff / 31536000000) // Convert to years
    return `${years} year${years > 1 ? "s" : ""} ago`
  }
}
