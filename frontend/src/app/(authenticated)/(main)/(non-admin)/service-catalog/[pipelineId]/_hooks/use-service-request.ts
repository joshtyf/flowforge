import { toast } from "@/components/ui/use-toast"
import useOrganizationId from "@/hooks/use-organization-id"
import usePipeline from "@/hooks/use-pipeline"
import { createServiceRequest } from "@/lib/service"
import { generateUiSchema } from "@/lib/rjsf-utils"
import { convertServiceRequestFormToRJSFSchema } from "@/lib/rjsf-utils"
import { FormFieldType, JsonFormComponents } from "@/types/json-form-components"
import { IChangeEvent } from "@rjsf/core"
import { RJSFSchema } from "@rjsf/utils"
import { useMemo, useState } from "react"

interface UseServiceRequestFormOptions {
  pipelineId: string
}

const DUMMY_SERVICE_REQUEST_FORM: JsonFormComponents = {
  fields: [
    {
      name: "input",
      title: "Input",
      description: "",
      type: FormFieldType.INPUT,
      required: true,
      placeholder: "Enter text...",
      min_length: 1,
    },
    {
      name: "select",
      title: "Select",
      description: "",
      type: FormFieldType.SELECT,
      required: true,
      placeholder: "Select an option",
      options: ["Option 1", "Option 2", "Option 3"],
      default: "Option 1",
    },
    {
      name: "checkbox",
      title: "Checkbox",
      description: "",
      type: FormFieldType.CHECKBOXES,
      options: ["Option 1", "Option 2", "Option 3"],
    },
  ],
}

const useServiceRequestForm = ({
  pipelineId,
}: UseServiceRequestFormOptions) => {
  const { pipeline: service, isPipelineLoading: isLoadingForm } = usePipeline({
    pipelineId,
  })
  const [isSubmittingRequest, setIsSubmittingRequest] = useState(false)
  const { organizationId } = useOrganizationId()
  const handleCreateServiceRequest = (
    data: IChangeEvent<object, RJSFSchema, object>
  ) => {
    const { formData } = data
    setIsSubmittingRequest(true)
    if (!formData) {
      toast({
        title: "No Form Data Error",
        description: "Form data cannot be found.",
        variant: "destructive",
      })
      return
    }
    createServiceRequest(organizationId, pipelineId, formData, service?.version)
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
