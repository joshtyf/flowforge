import { toast } from "@/components/ui/use-toast"
import { getServiceRequestSteps } from "@/lib/service"
import { useQuery } from "@tanstack/react-query"

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

  return {
    steps: data?.steps,
    firstStepName: data?.first_step_name,
    isLoading,
  }
}

export default useServiceRequestSteps
