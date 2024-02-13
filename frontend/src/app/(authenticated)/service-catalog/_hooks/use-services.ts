import { toast } from "@/components/ui/use-toast"
import { getAllPipeline } from "@/lib/service"
import { Pipeline } from "@/types/pipeline"
import { ServiceRequest } from "@/types/service-request"
import { useQuery } from "@tanstack/react-query"

const createDummyServices = (noOfServices: number) => {
  const services: Pipeline[] = []
  for (let i = 0; i < noOfServices; i++) {
    services.push({
      id: (i + 1).toString(),
      pipeline_name: `Service ${i + 1}`,
      // TODO: Add description once available
      // description: `Description ${i + 1}`,
    })
  }
  return services
}

const useServices = () => {
  // const services: Pipeline[] = createDummyServices(25)
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
  }
}

export default useServices
