import { getServiceRequest } from "@/lib/service"
import { ServiceRequest } from "@/types/service-request"
import { useEffect, useState } from "react"

interface UseServiceRequestOptions {
  serviceRequestId: string
}

const useServiceRequest = ({ serviceRequestId }: UseServiceRequestOptions) => {
  const [serviceRequest, setServiceRequest] = useState<ServiceRequest>()
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    setLoading(true)
    getServiceRequest(serviceRequestId)
      .then(setServiceRequest)
      .finally(() => setLoading(false))
  }, [serviceRequestId])

  return {
    serviceRequest,
    isServiceRequestLoading: loading,
  }
}

export default useServiceRequest
