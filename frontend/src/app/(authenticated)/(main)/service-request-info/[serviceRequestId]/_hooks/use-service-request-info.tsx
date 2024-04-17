import useServiceRequest from "@/hooks/use-service-request"
import { generateUiSchema } from "@/lib/rjsf-utils"
import { convertServiceRequestFormToRJSFSchema } from "@/lib/rjsf-utils"
import { FormFieldType, JsonFormComponents } from "@/types/json-form-components"
import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"
import { useMemo } from "react"

const DUMMY_PIPELINE_FORM: JsonFormComponents = {
  fields: [
    {
      name: "input",
      title: "Input",
      description: "",
      type: FormFieldType.INPUT,
      required: true,
      placeholder: "Enter text...",
      min_length: 1,
    },
    {
      name: "select",
      title: "Select",
      description: "",
      type: FormFieldType.SELECT,
      required: true,
      placeholder: "Select an option",
      options: ["Option 1", "Option 2", "Option 3"],
      default: "Option 1",
    },
    {
      name: "checkbox",
      title: "Checkbox",
      description: "",
      type: FormFieldType.CHECKBOXES,
      options: ["Option 1", "Option 2", "Option 3"],
    },
  ],
}

const DUMMY_SR_FORM_DATA = {
  input: "Input value",
  select: "Item 1",
  checkboxes: ["Item 1", "Item 2"],
}
const DUMMY_SERVICE_REQUEST: ServiceRequest = {
  id: "1",
  pipeline_id: "65d48c02d62a1281c4f4ba3e",
  pipeline_name: "Service 1",
  pipeline_version: "0",
  status: ServiceRequestStatus.NOT_STARTED,
  created_by: "User 1",
  created_on: "2024-02-21T19:50:01",
  last_updated: "2024-02-21T19:50:01",
  remarks: "Remarks",
  form: DUMMY_PIPELINE_FORM,
  form_data: DUMMY_SR_FORM_DATA,
  steps: [
    {
      name: "Approval",
      status: ServiceRequestStatus.NOT_STARTED,
    },
    {
      name: "Create EC2",
      status: ServiceRequestStatus.NOT_STARTED,
    },
    {
      name: "Create EC2",
      status: ServiceRequestStatus.NOT_STARTED,
    },
    {
      name: "Create EC2",
      status: ServiceRequestStatus.NOT_STARTED,
    },
    {
      name: "Create EC2",
      status: ServiceRequestStatus.NOT_STARTED,
    },
    {
      name: "Create EC2",
      status: ServiceRequestStatus.NOT_STARTED,
    },
  ],
}

interface UseServiceRequestInfoOptions {
  serviceRequestId: string
}

const useServiceRequestInfo = ({
  serviceRequestId,
}: UseServiceRequestInfoOptions) => {
  // TODO: Remove DUMMY_SERVICE_REQUEST when integrating with BE
  const { serviceRequest = DUMMY_SERVICE_REQUEST, isServiceRequestLoading } =
    useServiceRequest({
      serviceRequestId,
    })

  const uiSchema = useMemo(
    () => generateUiSchema(serviceRequest?.form),
    [serviceRequest]
  )
  const rjsfSchema = useMemo(
    () => convertServiceRequestFormToRJSFSchema(serviceRequest?.form),
    [serviceRequest]
  )

  return {
    pipelineName: serviceRequest?.pipeline_name,
    pipelineDescription: serviceRequest?.pipeline_description,
    formData: serviceRequest?.form_data,
    isServiceRequestLoading,
    uiSchema,
    rjsfSchema,
  }
}

export default useServiceRequestInfo
