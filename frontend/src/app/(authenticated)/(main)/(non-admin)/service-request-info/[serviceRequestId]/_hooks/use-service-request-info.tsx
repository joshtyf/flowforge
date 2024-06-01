import useServiceRequestDTO from "@/hooks/use-service-request-dto"
import { generateUiSchema } from "@/lib/rjsf-utils"
import { convertServiceRequestFormToRJSFSchema } from "@/lib/rjsf-utils"
import { useMemo } from "react"

interface UseServiceRequestInfoOptions {
  serviceRequestId: string
}

const useServiceRequestInfo = ({
  serviceRequestId,
}: UseServiceRequestInfoOptions) => {
  const { serviceRequest, isServiceRequestLoading } = useServiceRequestDTO({
    serviceRequestId,
  })

  const uiSchema = useMemo(
    () => generateUiSchema(serviceRequest?.pipeline?.form),
    [serviceRequest]
  )
  const rjsfSchema = useMemo(
    () => convertServiceRequestFormToRJSFSchema(serviceRequest?.pipeline?.form),
    [serviceRequest]
  )

  return {
    pipelineName: serviceRequest?.pipeline.name ?? "",
    pipelineDescription: "",
    formData: serviceRequest?.service_request.form_data ?? {},
    isServiceRequestLoading,
    uiSchema,
    rjsfSchema,
  }
}

export default useServiceRequestInfo
