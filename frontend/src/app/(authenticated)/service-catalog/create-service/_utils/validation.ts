import { isJson } from "@/lib/utils"
import { ServiceRequestWithSteps } from "@/types/service"
import {
  FormCheckboxes,
  FormComponent,
  FormComponentWithOptions,
  FormInput,
  FormSelect,
} from "@/types/sevice-request-form"

const isFormAttribute = (
  formItemAttribute: string
): formItemAttribute is keyof FormComponent => {
  return (
    formItemAttribute === "title" ||
    formItemAttribute === "description" ||
    formItemAttribute === "type" ||
    formItemAttribute === "required"
  )
}

const isInputFormAttribute = (
  formItemAttribute: string
): formItemAttribute is keyof FormInput => {
  return isFormAttribute(formItemAttribute) || formItemAttribute === "minLength"
}

const isSelectFormAttribute = (
  formItemAttribute: string
): formItemAttribute is keyof FormComponentWithOptions => {
  return (
    isFormAttribute(formItemAttribute) ||
    formItemAttribute === "options" ||
    formItemAttribute === "disabled" ||
    formItemAttribute === "default" ||
    formItemAttribute === "placeholder"
  )
}

const isCheckboxesFormAttribute = (formItemAttribute: string) => {
  return (
    isFormAttribute(formItemAttribute) ||
    formItemAttribute === "options" ||
    formItemAttribute === "disabled"
  )
}

function checkForUnexpectedFormAttributes(
  formItemName: string,
  formItem: object,
  errorMessages: string[]
) {
  const formItemObject = formItem as FormComponent
  for (const formItemAttribute in formItemObject) {
    if (
      formItemObject.type === "input" &&
      !isInputFormAttribute(formItemAttribute)
    ) {
      errorMessages.push(
        `Not allowed to add '${formItemAttribute}' attribute to input form item '${formItemName}'`
      )
    } else if (
      formItemObject.type === "select" &&
      !isSelectFormAttribute(formItemAttribute)
    ) {
      errorMessages.push(
        `Not allowed to add '${formItemAttribute}' attribute to select form item '${formItemName}'`
      )
    } else if (
      formItemObject.type === "checkboxes" &&
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
  if (title === undefined) {
    errorMessages.push(`title is missing from form item '${formItemName}'`)
  }

  if (description === undefined) {
    errorMessages.push(
      `description is missing from form item '${formItemName}'`
    )
  }

  if (type === undefined) {
    errorMessages.push(`type is missing from form item '${formItemName}'`)
  }
}

function checkForFormAttributeTypes(
  formItemName: string,
  formItem: object,
  errorMessages: string[]
) {
  const { title, description, type, required } = formItem as FormComponent
  if (typeof title !== "string" || typeof description !== "string") {
    errorMessages.push(
      `title and description of form item '${formItemName}' can only be string.`
    )
  }

  if (type !== "input" && type !== "select" && type !== "checkboxes") {
    errorMessages.push(
      `type of form item '${formItemName}' can only be 'input', 'select' or 'checkboxes'.`
    )
  }

  if (required !== undefined && typeof required !== "boolean") {
    errorMessages.push(
      `required of form item '${formItemName}' can only be boolean.`
    )
  }

  if (type === "input") {
    const inputItem = formItem as FormInput
    if (
      inputItem.minLength !== undefined &&
      typeof inputItem.minLength !== "number"
    ) {
      errorMessages.push(
        `minLength of form item '${formItemName}' can only be number.`
      )
    }
  } else if (type === "select" || type === "checkboxes") {
    const selectItem = formItem as FormComponentWithOptions
    if (
      selectItem.disabled !== undefined &&
      typeof selectItem.disabled !== "boolean"
    ) {
      errorMessages.push(
        `disabled of form item '${formItemName}' can only be boolean.`
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
    errorMessages.push(`Form item name for '${formItemName}' cannot be empty`)
  }
  if (type === "select" || type === "checkboxes") {
    const formWithOptions = formItem as FormComponentWithOptions
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

export function validateFormSchema(jsonString: string) {
  if (!isJson(jsonString)) {
    return []
  }

  const formJson = JSON.parse(jsonString)
  const errorList: string[] = []
  for (const formItemName in formJson) {
    const formItem = formJson[formItemName]

    if (
      formItem.type !== "input" &&
      formItem.type !== "select" &&
      formItem.type !== "checkboxes"
    ) {
      errorList.push(
        `Please define a type for form item '${formItemName}' (Only 'input', 'select' and 'checkboxes' type are supported)`
      )
    }

    checkForUnexpectedFormAttributes(formItemName, formItem, errorList)
    checkForMissingFormAttributes(formItemName, formItem, errorList)
    checkForFormAttributeTypes(formItemName, formItem, errorList)
    checkForEmptyAttributes(formItemName, formItem, errorList)
  }

  return errorList
}
