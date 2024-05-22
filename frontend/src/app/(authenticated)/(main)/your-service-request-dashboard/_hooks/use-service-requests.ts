import { toast } from "@/components/ui/use-toast"
import useOrganizationId from "@/hooks/use-organization-id"
import { getAllServiceRequest } from "@/lib/service"
import { FormFieldType, JsonFormComponents } from "@/types/json-form-components"
import { StepStatus } from "@/types/pipeline"
import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"
import { useQuery } from "@tanstack/react-query"
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

const DUMMY_SERVICE_REQUESTS: ServiceRequest[] = [
  {
    id: "1",
    user_id: "123456", // DUMMY
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.NOT_STARTED,
    created_by: "User 1",
    created_on: "2024-02-21T19:50:01",
    last_updated: "2024-02-21T19:50:01",
    remarks: "Remarks",
    pipeline: {
      form: DUMMY_PIPELINE_FORM,
    },
    first_step_name: "Approval",
    form_data: {},
    steps: {
      Approval: {
        name: "Approval",
        status: StepStatus.STEP_NOT_STARTED,
        next_step_name: "Create EC2",
      },

      "Create EC2": {
        name: "Create EC2",
        status: StepStatus.STEP_NOT_STARTED,
        next_step_name: "",
      },
    },
  },
  {
    id: "2",
    user_id: "123456", // DUMMY
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_version: "0",
    pipeline_name: "Service 1",
    status: ServiceRequestStatus.PENDING,
    created_by: "User 1",
    created_on: "2024-02-21T18:50:01",
    last_updated: "2024-02-21T18:50:01",
    remarks: "Remarks",
    pipeline: {
      form: DUMMY_PIPELINE_FORM,
    },
    first_step_name: "Approval",
    form_data: {},
    steps: {
      Approval: {
        name: "Approval",
        status: StepStatus.STEP_RUNNING,
        next_step_name: "Create EC2",
      },

      "Create EC2": {
        name: "Create EC2",
        status: StepStatus.STEP_NOT_STARTED,
        next_step_name: "",
      },
    },
  },
  {
    id: "3",
    user_id: "123456", // DUMMY
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",

    status: ServiceRequestStatus.RUNNING,
    created_by: "User 1",
    created_on: "2024-02-21T17:00:00",
    last_updated: "2024-02-21T17:00:00",
    remarks: "Remarks",
    pipeline: {
      form: DUMMY_PIPELINE_FORM,
    },
    first_step_name: "Approval",
    form_data: {},
    steps: {
      Approval: {
        name: "Approval",
        status: StepStatus.STEP_COMPLETED,
        next_step_name: "Create EC2",
      },

      "Create EC2": {
        name: "Create EC2",
        status: StepStatus.STEP_RUNNING,
        next_step_name: "",
      },
    },
  },
  {
    id: "4",
    user_id: "123456", // DUMMY
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.COMPLETED,
    created_by: "User 1",
    created_on: "2024-02-21T00:00:00",
    last_updated: "2024-02-21T00:00:00",
    remarks: "",
    pipeline: {
      form: DUMMY_PIPELINE_FORM,
    },
    first_step_name: "Approval",
    form_data: {},
    steps: {
      Approval: {
        name: "Approval",
        status: StepStatus.STEP_COMPLETED,
        next_step_name: "Create EC2",
      },

      "Create EC2": {
        name: "Create EC2",
        status: StepStatus.STEP_COMPLETED,
        next_step_name: "",
      },
    },
  },
  {
    id: "4",
    user_id: "123456", // DUMMY
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.FAILED,
    created_by: "User 1",
    created_on: "2024-02-21T00:00:00",
    last_updated: "2024-02-21T00:00:00",
    remarks: "",
    pipeline: {
      form: DUMMY_PIPELINE_FORM,
    },
    first_step_name: "Approval",
    form_data: {},
    steps: {
      Approval: {
        name: "Approval",
        status: StepStatus.STEP_FAILURE,
        next_step_name: "Create EC2",
      },

      "Create EC2": {
        name: "Create EC2",
        status: StepStatus.STEP_NOT_STARTED,
        next_step_name: "",
      },
    },
  },
  {
    id: "5",
    user_id: "123456", // DUMMY
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.FAILED,
    created_by: "User 1",
    created_on: "2024-02-20T00:00:00",
    last_updated: "2024-02-20T00:00:00",
    remarks: "",
    pipeline: {
      form: DUMMY_PIPELINE_FORM,
    },
    first_step_name: "Approval",
    form_data: {},
    steps: {
      Approval: {
        name: "Approval",
        status: StepStatus.STEP_COMPLETED,
        next_step_name: "Create EC2",
      },

      "Create EC2": {
        name: "Create EC2",
        status: StepStatus.STEP_FAILURE,
        next_step_name: "",
      },
    },
  },
  {
    id: "6",
    user_id: "123456", // DUMMY
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.CANCELLED,
    created_on: "2024-02-10T00:00:00",
    last_updated: "2024-02-10T00:00:00",
    remarks: "",
    pipeline: {
      form: DUMMY_PIPELINE_FORM,
    },
    first_step_name: "Approval",
    form_data: {},
    steps: {
      Approval: {
        name: "Approval",
        status: StepStatus.STEP_FAILURE,
        next_step_name: "Create EC2",
      },

      "Create EC2": {
        name: "Create EC2",
        status: StepStatus.STEP_RUNNING,
        next_step_name: "",
      },
    },
  },
]
interface UseServiceRequestProps {
  page: number
  pageSize: number
}

const useServiceRequests = ({ page, pageSize }: UseServiceRequestProps) => {
  const { organizationId } = useOrganizationId()
  const { isLoading, data } = useQuery({
    queryKey: ["user_service_requests", page, pageSize],
    queryFn: () => {
      return getAllServiceRequest(organizationId, page, pageSize).catch(
        (err) => {
          console.error(err)
          toast({
            title: "Fetching Service Requests Error",
            description:
              "Failed to fetch Service Requests for user. Please try again later.",
            variant: "destructive",
          })
        }
      )
    },
    refetchInterval: 2000,
  })

  const noOfPages = useMemo(
    () =>
      data?.metadata.total_count
        ? Math.ceil(data?.metadata.total_count / pageSize)
        : undefined,
    [data, pageSize]
  )

  return {
    response: data,
    isLoading,
    noOfPages,
  }
}

export default useServiceRequests
