import { Pipeline } from "@/types/pipeline"
import { ServiceRequest } from "@/types/service-request"

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
  const services: Pipeline[] = createDummyServices(25)
  return {
    services,
  }
}

export default useServices
