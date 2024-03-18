import {
  ServiceRequest,
  ServiceRequestStatus,
  ServiceRequestStep,
} from "@/types/service-request"

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
    pipeline_version: "0",
    status: ServiceRequestStatus.NOT_STARTED,
    created_by: "User 1",
    created_on: "2024-02-21T19:50:01",
    last_updated: "2024-02-21T19:50:01",
    remarks: "Remarks",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
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
    status: ServiceRequestStatus.PENDING,
    created_by: "User 1",
    created_on: "2024-02-21T18:50:01",
    last_updated: "2024-02-21T18:50:01",
    remarks: "Remarks",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
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
    pipeline_version: "0",
    status: ServiceRequestStatus.RUNNING,
    created_by: "User 1",
    created_on: "2024-02-21T17:00:00",
    last_updated: "2024-02-21T17:00:00",
    remarks: "Remarks",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
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
    pipeline_version: "0",
    status: ServiceRequestStatus.SUCCESS,
    created_by: "User 1",
    created_on: "2024-02-21T00:00:00",
    last_updated: "2024-02-21T00:00:00",
    remarks: "",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
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
    pipeline_version: "0",
    status: ServiceRequestStatus.REJECTED,
    created_by: "User 1",
    created_on: "2024-02-21T00:00:00",
    last_updated: "2024-02-21T00:00:00",
    remarks: "",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
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
    pipeline_version: "0",
    status: ServiceRequestStatus.FAILURE,
    created_by: "User 1",
    created_on: "2024-02-20T00:00:00",
    last_updated: "2024-02-20T00:00:00",
    remarks: "",
    form_data: {
      user_id: "1",
      user_name: "Test User",
      name: "Pipeline 1",
      description: "Pipeline 1 description",
    },
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
  return { serviceRequests: DUMMY_SERVICE_REQUESTS }
}

export default useServiceRequests
