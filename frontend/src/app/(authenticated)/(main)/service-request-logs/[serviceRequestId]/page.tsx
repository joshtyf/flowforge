"use client"

import { useParams } from "next/navigation"
import ServiceRequestLogsView from "../_components/service-request-logs-view"
import useServiceRequestLogs from "../_hooks/use-service-request-logs"
import ServiceRequestLogsSkeletonView from "../_components/service-request-logs-skeleton"

interface ServiceRequestLogsPageProps {}

export default function ServiceRequestLogsPage({}: ServiceRequestLogsPageProps) {
  const { serviceRequestId } = useParams()
  const serviceRequestIdString = Array.isArray(serviceRequestId)
    ? serviceRequestId[0]
    : serviceRequestId
  const { serviceRequestSteps, isLoading } = useServiceRequestLogs({
    serviceRequestId: serviceRequestIdString,
  })
  return (
    <>
      {isLoading ? (
        <ServiceRequestLogsSkeletonView />
      ) : (
        <ServiceRequestLogsView serviceRequestSteps={serviceRequestSteps} />
      )}
    </>
  )
}
