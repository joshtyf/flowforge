import { toast } from "@/components/ui/use-toast"
import { createPipeline } from "@/lib/service"
import { isJson } from "@/lib/utils"
import {
  FormCheckboxes,
  FormFieldType,
  FormInput,
  FormSelect,
  JsonFormComponents,
} from "@/types/json-form-components"
import { Pipeline } from "@/types/pipeline"
import { zodResolver } from "@hookform/resolvers/zod"
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime"
import { KeyboardEvent, useState } from "react"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { crossValidateSchema, validateFormSchema } from "../_utils/validation"

const DEFAULT_FORM: JsonFormComponents = {
  fields: [
    {
      name: "",
      title: "",
      description: "",
      type: FormFieldType.INPUT,
      required: true,
      min_length: 1,
      placeholder: "Enter text...",
    } as FormInput,
    {
      name: "",
      title: "",
      description: "",
      type: FormFieldType.SELECT,
      required: true,
      options: ["Option 1", "Option 2", "Option 3"],
      default: "Option 1",
      placeholder: "Select an option",
    } as FormSelect,
    {
      name: "",
      title: "",
      description: "",
      type: FormFieldType.CHECKBOXES,
      options: ["Option 1", "Option 2", "Option 3"],
    } as FormCheckboxes,
  ],
}

const DEFAULT_PIPELINE = {
  version: 1,
  first_step_name: "step1",
  steps: [
    {
      step_name: "step1",
      step_type: "API",
      next_step_name: "step2",
      prev_step_name: "",
      parameters: {
        method: "GET",
        url: "https://example.com",
      },
      is_terminal_step: false,
    },
    {
      step_name: "step2",
      step_type: "WAIT_FOR_APPROVAL",
      next_step_name: "",
      prev_step_name: "step1",
      parameters: {},
      is_terminal_step: true,
    },
  ],
}

const createServiceSchema = z
  .object({
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
    pipeline: z
      .string()
      .min(1, "Pipeline Steps Schema is required")
      .refine((arg) => isJson(arg), {
        message: "Ensure that Form is valid JSON Schema",
      }),
  })
  .superRefine((val, ctx) => {
    crossValidateSchema(val.form, val.pipeline).forEach((error) => {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: error,
        path: ["form"],
      })
    })
  })

interface UseCreateServiceProps {
  router: AppRouterInstance
}

const useCreateService = ({ router }: UseCreateServiceProps) => {
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [submitted, setSubmitted] = useState(false)
  const form = useForm<z.infer<typeof createServiceSchema>>({
    resolver: zodResolver(createServiceSchema),
    defaultValues: {
      name: "",
      description: "",
      form: JSON.stringify(DEFAULT_FORM, null, 4),
      pipeline: JSON.stringify(DEFAULT_PIPELINE, null, 4),
    },
  })

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

  const handleSubmitForm = (values: z.infer<typeof createServiceSchema>) => {
    const { description, form, name, pipeline } = values

    const formJson: JsonFormComponents = JSON.parse(form)

    const pipelineJson: Pipeline = {
      pipeline_name: name,
      pipeline_description: description,
      form: formJson,
      ...JSON.parse(pipeline),
    }

    setIsSubmitting(true)

    createPipeline(pipelineJson)
      .then(() => {
        toast({
          title: "Service Creation Successful",
          description: "Redirecting to service catalog...",
          variant: "success",
        })
        setSubmitted(true)
        setTimeout(() => router.push("/service-catalog"), 1500)
      })
      .catch((err) => {
        console.error(err)
        toast({
          title: "Submission Error",
          description:
            "Encountered error submitting pipeline. Please try again later",
          variant: "destructive",
        })
      })
      .finally(() => setIsSubmitting(false))
  }

  return {
    form,
    handleTextAreaTabKeyDown,
    handleSubmitForm,
    submitted,
    isSubmitting,
  }
}

export default useCreateService
