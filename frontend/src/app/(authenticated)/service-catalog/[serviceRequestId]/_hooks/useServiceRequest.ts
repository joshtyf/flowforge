import { ServiceRequest } from "@/types/service"
import { IChangeEvent } from "@rjsf/core"
import { RJSFSchema } from "@rjsf/utils"

interface UseServiceRequestProps {
  serviceRequestId: string
}

const useServiceRequest = ({ serviceRequestId }: UseServiceRequestProps) => {
  // TODO: Once service request follows our own custom validation, write a conversion util function to convert service request to RJSF friendly object
  const serviceRequest: ServiceRequest = {
    name: "Sample Service Request",
    description: "Sample Service Request Form",
    form: {
      type: "object",
      required: ["input", "select"],
      properties: {
        input: {
          title: "Input",
          description: "Input Description with minimum length 1",
          type: "string",
          minLength: 1,
        },
        select: {
          type: "string",
          title: "Select Option",
          description: "Dropdown selection with default value as Item 1",
          enum: ["Item 1", "Item 2", "Item 3"],
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
