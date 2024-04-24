import { Pipeline } from "@/types/pipeline"
import apiClient from "./apiClient"
import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"
import { UserFromBackend } from "@/types/user-profile"

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
  organizationId: number
): Promise<ServiceRequest[]> {
  const res = await apiClient.get("/service_request", {
    params: { org_id: organizationId },
  })
  const serviceRequests = res.data
  return Promise.all(
    serviceRequests.map(
      async (serviceRequest: ServiceRequest): Promise<ServiceRequest> => {
        const pipelineId = serviceRequest.pipeline_id
        const pipeline = await getPipeline(pipelineId)
        const userId = (serviceRequest as unknown as { user_id: string })
          .user_id
        const user = await getUserById(userId)
        return {
          ...serviceRequest,
          created_by: user.name, // TODO: Should return user id instead. Render the proper user name in the UI separately
          pipeline_name: pipeline.pipeline_name, // TODO: Similar for pipeline
          pipeline_description: pipeline.pipeline_description,
        }
      }
    )
  )
  return apiClient
    .get("/service_request", { params: { org_id: organizationId } })
    .then((res) => res.data)
}

export async function getServiceRequest(
  serviceRequestId: string
): Promise<ServiceRequest> {
  return apiClient.post(`/service_request/${serviceRequestId}`)
}

export async function approveServiceRequest(
  serviceRequestId: string,
  organizationId: string
) {
  return apiClient.post(`/service_request/${serviceRequestId}/approve`, {
    org_id: organizationId,
  })
}

/* Organization */

export async function getAllOrgsForUser() {
  return apiClient.get("/organization").then((res) => res.data)
}

export async function getUserById(userId: string): Promise<UserInfo> {
  return apiClient.get(`/user/${userId}`).then((res) => res.data)
}
