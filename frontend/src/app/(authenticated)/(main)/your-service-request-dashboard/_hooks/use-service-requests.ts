import { toast } from "@/components/ui/use-toast"
import useOrganizationId from "@/hooks/use-organization-id"
import { getAllServiceRequest, getPipeline, getUserById } from "@/lib/service"
import { FormFieldType, JsonFormComponents } from "@/types/json-form-components"
import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"
import { useQuery } from "@tanstack/react-query"

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
    form: DUMMY_PIPELINE_FORM,
    form_data: {},
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
    form: DUMMY_PIPELINE_FORM,
    form_data: {},
    steps: [
      {
        name: "Approval",
        status: ServiceRequestStatus.PENDING,
      },
      {
        name: "Create EC2",
        status: ServiceRequestStatus.NOT_STARTED,
      },
    ],
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
    form: DUMMY_PIPELINE_FORM,
    form_data: {},
    steps: [
      {
        name: "Approval",
        status: ServiceRequestStatus.COMPLETED,
      },
      {
        name: "Create EC2",
        status: ServiceRequestStatus.RUNNING,
      },
    ],
  },
  {
    id: "4",
    user_id: "123456", // DUMMY
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.SUCCESS,
    created_by: "User 1",
    created_on: "2024-02-21T00:00:00",
    last_updated: "2024-02-21T00:00:00",
    remarks: "",
    form: DUMMY_PIPELINE_FORM,
    form_data: {},
    steps: [
      {
        name: "Approval",
        status: ServiceRequestStatus.COMPLETED,
      },
      {
        name: "Create EC2",
        status: ServiceRequestStatus.COMPLETED,
      },
    ],
  },
  {
    id: "4",
    user_id: "123456", // DUMMY
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.REJECTED,
    created_by: "User 1",
    created_on: "2024-02-21T00:00:00",
    last_updated: "2024-02-21T00:00:00",
    remarks: "",
    form: DUMMY_PIPELINE_FORM,
    form_data: {},
    steps: [
      {
        name: "Approval",
        status: ServiceRequestStatus.REJECTED,
      },
      {
        name: "Create EC2",
        status: ServiceRequestStatus.NOT_STARTED,
      },
    ],
  },
  {
    id: "5",
    user_id: "123456", // DUMMY
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.FAILURE,
    created_by: "User 1",
    created_on: "2024-02-20T00:00:00",
    last_updated: "2024-02-20T00:00:00",
    remarks: "",
    form: DUMMY_PIPELINE_FORM,
    form_data: {},
    steps: [
      {
        name: "Approval",
        status: ServiceRequestStatus.COMPLETED,
      },
      {
        name: "Create EC2",
        status: ServiceRequestStatus.FAILURE,
      },
    ],
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
    form: DUMMY_PIPELINE_FORM,
    form_data: {},
    steps: [
      {
        name: "Approval",
        status: ServiceRequestStatus.CANCELLED,
      },
      {
        name: "Create EC2",
        status: ServiceRequestStatus.NOT_STARTED,
      },
    ],
  },
]

const useServiceRequests = () => {
  const { organizationId } = useOrganizationId()
  const { isLoading, data: serviceRequests } = useQuery({
    queryKey: ["user_service_requests"],
    queryFn: async () => {
      let serviceRequests
      try {
        serviceRequests = await getAllServiceRequest(organizationId)
        // Fetch pipeline details for each service request
        serviceRequests = await Promise.all(
          serviceRequests.map(async (serviceRequest) => {
            const pipelineId = serviceRequest.pipeline_id
            const pipeline = await getPipeline(pipelineId)
            return {
              ...serviceRequest,
              pipeline_name: pipeline.pipeline_name,
              pipeline_description: pipeline.pipeline_description,
            }
          })
        )
        // Fetch user details for each service request
        serviceRequests = await Promise.all(
          serviceRequests.map(async (serviceRequest) => {
            const userId = serviceRequest.user_id
            const user = await getUserById(userId)
            return {
              ...serviceRequest,
              created_by: user.name,
            }
          })
        )
        return serviceRequests
      } catch (err) {
        console.log(err)
        toast({
          title: "Fetching Service Requests Error",
          description:
            "Failed to fetch Service Requests for user. Please try again later.",
          variant: "destructive",
        })
      }
    },
  })

  return {
    serviceRequests: serviceRequests,
    isLoading,
  }
}

export default useServiceRequests
