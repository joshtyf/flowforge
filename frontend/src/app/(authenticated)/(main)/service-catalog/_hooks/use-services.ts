import { toast } from "@/components/ui/use-toast"
import { getAllPipeline } from "@/lib/service"
import { useQuery } from "@tanstack/react-query"

const useServices = () => {
  const { isLoading, data: pipelines } = useQuery({
    queryKey: ["pipelines"],
    queryFn: () =>
      getAllPipeline().catch((err) => {
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
