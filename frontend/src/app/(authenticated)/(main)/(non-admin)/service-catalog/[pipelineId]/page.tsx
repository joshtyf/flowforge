"use client"

import React from "react"
import { useParams, useRouter } from "next/navigation"
import useServiceRequest from "./_hooks/use-service-request-form"
import ServiceRequestSkeletonView from "@/components/layouts/service-request-skeleton-view"
import ServiceRequestView from "@/components/layouts/service-request-view"

export default function ServiceRequestPage() {
  const { pipelineId } = useParams()
  const pipelineIdString = Array.isArray(pipelineId)
    ? pipelineId[0]
    : pipelineId
  const router = useRouter()
  const {
    pipelineName,
    pipelineDescription,
    rjsfSchema,
    uiSchema,
    handleSubmit,
    isLoadingForm,
    isSubmittingRequest,
  } = useServiceRequest({
    pipelineId: pipelineIdString,
  })

  return isLoadingForm ? (
    <ServiceRequestSkeletonView
      router={router}
      returnRoute={"/service-catalog"}
    />
  ) : (
    <ServiceRequestView
      router={router}
      pipelineName={pipelineName ?? ""}
      pipelineDescription={pipelineDescription}
      rjsfSchema={rjsfSchema}
      uiSchema={uiSchema}
      handleSubmit={handleSubmit}
      isSubmittingRequest={isSubmittingRequest}
    />
  )
}
