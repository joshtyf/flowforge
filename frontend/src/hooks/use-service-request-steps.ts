import { toast } from "@/components/ui/use-toast"
import { getServiceRequestSteps } from "@/lib/service"
import { createStepsFromObject } from "@/lib/utils"
import { useQuery } from "@tanstack/react-query"
import { useMemo } from "react"

type UseServiceRequestStepsProps = {
  serviceRequestId: string
}
const useServiceRequestSteps = ({
  serviceRequestId,
}: UseServiceRequestStepsProps) => {
  const { isLoading, data } = useQuery({
    queryKey: ["pipelines"],
    queryFn: () =>
      getServiceRequestSteps(serviceRequestId).catch((err) => {
        console.log(err)
        toast({
          title: "Fetching Services Error",
          description: "Failed to fetch the services. Please try again later.",
          variant: "destructive",
        })
      }),
    refetchInterval: 2000,
  })

  const stepsList = useMemo(
    () => createStepsFromObject(data?.first_step_name, data?.steps),
    [data]
  )

  return {
    steps: stepsList,
    isLoading,
  }
}

export default useServiceRequestSteps
