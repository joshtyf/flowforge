import { toast } from "@/components/ui/use-toast"
import { getServiceRequestDTO } from "@/lib/service"
import { ServiceRequestDTO } from "@/types/service-request"
import { useEffect, useState } from "react"

interface UseServiceRequestOptions {
  serviceRequestId: string
}

const useServiceRequestDTO = ({
  serviceRequestId,
}: UseServiceRequestOptions) => {
  const [serviceRequest, setServiceRequest] = useState<ServiceRequestDTO>()
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    setLoading(true)
    getServiceRequestDTO(serviceRequestId)
      .then(setServiceRequest)
      .catch((err) => {
        console.error(err)
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

export default useServiceRequestDTO
