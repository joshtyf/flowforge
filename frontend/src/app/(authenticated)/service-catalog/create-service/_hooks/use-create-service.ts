import { ServiceRequestWithSteps } from "@/types/service"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { isJson } from "@/lib/utils"
import { KeyboardEvent, KeyboardEventHandler } from "react"
import { validateFormSchema } from "../_utils/validation"

const DEFAULT_FORM = {
  input: { title: "", description: "", type: "input", required: true },
  select: {
    title: "",
    description: "",
    type: "select",
    required: true,
    options: ["Option 1", "Option 2", "Option 3"],
  },
  checkBoxes: {
    title: "",
    description: "",
    type: "checkboxes",
    required: false,
    options: ["Option 1", "Option 2", "Option 3"],
  },
}

const DEFAULT_STEPS = {
  Approval: {
    name: "Approval",
    type: "approval",
    next: "",
    start: true,
  },
}

const createServiceSchema = z.object({
  name: z.string().min(1, "Name is required"),
  description: z.string(),
  form: z
    .string()
    .min(1, "Form Schema is required")
    .superRefine((val, ctx) => {
      const errorList = validateFormSchema(val)
      if (errorList.length > 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: errorList.join(" , "),
        })
      }
    })
    .refine((arg) => isJson(arg), {
      message: "Ensure that Form is valid JSON Schema",
    }),
  steps: z
    .string()
    .min(1, "Pipeline Steps Schema is required")
    .refine((arg) => isJson(arg), {
      message: "Ensure that Form is valid JSON Schema",
    }),
})
const useCreateService = () => {
  const form = useForm<z.infer<typeof createServiceSchema>>({
    resolver: zodResolver(createServiceSchema),
    defaultValues: {
      name: "",
      description: "",
      form: JSON.stringify(DEFAULT_FORM, null, 4),
      steps: JSON.stringify(DEFAULT_STEPS, null, 4),
    },
  })

  const handleSubmitForm = (values: z.infer<typeof createServiceSchema>) => {
    const { description, form, name, steps } = values

    // TODO: Replace with API call
    console.log("Submitting:", {
      name,
      description,
      form: JSON.parse(form),
      steps: JSON.parse(steps),
    })
  }

  function handleTextAreaTabKeyDown(event: KeyboardEvent): void {
    if (event.key == "Tab") {
      event.preventDefault()
      const htmlTextElement = event.target as HTMLTextAreaElement
      const start = htmlTextElement.selectionStart
      const end = htmlTextElement.selectionEnd

      htmlTextElement.value =
        htmlTextElement.value.substring(0, start) +
        "\t" +
        htmlTextElement.value.substring(end)

      htmlTextElement.selectionStart = htmlTextElement.selectionEnd = start + 1
    }
  }

  return {
    form,
    handleSubmitForm,
    handleTextAreaTabKeyDown,
  }
}

export default useCreateService
