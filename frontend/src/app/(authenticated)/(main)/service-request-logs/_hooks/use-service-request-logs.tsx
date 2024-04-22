import useServiceRequest from "@/hooks/use-service-request"

type UseServiceRequestLogsOptions = {
  serviceRequestId: string
}
const useServiceRequestLogs = ({
  serviceRequestId,
}: UseServiceRequestLogsOptions) => {
  const { serviceRequest, isServiceRequestLoading } = useServiceRequest({
    serviceRequestId,
  })

  return {
    serviceRequestLogs: [],
    serviceRequestSteps: serviceRequest?.steps ?? [],
    isLoading: isServiceRequestLoading,
  }
}

export default useServiceRequestLogs
