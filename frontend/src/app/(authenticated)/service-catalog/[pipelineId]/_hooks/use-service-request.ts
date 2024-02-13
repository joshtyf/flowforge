import { toast } from "@/components/ui/use-toast"
import { createServiceRequest, getPipeline } from "@/lib/service"
import {
  convertServiceRequestFormToRJSFSchema,
  generateUiSchema,
} from "@/lib/utils"
import { JsonFormComponents } from "@/types/json-form-components"
import { Pipeline } from "@/types/pipeline"
import { IChangeEvent } from "@rjsf/core"
import { RJSFSchema } from "@rjsf/utils"
import { useEffect, useState } from "react"

interface UseServiceRequestProps {
  pipelineId: string
}

const DUMMY_SERVICE_REQUEST_FORM: JsonFormComponents = {
  input: {
    title: "Input",
    type: "input",
    description: "Input Description with minimum length 1",
    minLength: 1,
    required: true,
  },
  select: {
    title: "Select Option",
    type: "select",
    description: "Dropdown selection with default value as Item 1",
    options: ["Item 1", "Item 2", "Item 3"],
    required: true,
  },
  checkboxes: {
    title: "Checkboxes",
    type: "checkboxes",
    description: "You can select more than 1 item",
    options: ["Item 1", "Item 2", "Item 3"],
    required: false,
  },
}

const useServiceRequest = ({ pipelineId }: UseServiceRequestProps) => {
  const [service, setService] = useState<Pipeline>()
  const [isLoadingForm, setIsLoadingForm] = useState(true)
  const [isSubmittingRequest, setIsSubmittingRequest] = useState(false)

  useEffect(() => {
    getPipeline(pipelineId)
      .then((data) => {
        // TODO: To remove once data contains form and description
        data.form = DUMMY_SERVICE_REQUEST_FORM
        data.pipeline_description = "Pipeline description"
        setService(data)
      })
      .catch((err) => {
        console.log(err)
        toast({
          title: "Fetching Service Error",
          description: "Failed to fetch the service for service request.",
          variant: "destructive",
        })
      })
      .finally(() => {
        setIsLoadingForm(false)
      })
  }, [pipelineId])

  const handleSubmit = (data: IChangeEvent<object, RJSFSchema, object>) => {
    setIsSubmittingRequest(true)
    // TODO: Add validations if needed
    // TODO: Add submission of form data
    createServiceRequest(pipelineId)
      .then((data) => {
        toast({
          title: "Request Submission Successful",
          description:
            "Please check the dashboard for the status of the request.",
        })
        console.log("Response: ", data)
      })
      .catch((err) => {
        console.log(err)
        toast({
          title: "Request Submission Error",
          description: "Failed to submit the service request.",
          variant: "destructive",
        })
      })
      .finally(() => {
        setIsSubmittingRequest(false)
      })
  }

  const uiSchema = generateUiSchema(service?.form)
  const rjsfSchema = convertServiceRequestFormToRJSFSchema(service?.form)

  return {
    service,
    rjsfSchema,
    uiSchema,
    handleSubmit,
    isLoadingForm,
    isSubmittingRequest,
  }
}

export default useServiceRequest
