"use client"

import React from "react"
import { useParams, useRouter } from "next/navigation"
import ServiceRequestSkeletonView from "@/components/layouts/service-request-skeleton-view"
import ServiceRequestView from "@/components/layouts/service-request-view"
import useServiceRequestInfo from "./_hooks/use-service-request-info"

export default function ServiceRequestInfoPage() {
  const { serviceRequestId } = useParams()
  const serviceRequestIdString = Array.isArray(serviceRequestId)
    ? serviceRequestId[0]
    : serviceRequestId
  const router = useRouter()
  const {
    pipelineName,
    pipelineDescription,
    formData,
    rjsfSchema,
    uiSchema,
    isServiceRequestLoading,
  } = useServiceRequestInfo({
    serviceRequestId: serviceRequestIdString,
  })

  return isServiceRequestLoading ? (
    <ServiceRequestSkeletonView router={router} />
  ) : (
    <ServiceRequestView
      router={router}
      formData={formData}
      pipelineName={pipelineName ?? ""}
      pipelineDescription={pipelineDescription}
      rjsfSchema={rjsfSchema}
      uiSchema={uiSchema}
      viewOnly
    />
  )
}
