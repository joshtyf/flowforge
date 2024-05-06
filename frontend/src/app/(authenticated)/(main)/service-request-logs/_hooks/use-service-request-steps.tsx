import useServiceRequest from "@/hooks/use-service-request"
import { createStepsFromObject } from "@/lib/utils"
import { useEffect, useMemo, useState } from "react"

type UseServiceRequestLogsOptions = {
  serviceRequestId: string
}
const useServiceRequestSteps = ({
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

  const [currentStep, setCurrentStep] = useState<string>("")

  useEffect(() => {
    if (serviceRequest) {
      // Initialize logs page with the first step logs
      setCurrentStep(serviceRequest.first_step_name)
    }
  }, [serviceRequest])

  return {
    serviceRequestLogs: [],
    serviceRequestSteps,
    currentStep,
    isLoading: isServiceRequestLoading,
    handleCurrentStepChange: setCurrentStep,
  }
}

export default useServiceRequestSteps
