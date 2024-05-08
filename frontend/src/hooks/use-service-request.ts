import { toast } from "@/components/ui/use-toast"
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
      .catch((err) => {
        console.log(err)
        toast({
          title: "Fetching Service Request Error",
          description: `Failed to fetch Service Request Info. Please try again later.`,
          variant: "destructive",
        })
      })
      .finally(() => setLoading(false))
  }, [serviceRequestId])

  return {
    serviceRequest,
    isServiceRequestLoading: loading,
  }
}

export default useServiceRequest
