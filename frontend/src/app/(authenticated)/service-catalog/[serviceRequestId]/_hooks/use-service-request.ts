import { toast } from "@/components/ui/use-toast"
import { getPipeline } from "@/lib/service"
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
  serviceRequestId: string
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

const useServiceRequest = ({ serviceRequestId }: UseServiceRequestProps) => {
  const [service, setService] = useState<Pipeline>()
  const [isLoadingForm, setIsLoadingForm] = useState(true)
  const [isSubmittingRequest, setIsSubmittingRequest] = useState(false)

  useEffect(() => {
    getPipeline(serviceRequestId)
      .then((data) => {
        data.form = DUMMY_SERVICE_REQUEST_FORM
        data.pipeline_description = "Pipeline description"
        setService(data)
      })
      .catch((err) => {
        console.log(err)
        toast({
          title: "Fetching Service Error",
          description:
            "Failed to fetch the service for service request. Please try again later.",
          variant: "destructive",
        })
      })
      .finally(() => {
        setIsLoadingForm(false)
      })
  }, [serviceRequestId])

  const handleSubmit = (data: IChangeEvent<object, RJSFSchema, object>) => {
    setIsSubmittingRequest(true)
    // TODO: Replace with API call
    // TODO: Add validations
    console.log(
      "Data submitted: ",
      "Service id: " + serviceRequestId,
      data.formData
    )
    setIsSubmittingRequest(false)
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
