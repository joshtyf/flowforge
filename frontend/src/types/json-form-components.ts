export enum FormFieldType {
  INPUT = "input",
  SELECT = "select",
  CHECKBOXES = "checkboxes",
}

type FormComponent = {
  name: string
  title: string
  description: string
  type: FormFieldType
}

type FormInput = FormComponent & {
  type: FormFieldType.INPUT
  required?: boolean
  min_length?: number
  placeholder?: string
}

type Options = {
  options: string[]
}

type FormSelect = FormComponent &
  Options & {
    type: FormFieldType.SELECT
    required?: boolean
    default?: string
    placeholder?: string
  }

type FormCheckboxes = FormComponent &
  Options & {
    type: FormFieldType.CHECKBOXES
  }

type JsonFormComponents = {
  fields: (FormInput | FormSelect | FormCheckboxes)[]
}

export type {
  FormComponent,
  FormInput,
  Options,
  FormSelect,
  FormCheckboxes,
  JsonFormComponents,
}
