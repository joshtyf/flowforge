import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"

const DUMMY_SERVICE_REQUESTS: ServiceRequest[] = [
  {
    id: "1",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_version: "0",
    status: ServiceRequestStatus.NOT_STARTED,
    created_on: "2024-02-21T19:50:01",
    last_updated: "2024-02-21T19:50:01",
    remarks: "",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
  },
  {
    id: "2",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_version: "0",
    status: ServiceRequestStatus.PENDING,
    created_on: "2024-02-21T18:50:01",
    last_updated: "2024-02-21T18:50:01",
    remarks: "",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
  },
  {
    id: "3",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_version: "0",
    status: ServiceRequestStatus.RUNNING,
    created_on: "2024-02-21T17:00:00",
    last_updated: "2024-02-21T17:00:00",

    remarks: "",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
  },
  {
    id: "4",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_version: "0",
    status: ServiceRequestStatus.SUCCESS,
    created_on: "2024-02-21T00:00:00",
    last_updated: "2024-02-21T00:00:00",

    remarks: "",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
  },
  {
    id: "4",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_version: "0",
    status: ServiceRequestStatus.REJECTED,
    created_on: "2024-02-21T00:00:00",
    last_updated: "2024-02-21T00:00:00",

    remarks: "",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
  },
  {
    id: "5",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_version: "0",
    status: ServiceRequestStatus.FAILURE,
    created_on: "2024-02-20T00:00:00",
    last_updated: "2024-02-20T00:00:00",

    remarks: "",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
  },
  {
    id: "6",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_version: "0",
    status: ServiceRequestStatus.CANCELLED,
    created_on: "2024-02-10T00:00:00",
    last_updated: "2024-02-10T00:00:00",

    remarks: "",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
  },
]

const useServiceRequests = () => {
  return { serviceRequests: DUMMY_SERVICE_REQUESTS }
}

export default useServiceRequests
