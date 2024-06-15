import { toast } from "@/components/ui/use-toast"
import useOrganizationId from "@/hooks/use-organization-id"
import { getAllPipeline } from "@/lib/service"
import { useQuery } from "@tanstack/react-query"

const useServices = () => {
  const { organizationId } = useOrganizationId()

  const { isLoading, data: pipelines } = useQuery({
    queryKey: ["pipelines"],
    queryFn: () =>
      getAllPipeline(organizationId).catch((err) => {
        console.log(err)
        toast({
          title: "Fetching Services Error",
          description: "Failed to fetch the services. Please try again later.",
          variant: "destructive",
        })
      }),
  })

  return {
    services: pipelines,
    isServicesLoading: isLoading,
  }
}

export default useServices
