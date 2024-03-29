import { toast } from "@/components/ui/use-toast"
import usePipeline from "@/hooks/use-pipeline"
import { createServiceRequest, getPipeline } from "@/lib/service"
import {
  convertServiceRequestFormToRJSFSchema,
  generateUiSchema,
} from "@/lib/utils"
import { JsonFormComponents } from "@/types/json-form-components"
import { Pipeline } from "@/types/pipeline"
import { IChangeEvent } from "@rjsf/core"
import { RJSFSchema } from "@rjsf/utils"
import { useEffect, useMemo, useState } from "react"

interface UseServiceRequestFormOptions {
  pipelineId: string
}

const DUMMY_SERVICE_REQUEST_FORM: JsonFormComponents = {
  input: {
    title: "Input",
    type: "input",
    description: "Input Description with minimum length 1",
    minLength: 1,
    required: true,
    placeholder: "Input placeholder...",
  },
  select: {
    title: "Select Option",
    type: "select",
    placeholder: "Select placeholder",
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

const useServiceRequestForm = ({
  pipelineId,
}: UseServiceRequestFormOptions) => {
  const { pipeline: service, isPipelineLoading: isLoadingForm } = usePipeline({
    pipelineId,
  })
  const [isSubmittingRequest, setIsSubmittingRequest] = useState(false)

  const handleCreateServiceRequest = (
    data: IChangeEvent<object, RJSFSchema, object>
  ) => {
    const { formData } = data
    setIsSubmittingRequest(true)
    createServiceRequest(pipelineId, formData, service?.version)
      .then((data) => {
        toast({
          title: "Request Submission Successful",
          description:
            "Please check the dashboard for the status of the request.",
          variant: "success",
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

  const uiSchema = useMemo(() => generateUiSchema(service?.form), [service])
  const rjsfSchema = useMemo(
    () => convertServiceRequestFormToRJSFSchema(service?.form),
    [service]
  )

  return {
    pipelineName: service?.pipeline_name,
    pipelineDescription: service?.pipeline_description,
    rjsfSchema,
    uiSchema,
    handleSubmit: handleCreateServiceRequest,
    isLoadingForm,
    isSubmittingRequest,
  }
}

export default useServiceRequestForm
