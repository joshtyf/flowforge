import { Pipeline } from "@/types/pipeline"
import {
  ServiceRequest,
  ServiceRequestDTO,
  ServiceRequestLogs,
  ServiceRequestSteps,
} from "@/types/service-request"
import { UserInfo } from "@/types/user-profile"
import apiClient from "./apiClient"
import { UserMemberships } from "@/types/membership"

/* Pipeline */

export async function createPipeline(
  pipeline: Pipeline,
  organizationId: number
): Promise<Pipeline> {
  return apiClient.post("/pipeline", { ...pipeline, org_id: organizationId })
}

export async function getAllPipeline(
  organizationId: number
): Promise<Pipeline[]> {
  return apiClient
    .get("/pipeline", {
      params: {
        org_id: organizationId,
      },
    })
    .then((res) => res.data)
}

export async function getPipeline(
  pipelineId: string,
  organizationId: number
): Promise<Pipeline> {
  return apiClient
    .get(`/pipeline/${pipelineId}`, {
      params: {
        org_id: organizationId,
      },
    })
    .then((res) => res.data)
}

/* Service Request */

export async function createServiceRequest(
  organizationId: number,
  pipelineId: string,
  formData: object,
  pipelineVersion?: number,
  remarks?: string
): Promise<ServiceRequest> {
  return apiClient.post("/service_request", {
    pipeline_id: pipelineId,
    pipeline_version: pipelineVersion,
    form_data: formData,
    remarks: remarks,
    org_id: organizationId,
  })
}

export async function getAllServiceRequest(
  organizationId: number,
  page?: number,
  pageSize?: number
): Promise<{
  data: ServiceRequest[]
  metadata: {
    total_count: number
  }
}> {
  return apiClient
    .get("/service_request", {
      params: {
        org_id: organizationId,
        page: page ?? 1,
        page_size: pageSize ?? 10,
      },
    })
    .then((res) => res.data)
}

export async function getAllServiceRequestForAdmin(
  organizationId: number,
  page?: number,
  pageSize?: number
): Promise<{
  data: ServiceRequest[]
  metadata: {
    total_count: number
  }
}> {
  return apiClient
    .get("/service_request/admin", {
      params: {
        org_id: organizationId,
        page: page ?? 1,
        page_size: pageSize ?? 10,
      },
    })
    .then((res) => res.data)
}

export async function getServiceRequestDTO(
  serviceRequestId: string
): Promise<ServiceRequestDTO> {
  return apiClient
    .get(`/service_request/${serviceRequestId}`)
    .then((res) => res.data)
}

export async function approveServiceRequest(serviceRequestId: string) {
  return apiClient.put(`/service_request/${serviceRequestId}/approve`)
}

export async function startServiceRequest(serviceRequestId: string) {
  return apiClient.put(`/service_request/${serviceRequestId}/start`)
}

export async function cancelServiceRequest(serviceRequestId: string) {
  return apiClient.put(`/service_request/${serviceRequestId}/cancel`)
}

export async function rejectServiceRequest(
  serviceRequestId: string,
  remarks?: string
) {
  return apiClient.put(`/service_request/${serviceRequestId}/reject`, {
    remarks,
  })
}

export async function getServiceRequestSteps(
  serviceRequestId: string
): Promise<{
  service_request_id: string
  first_step_name: string
  pipeline_id: string
  pipeline_version: number
  steps: ServiceRequestSteps
}> {
  return apiClient
    .get(`/service_request/${serviceRequestId}/steps`)
    .then((res) => res.data)
}

export async function getServiceRequestLogs(
  serviceRequestId: string,
  stepName: string,
  offset?: number
): Promise<ServiceRequestLogs> {
  return apiClient
    .get(`/service_request/${serviceRequestId}/logs/${stepName}`, {
      params: offset
        ? {
            offset,
          }
        : {},
    })
    .then((res) => res.data)
}

/* Organization */

export async function getAllOrgsForUser() {
  return apiClient.get("/organization").then((res) => res.data)
}

export async function createOrg(orgName: string) {
  return apiClient
    .post("/organization", { name: orgName })
    .then((res) => res.data)
}

export async function updateOrgName(orgId: number, orgName: string) {
  return apiClient
    .patch("/organization", { org_id: orgId, name: orgName })
    .then((res) => res.data)
}

export async function getUserById(userId: string): Promise<UserInfo> {
  return apiClient.get(`/user/${userId}`).then((res) => res.data)
}

export async function getUserMemberships(): Promise<UserMemberships> {
  return apiClient.get(`/membership`).then((res) => res.data)
}

export async function login(): Promise<UserInfo> {
  return apiClient.get("login").then((res) => res.data)
}

export async function createUser(name: string): Promise<UserInfo> {
  return apiClient.post("/user", { name }).then((res) => res.data)
}
