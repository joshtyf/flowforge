import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"

const DUMMY_SERVICE_REQUESTS: ServiceRequest[] = [
  {
    id: "1",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_version: "0",
    status: ServiceRequestStatus.PENDING,
    created_on: "2024-02-20T19:50:01",
    last_updated: "2024-02-20T19:50:01",
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
    created_on: "2024-02-20T15:50:01",
    last_updated: "2024-02-20T15:50:01",
    remarks: "",
    form_data: {
      user_id: "2",
      user_name: "Test User",
      name: "Pipeline 2",
      description: "Pipeline 2 description",
    },
  },
  {
    id: "3",
    pipeline_id: "65d48c02d62a1281c4f4ba3e",
    pipeline_version: "0",
    status: ServiceRequestStatus.PENDING,
    created_on: "2024-02-19T00:00:00",
    last_updated: "2024-02-19T00:00:00",

    remarks: "",
    form_data: {
      user_id: "3",
      user_name: "Test User",
      name: "Pipeline 3",
      description: "Pipeline 3 description",
    },
  },
]

const useServiceRequests = () => {
  return { serviceRequests: DUMMY_SERVICE_REQUESTS }
}

export default useServiceRequests
