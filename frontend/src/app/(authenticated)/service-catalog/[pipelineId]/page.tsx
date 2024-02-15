"use client"

import React from "react"
import { useParams, useRouter } from "next/navigation"
import useServiceRequest from "./_hooks/use-service-request"
import { RegistryWidgetsType } from "@rjsf/utils"
import CustomCheckboxes from "@/components/form/custom-widgets/custom-checkboxes"
import CustomSelect from "@/components/form/custom-widgets/custom-select"
import ServiceRequestSkeletonView from "./_views/service-request-skeleton-view"
import ServiceRequestView from "./_views/service-request-view"

const widgets: RegistryWidgetsType = {
  CheckboxesWidget: CustomCheckboxes,
  SelectWidget: CustomSelect,
}

export default function ServiceRequestPage() {
  const { pipelineId } = useParams()
  const pipelineIdString = Array.isArray(pipelineId)
    ? pipelineId[0]
    : pipelineId
  const router = useRouter()
  const {
    service,
    rjsfSchema,
    uiSchema,
    handleSubmit,
    isLoadingForm,
    isSubmittingRequest,
  } = useServiceRequest({
    pipelineId: pipelineIdString,
  })

  return isLoadingForm ? (
    <ServiceRequestSkeletonView router={router} />
  ) : (
    <ServiceRequestView
      router={router}
      service={service}
      rjsfSchema={rjsfSchema}
      uiSchema={uiSchema}
      handleSubmit={handleSubmit}
      isSubmittingRequest={isSubmittingRequest}
    />
  )
}
