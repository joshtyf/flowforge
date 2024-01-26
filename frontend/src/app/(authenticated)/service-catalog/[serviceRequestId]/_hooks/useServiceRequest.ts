import { ServiceRequest } from "@/types/service"
import { IChangeEvent } from "@rjsf/core"
import { RJSFSchema } from "@rjsf/utils"

interface UseServiceRequestProps {
  serviceRequestId: string
}

const useServiceRequest = ({ serviceRequestId }: UseServiceRequestProps) => {
  console.log(serviceRequestId)
  const serviceRequest: ServiceRequest = {
    name: "Sample Service Request",
    description: "Sample Service Request Form",
    form: {
      type: "object",
      required: ["resourceName", "type"],
      properties: {
        resourceName: {
          title: "Resource Name",
          description: "Test Description",
          type: "string",
          minLength: 1,
        },
        type: {
          title: "Type",
          type: "string",
          minLength: 1,
        },
      },
    },
  }

  const handleSubmit = (data: IChangeEvent<object, RJSFSchema, object>) => {
    // TODO: Replace with API call
    // TODO: Add validations
    console.log("Data submitted: ", data.formData)
  }
  return {
    serviceRequest,
    handleSubmit,
  }
}

export default useServiceRequest
