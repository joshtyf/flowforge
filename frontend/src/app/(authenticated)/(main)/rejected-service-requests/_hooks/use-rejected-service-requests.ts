import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"

const DUMMY_SERVICE_REQUESTS: ServiceRequest[] = [
  {
    id: "1",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.REJECTED,
    created_on: "2024-02-21T19:50:01",
    created_by: "User 1",
    last_updated: "2024-02-21T19:50:01",
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
    id: "2",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.REJECTED,
    created_on: "2024-02-21T18:50:01",
    created_by: "User 2",
    last_updated: "2024-02-21T18:50:01",
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
    id: "3",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.REJECTED,
    created_on: "2024-02-21T17:00:00",
    created_by: "User 3",
    last_updated: "2024-02-21T17:00:00",
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
    id: "4",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_name: "Service 1",
    pipeline_version: "0",
    status: ServiceRequestStatus.REJECTED,
    created_on: "2024-02-21T00:00:00",
    created_by: "User 4",
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
]

const useRejectedServiceRequests = () => {
  return { serviceRequests: DUMMY_SERVICE_REQUESTS }
}

export default useRejectedServiceRequests
