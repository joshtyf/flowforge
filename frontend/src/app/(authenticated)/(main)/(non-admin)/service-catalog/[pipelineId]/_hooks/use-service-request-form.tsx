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
import RequestCreatedTextWithCountdown from "../_components/request-created-text-with-countdown"
import { useRouter } from "next/navigation"

interface UseServiceRequestFormOptions {
  pipelineId: string
}

const useServiceRequestForm = ({
  pipelineId,
}: UseServiceRequestFormOptions) => {
  const { pipeline: service, isPipelineLoading: isLoadingForm } = usePipeline({
    pipelineId,
  })
  const [isSubmittingRequest, setIsSubmittingRequest] = useState(false)
  const { organizationId } = useOrganizationId()
  const router = useRouter()
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
      .then(() => {
        toast({
          title: "Request Submission Successful",
          description: <RequestCreatedTextWithCountdown />,
          variant: "success",
        })
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
