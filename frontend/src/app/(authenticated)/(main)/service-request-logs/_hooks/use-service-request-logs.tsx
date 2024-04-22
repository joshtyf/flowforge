import useServiceRequest from "@/hooks/use-service-request"

type UseServiceRequestLogsOptions = {
  serviceRequestId: string
}
const useServiceRequestLogs = ({
  serviceRequestId,
}: UseServiceRequestLogsOptions) => {
  const serviceRequest = useServiceRequest({ serviceRequestId })

  return {
    serviceRequestLogs: [],
    serviceRequest,
  }
}

export default useServiceRequestLogs
