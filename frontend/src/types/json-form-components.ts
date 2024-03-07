type FormComponent = {
  title: string
  description: string
  type: "input" | "select" | "checkboxes"
  required?: boolean
}

type FormInput = FormComponent & {
  minLength?: number
  type: "input"
  placeholder?: string
}

type FormComponentWithOptions = FormComponent & {
  options: string[]
  disabled?: string[]
}

type FormSelect = FormComponentWithOptions & {
  type: "select"
  default?: string
  placeholder?: string
}

type FormCheckboxes = FormComponentWithOptions & {
  type: "checkboxes"
  required?: false
}

type JsonFormComponents = {
  [key: string]: FormInput | FormSelect | FormCheckboxes
}

export type {
  FormComponent,
  FormComponentWithOptions,
  FormInput,
  FormSelect,
  FormCheckboxes,
  JsonFormComponents,
}
