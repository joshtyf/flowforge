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
    console.log("Submitting:", serviceObject)
  }

  return { serviceObject, setServiceObject, handleSubmitObject }
}

export default useCreateService
