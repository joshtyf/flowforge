import useServiceRequest from "@/hooks/use-service-request"
import { createStepsFromObject } from "@/lib/utils"
import { useMemo } from "react"

type UseServiceRequestLogsOptions = {
  serviceRequestId: string
}
const useServiceRequestLogs = ({
  serviceRequestId,
}: UseServiceRequestLogsOptions) => {
  const { serviceRequest, isServiceRequestLoading } = useServiceRequest({
    serviceRequestId,
  })

  const serviceRequestSteps = useMemo(
    () =>
      createStepsFromObject(
        serviceRequest?.first_step_name ?? "",
        serviceRequest?.steps
      ),
    [serviceRequest?.steps, serviceRequest?.first_step_name]
  )

  return {
    serviceRequestLogs: [],
    serviceRequestSteps,
    isLoading: isServiceRequestLoading,
  }
}

export default useServiceRequestLogs
