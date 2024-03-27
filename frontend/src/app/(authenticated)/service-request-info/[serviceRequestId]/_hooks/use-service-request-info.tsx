import useServiceRequest from "@/hooks/use-service-request"
import { getServiceRequest } from "@/lib/service"
import {
  convertServiceRequestFormToRJSFSchema,
  generateUiSchema,
} from "@/lib/utils"
import { ServiceRequest } from "@/types/service-request"
import { useEffect, useMemo, useState } from "react"

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
