import { ServiceRequest } from "@/types/service"
import { IChangeEvent } from "@rjsf/core"
import { RJSFSchema } from "@rjsf/utils"

interface UseServiceRequestProps {
  serviceRequestId: string
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const ORIGINAL_RJSF_OBJECT: RJSFSchema =
  // TO SHOW: original RJSF Object
  {
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
  }

const useServiceRequest = ({ serviceRequestId }: UseServiceRequestProps) => {
  const serviceRequest: ServiceRequest = {
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
