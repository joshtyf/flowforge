import { ServiceRequestWithSteps } from "@/types/service"
import { useState } from "react"

const BASIC_SERVICE_OBJECT: ServiceRequestWithSteps = {
  name: "",
  description: "",
  form: {},
  steps: {
    Approval: {
      name: "Approval",
      type: "approval",
      next: "",
      start: true,
    },
  },
}
const useCreateService = () => {
  const [serviceObject, setServiceObject] =
    useState<object>(BASIC_SERVICE_OBJECT)

  const handleSubmitObject = () => {
    // TODO: Replace with API call
    console.log("Submitting:", serviceObject)
  }

  // TODO: Add validation of JSON object on edit/submit
  return { serviceObject, setServiceObject, handleSubmitObject }
}

export default useCreateService
