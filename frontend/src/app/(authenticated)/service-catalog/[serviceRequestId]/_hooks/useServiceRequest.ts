interface UseServiceRequestProps {
  serviceRequestId: string
}
const useServiceRequest = ({ serviceRequestId }: UseServiceRequestProps) => {
  console.log(serviceRequestId)
  const serviceRequest = {
    name: "Sample Service Request",
    description: "Sample Service Request Form",
    form: {},
  }
  return {
    serviceRequest,
  }
}

export default useServiceRequest
