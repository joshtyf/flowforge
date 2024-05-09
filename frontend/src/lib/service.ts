import { Pipeline } from "@/types/pipeline"
import {
  ServiceRequest,
  ServiceRequestLogs,
  ServiceRequestSteps,
} from "@/types/service-request"
import { UserInfo } from "@/types/user-profile"
import apiClient from "./apiClient"

/* Pipeline */

export async function createPipeline(pipeline: Pipeline): Promise<Pipeline> {
  return apiClient.post("/pipeline", pipeline)
}

export async function getAllPipeline(): Promise<Pipeline[]> {
  return apiClient.get("/pipeline").then((res) => res.data)
}

export async function getPipeline(pipelineId: string): Promise<Pipeline> {
  return apiClient.get(`/pipeline/${pipelineId}`).then((res) => res.data)
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

export async function getServiceRequest(
  serviceRequestId: string
): Promise<ServiceRequest> {
  return apiClient
    .get(`/service_request/${serviceRequestId}`)
    .then((res) => res.data)
}

export async function approveServiceRequest(
  serviceRequestId: string,
  organizationId: string
) {
  return apiClient.post(`/service_request/${serviceRequestId}/approve`, {
    org_id: organizationId,
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

export async function getUserById(userId: string): Promise<UserInfo> {
  return apiClient.get(`/user/${userId}`).then((res) => res.data)
}
