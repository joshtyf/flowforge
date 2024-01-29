import { ServiceRequest } from "@/types/service"
import { IChangeEvent } from "@rjsf/core"
import { RJSFSchema } from "@rjsf/utils"

interface UseServiceRequestProps {
  serviceRequestId: string
}

const useServiceRequest = ({ serviceRequestId }: UseServiceRequestProps) => {
  const serviceRequest: ServiceRequest = {
    name: "Sample Service Request",
    description: "Sample Service Request Form",
    form: {
      type: "object",
      required: ["input"],
      properties: {
        input: {
          title: "Input",
          description: "Input Description with minimum length 1",
          type: "string",
          minLength: 1,
        },
        dropdown: {
          type: "string",
          title: "Dropdown list",
          description: "Dropdown selection with default value as Item 1",
          enum: ["Item 1", "Item 2", "Item 3"],
          default: "Item 1",
        },
        checkboxes: {
          type: "array",
          title: "Checkboxes",
          description: "You can select more than 1 item",
          items: {
            enum: ["Item 1", "Item 2", "Item 3"],
          },
          uniqueItems: true,
        },
      },
    },
  }

  const handleSubmit = (data: IChangeEvent<object, RJSFSchema, object>) => {
    // TODO: Replace with API call
    // TODO: Add validations
    console.log(
      "Data submitted: ",
      "Service id: " + serviceRequestId,
      data.formData
    )
  }
  return {
    serviceRequest,
    handleSubmit,
  }
}

export default useServiceRequest
