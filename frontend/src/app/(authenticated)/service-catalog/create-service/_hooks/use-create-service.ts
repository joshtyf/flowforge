import { ServiceRequestWithSteps } from "@/types/service"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { isJson } from "@/lib/utils"
import { KeyboardEvent, KeyboardEventHandler } from "react"

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
  form: z.string(),
  steps: z.string(),
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

  // TODO: Add validation of JSON object on edit/submit
  return {
    form,
    handleSubmitForm,
    handleTextAreaTabKeyDown,
  }
}

export default useCreateService
