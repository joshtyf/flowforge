import { isArrayValuesString, isArrayValuesUnique, isJson } from "@/lib/utils"
import {
  FormCheckboxes,
  FormComponent,
  FormFieldType,
  FormInput,
  FormSelect,
  JsonFormComponents,
  Options,
} from "@/types/json-form-components"

const isNameUnique = (field: FormComponent[]) => {
  return new Set(field.map((f) => f.name)).size === field.length
}

const checkForName = (formItems: FormComponent[], errorMessages: string[]) => {
  for (let i = 0; i < formItems.length; i++) {
    const formItem = formItems[i]
    const { name } = formItem
    if (!name) {
      errorMessages.push(`Please define a name for form item ${i + 1}`)
    }
  }

  if (!isNameUnique(formItems)) {
    errorMessages.push("Please ensure that name is unique for the form items.")
  }
}

const isFormAttribute = (
  formItemAttribute: string
): formItemAttribute is keyof FormComponent => {
  return (
    formItemAttribute === "name" ||
    formItemAttribute === "title" ||
    formItemAttribute === "description" ||
    formItemAttribute === "type"
  )
}

const isInputFormAttribute = (
  formItemAttribute: string
): formItemAttribute is keyof FormInput => {
  return (
    isFormAttribute(formItemAttribute) ||
    formItemAttribute === "min_length" ||
    formItemAttribute === "required" ||
    formItemAttribute === "placeholder"
  )
}

const isSelectFormAttribute = (
  formItemAttribute: string
): formItemAttribute is keyof FormSelect => {
  return (
    isFormAttribute(formItemAttribute) ||
    formItemAttribute === "required" ||
    formItemAttribute === "options" ||
    formItemAttribute === "default" ||
    formItemAttribute === "placeholder"
  )
}

const isCheckboxesFormAttribute = (
  formItemAttribute: string
): formItemAttribute is keyof FormCheckboxes => {
  return isFormAttribute(formItemAttribute) || formItemAttribute === "options"
}

function checkForUnexpectedFormAttributes(
  formItemName: string,
  formItem: object,
  errorMessages: string[]
) {
  const formItemObject = formItem as FormComponent
  for (const formItemAttribute in formItemObject) {
    if (
      formItemObject.type === FormFieldType.INPUT &&
      !isInputFormAttribute(formItemAttribute)
    ) {
      errorMessages.push(
        `Not allowed to add '${formItemAttribute}' attribute to input form item '${formItemName}'`
      )
    } else if (
      formItemObject.type === FormFieldType.SELECT &&
      !isSelectFormAttribute(formItemAttribute)
    ) {
      errorMessages.push(
        `Not allowed to add '${formItemAttribute}' attribute to select form item '${formItemName}'`
      )
    } else if (
      formItemObject.type === FormFieldType.CHECKBOXES &&
      !isCheckboxesFormAttribute(formItemAttribute)
    ) {
      errorMessages.push(
        `Not allowed to add '${formItemAttribute}' attribute to checkboxes form item '${formItemName}'`
      )
    }
  }
}

function checkForMissingFormAttributes(
  formItemName: string,
  formItem: object,
  errorMessages: string[]
) {
  const { title, description, type } = formItem as FormComponent

  if (!title) {
    errorMessages.push(`title is missing from form item '${formItemName}'`)
  }

  // Check for undefined equality directly as empty string is allowed
  if (description === undefined) {
    errorMessages.push(
      `description is missing from form item '${formItemName}'`
    )
  }

  if (!type) {
    errorMessages.push(`type is missing from form item '${formItemName}'`)
  }
}

function checkForFormAttributeTypes(
  formItemName: string,
  formItem: object,
  errorMessages: string[]
) {
  const { title, description, type } = formItem as FormComponent
  if (typeof title !== "string" || typeof description !== "string") {
    errorMessages.push(
      `title and description of form item '${formItemName}' can only be string.`
    )
  }

  switch (type) {
    case FormFieldType.INPUT: {
      const inputItem = formItem as FormInput
      if (
        inputItem.min_length !== undefined &&
        typeof inputItem.min_length !== "number"
      ) {
        errorMessages.push(
          `min_length of form item '${formItemName}' can only be number.`
        )
      }

      if (
        inputItem.required !== undefined &&
        typeof inputItem.required !== "boolean"
      ) {
        errorMessages.push(
          `required of form item '${formItemName}' can only be boolean.`
        )
      }
      break
    }
    case FormFieldType.SELECT: {
      const selectItem = formItem as FormSelect
      if (
        selectItem.required !== undefined &&
        typeof selectItem.required !== "boolean"
      ) {
        errorMessages.push(
          `required of form item '${formItemName}' can only be boolean.`
        )
      }
      break
    }
    case FormFieldType.CHECKBOXES: {
      const checkboxesItem = formItem as FormCheckboxes

      break
    }
    default: {
      errorMessages.push(
        `type of form item '${formItemName}' can only be 'input', 'select' or 'checkboxes'.`
      )
    }
  }
}

function checkForEmptyAttributes(
  formItemName: string,
  formItem: object,
  errorMessages: string[]
) {
  const { title, type } = formItem as FormComponent
  if (!title) {
    errorMessages.push(`Form item title for '${formItemName}' cannot be empty`)
  }
  if (type === FormFieldType.SELECT || type === FormFieldType.CHECKBOXES) {
    const formWithOptions = formItem as Options
    const { options } = formWithOptions
    if (!options) {
      errorMessages.push(
        `Please define an array of options for '${formItemName}'`
      )
    } else if (options.length === 0) {
      errorMessages.push(
        `Form item options for '${formItemName}' cannot be empty`
      )
    }
  }
}

function checkForFormAttributesValues(
  formItemName: string,
  formItem: object,
  errorMessages: string[]
) {
  const { type } = formItem as FormComponent

  switch (type) {
    case FormFieldType.INPUT: {
      break
    }
    case FormFieldType.SELECT:
    case FormFieldType.CHECKBOXES: {
      const { options } = formItem as Options
      if (!isArrayValuesUnique(options)) {
        errorMessages.push(`Option values for '${formItemName}' must be unique`)
      }

      if (!isArrayValuesString(options)) {
        errorMessages.push(
          `Option values for '${formItemName}' can only be string`
        )
      }
      break
    }
    default:
      break
  }
}

export function validateFormSchema(jsonString: string) {
  if (!isJson(jsonString)) {
    return []
  }

  const formJson: JsonFormComponents = JSON.parse(
    jsonString
  ) as JsonFormComponents
  const errorList: string[] = []
  if (!formJson.fields) {
    errorList.push("Please define 'fields' in the form")
    return errorList
  }

  checkForName(formJson.fields, errorList)

  for (const formItem of formJson.fields) {
    const formItemName = formItem.name
    if (
      formItem.type !== FormFieldType.INPUT &&
      formItem.type !== FormFieldType.SELECT &&
      formItem.type !== FormFieldType.CHECKBOXES
    ) {
      errorList.push(
        `Please define a type for form item '${formItemName}' (Only 'input', 'select' and 'checkboxes' type are supported)`
      )
    }

    checkForUnexpectedFormAttributes(formItem.name, formItem, errorList)
    checkForMissingFormAttributes(formItem.name, formItem, errorList)
    checkForFormAttributeTypes(formItem.name, formItem, errorList)
    checkForEmptyAttributes(formItem.name, formItem, errorList)
    checkForFormAttributesValues(formItem.name, formItem, errorList)
  }

  return errorList
}

/**
 * Validates the schema of a form against a pipeline.
 * @param form - The JSON schema of the form.
 * @param pipeline - The pipeline string containing parameters.
 * @returns An array of error messages indicating missing parameters in the form schema.
 */
export function crossValidateSchema(form: string, pipeline: string): string[] {
  const pipelineParameters = pipeline
    .match(/\${(.*?)}/g)
    ?.map((match) => match.slice(2, -1)) // strip ${ and }
  if (!pipelineParameters) return []
  const formJson: JsonFormComponents = JSON.parse(form) as JsonFormComponents
  return pipelineParameters
    .filter(
      (pipelineParameter) =>
        !formJson.fields.find((field) => field.name === pipelineParameter)
    )
    .map((undefinedPipelineParameter) => {
      return `Pipeline schema parameter '${undefinedPipelineParameter}' not found in form schema. Declare a field with the name '${undefinedPipelineParameter}' in the form schema.`
    })
}
