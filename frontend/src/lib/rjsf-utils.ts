import { FormFieldType, JsonFormComponents } from "@/types/json-form-components"
import { RJSFSchema, UiSchema } from "@rjsf/utils"

export const convertServiceRequestFormToRJSFSchema = (
  jsonFormComponents?: JsonFormComponents
) => {
  const schema: RJSFSchema = {
    type: "object",
  }

  const properties: { [key: string]: object } = {}
  const required: string[] = []

  for (const component of jsonFormComponents?.fields ?? []) {
    const hasRequiredField =
      component.type === FormFieldType.INPUT ||
      component.type === FormFieldType.SELECT
    // To create required array
    if (hasRequiredField && component.required) {
      required.push(component.name)
    }

    switch (component.type) {
      case FormFieldType.INPUT:
        properties[component.name] = {
          type: "string",
          title: component.title,
          description: component.description,
          minLength: component.min_length,
        }
        break
      case FormFieldType.SELECT:
        properties[component.name] = {
          type: "string",
          title: component.title,
          description: component.description,
          enum: component.options,
        }
        break
      case FormFieldType.CHECKBOXES:
        properties[component.name] = {
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

  for (const itemOptions of jsonFormComponents?.fields ?? []) {
    switch (itemOptions.type) {
      case FormFieldType.INPUT:
        uiSchema[itemOptions.name] = {
          "ui:placeholder": itemOptions.placeholder,
        }
        break
      case FormFieldType.SELECT:
        uiSchema[itemOptions.name] = {
          "ui:placeholder": itemOptions.placeholder ?? "Select an item...",
        }
        break
      case FormFieldType.CHECKBOXES:
        uiSchema[itemOptions.name] = {
          "ui:widget": "checkboxes",
        }
        break

      default:
        break
    }
  }

  return uiSchema
}
