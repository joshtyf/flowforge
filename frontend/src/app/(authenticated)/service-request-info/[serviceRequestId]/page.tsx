"use client"

import React from "react"
import { useParams, useRouter } from "next/navigation"
import { RegistryWidgetsType } from "@rjsf/utils"
import CustomCheckboxes from "@/components/form/custom-widgets/custom-checkboxes"
import CustomSelect from "@/components/form/custom-widgets/custom-select"
import ServiceRequestSkeletonView from "@/components/layouts/service-request-skeleton-view"
import ServiceRequestView from "@/components/layouts/service-request-view"
import { ServiceRequest } from "@/types/service-request"
import useServiceRequestInfo from "./_hooks/use-service-request-info"

const widgets: RegistryWidgetsType = {
  CheckboxesWidget: CustomCheckboxes,
  SelectWidget: CustomSelect,
}

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
    <ServiceRequestSkeletonView
      router={router}
      returnRoute={"/service-catalog"}
    />
  ) : (
    <ServiceRequestView
      router={router}
      returnRoute={"/service-catalog"}
      formData={formData}
      pipelineName={pipelineName ?? ""}
      pipelineDescription={pipelineDescription}
      rjsfSchema={rjsfSchema}
      uiSchema={uiSchema}
    />
  )
}
