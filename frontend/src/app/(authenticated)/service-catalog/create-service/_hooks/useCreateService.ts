import { useState } from "react"

const BASIC_SERVICE_OBJECT = {
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

  form: {
    name: "",
  },
}

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
