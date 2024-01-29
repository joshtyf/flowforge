import { ServiceRequestWithSteps } from "@/types/service"
import { useState } from "react"

const BASIC_SERVICE_OBJECT: ServiceRequestWithSteps = {
  name: "",
  description: "",
  steps: {
    Approval: {
      name: "Approval",
      type: "approval",
      next: "",
      start: true,
    },
  },
  form: {},
}
// TODO: Create own custom form object schema to limit RJSF features. Limit to only these for now: Basic input, Select Dropdown, Checkboxes
const useCreateService = () => {
  const [serviceObject, setServiceObject] =
    useState<object>(BASIC_SERVICE_OBJECT)

  const handleSubmitObject = () => {
    // TODO: Replace with API call
    // TODO: Add loading for submit and re-direct to dashboard upon submission
    console.log("Submitting:", serviceObject)
  }

  // TODO: Add validation of JSON object on edit/add

  return { serviceObject, setServiceObject, handleSubmitObject }
}

export default useCreateService
