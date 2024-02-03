import { ServiceRequest } from "@/types/service"

const createDummyServices = (noOfServices: number) => {
  const services: ServiceRequest[] = []
  for (let i = 0; i < noOfServices; i++) {
    services.push({
      id: i + 1,
      name: `Service ${i + 1}`,
      description: `Description ${i + 1}`,
      form: {},
    })
  }
  return services
}

const useServices = () => {
  const services: ServiceRequest[] = createDummyServices(25)
  return {
    services,
  }
}

export default useServices
