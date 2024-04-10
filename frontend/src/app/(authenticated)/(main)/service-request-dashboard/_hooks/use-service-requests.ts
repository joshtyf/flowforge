import { toast } from "@/components/ui/use-toast"
import useOrganisationId from "@/hooks/use-organisation-id"
import { getAllServiceRequest } from "@/lib/service"
import {
  ServiceRequest,
  ServiceRequestStatus,
  ServiceRequestStep,
} from "@/types/service-request"
import { useQuery } from "@tanstack/react-query"

const PIPELINE_1_DUMMY_STEPS: ServiceRequestStep[] = [
  {
    name: "Approval",
    status: ServiceRequestStatus.NOT_STARTED,
  },
  {
    name: "Create EC2",
    status: ServiceRequestStatus.NOT_STARTED,
  },
]

const DUMMY_SERVICE_REQUESTS: ServiceRequest[] = [
  {
    id: "1",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.NOT_STARTED,
    created_by: "User 1",
    created_on: "2024-02-21T19:50:01",
    last_updated: "2024-02-21T19:50:01",
    remarks: "Remarks",
    form: {},
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
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_version: "0",
    pipeline_name: "Service 1",
    status: ServiceRequestStatus.PENDING,
    created_by: "User 1",
    created_on: "2024-02-21T18:50:01",
    last_updated: "2024-02-21T18:50:01",
    remarks: "Remarks",
    form: {},
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
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",

    status: ServiceRequestStatus.RUNNING,
    created_by: "User 1",
    created_on: "2024-02-21T17:00:00",
    last_updated: "2024-02-21T17:00:00",
    remarks: "Remarks",
    form: {},
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
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.SUCCESS,
    created_by: "User 1",
    created_on: "2024-02-21T00:00:00",
    last_updated: "2024-02-21T00:00:00",
    remarks: "",
    form: {},
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
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.REJECTED,
    created_by: "User 1",
    created_on: "2024-02-21T00:00:00",
    last_updated: "2024-02-21T00:00:00",
    remarks: "",
    form: {},
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
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.FAILURE,
    created_by: "User 1",
    created_on: "2024-02-20T00:00:00",
    last_updated: "2024-02-20T00:00:00",
    remarks: "",
    form: {},
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
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.CANCELLED,
    created_on: "2024-02-10T00:00:00",
    last_updated: "2024-02-10T00:00:00",
    remarks: "",
    form: {},
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
  // TODO: Integrate Service Request API by uncommenting below
  // const { organisationId } = useOrganisationId()
  // const { isLoading, data: serviceRequests } = useQuery({
  //   queryKey: ["user_service_requests"],
  //   queryFn: () =>
  //     getAllServiceRequest(organisationId).catch((err) => {
  //       console.log(err)
  //       toast({
  //         title: "Fetching Service Requests Error",
  //         description:
  //           "Failed to fetch Service Requests for user. Please try again later.",
  //         variant: "destructive",
  //       })
  //     }),
  // })
  return {
    serviceRequests: DUMMY_SERVICE_REQUESTS,
    // isLoading
  }
}

export default useServiceRequests