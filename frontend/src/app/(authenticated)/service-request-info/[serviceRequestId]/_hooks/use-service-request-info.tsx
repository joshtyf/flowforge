import useServiceRequest from "@/hooks/use-service-request"
import {
  convertServiceRequestFormToRJSFSchema,
  generateUiSchema,
} from "@/lib/utils"
import { useMemo } from "react"

interface UseServiceRequestInfoOptions {
  serviceRequestId: string
}

const useServiceRequestInfo = ({
  serviceRequestId,
}: UseServiceRequestInfoOptions) => {
  const { serviceRequest, isServiceRequestLoading } = useServiceRequest({
    serviceRequestId,
  })

  const uiSchema = useMemo(
    () => generateUiSchema(serviceRequest?.form),
    [serviceRequest]
  )
  const rjsfSchema = useMemo(
    () => convertServiceRequestFormToRJSFSchema(serviceRequest?.form),
    [serviceRequest]
  )

  return {
    pipelineName: serviceRequest?.pipeline_name,
    pipelineDescription: serviceRequest?.pipeline_description,
    formData: serviceRequest?.form_data,
    isServiceRequestLoading,
    uiSchema,
    rjsfSchema,
  }
}

export default useServiceRequestInfo
