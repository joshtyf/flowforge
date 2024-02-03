import { ServiceRequestForm } from "@/types/service-request"
import { IChangeEvent } from "@rjsf/core"
import { RJSFSchema } from "@rjsf/utils"

interface UseServiceRequestProps {
  serviceRequestId: string
}

const useServiceRequest = ({ serviceRequestId }: UseServiceRequestProps) => {
  const serviceRequest: ServiceRequestForm = {
    name: "Sample Service Request",
    description: "Sample Service Request Form",
    form: {
      input: {
        title: "Input",
        type: "input",
        description: "Input Description with minimum length 1",
        minLength: 1,
        required: true,
      },
      select: {
        title: "Select Option",
        type: "select",
        description: "Dropdown selection with default value as Item 1",
        options: ["Item 1", "Item 2", "Item 3"],
        required: true,
      },
      checkboxes: {
        title: "Checkboxes",
        type: "checkboxes",
        description: "You can select more than 1 item",
        options: ["Item 1", "Item 2", "Item 3"],
        required: false,
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
