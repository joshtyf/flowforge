"use client"

import { useParams } from "next/navigation"
import ServiceRequestLogsView from "../_components/service-request-logs-view"
import useServiceRequestSteps from "../_hooks/use-service-request-steps"
import ServiceRequestLogsSkeletonView from "../_components/service-request-logs-skeleton"

interface ServiceRequestLogsPageProps {}

export default function ServiceRequestLogsPage({}: ServiceRequestLogsPageProps) {
  const { serviceRequestId } = useParams()
  const serviceRequestIdString = Array.isArray(serviceRequestId)
    ? serviceRequestId[0]
    : serviceRequestId
  const { serviceRequestSteps, currentStep, isLoading } =
    useServiceRequestSteps({
      serviceRequestId: serviceRequestIdString,
    })
  return (
    <>
      {isLoading ? (
        <ServiceRequestLogsSkeletonView />
      ) : (
        <ServiceRequestLogsView
          currentStep={currentStep}
          serviceRequestSteps={serviceRequestSteps}
        />
      )}
    </>
  )
}
