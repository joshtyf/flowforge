import { Pipeline } from "@/types/pipeline"
import apiClient from "./apiClient"
import { ServiceRequest } from "@/types/service-request"

export async function createPipeline(pipeline: Pipeline): Promise<Pipeline> {
  return apiClient.post("/pipeline", pipeline)
}

export async function getAllPipeline(): Promise<Pipeline[]> {
  return apiClient.get("/pipeline")
}

export async function getPipeline(pipelineId: string): Promise<Pipeline> {
  return apiClient.get(`/pipeline/${pipelineId}`)
}

export async function createServiceRequest(
  pipelineId: string,
  pipelineVersion?: string,
  remarks?: string
): Promise<ServiceRequest> {
  return apiClient.post("/service-request", {
    pipeline_id: pipelineId,
    pipeline_version: pipelineVersion,
    remarks: remarks,
  })
}

export async function getAllServiceRequest(): Promise<ServiceRequest[]> {
  return apiClient.post("/service-request")
}

export async function getServiceRequest(
  serviceRequestId: string
): Promise<ServiceRequest> {
  return apiClient.post(`/service-request/${serviceRequestId}`)
}

export async function approveServiceRequest(serviceRequestId: string) {
  return apiClient.post(`/service-request/${serviceRequestId}/approve`, {
    step_name: "step2",
  })
}
