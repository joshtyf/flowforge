"use client"

import { formatDateString, formatTimeDifference } from "@/lib/utils"
import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"
import { ColumnDef } from "@tanstack/react-table"
import Link from "next/link"
import StatusBadge from "../../../components/layouts/status-badge"
import ServiceRequestActions from "./_components/service-request-actions"

export const columns: ColumnDef<ServiceRequest>[] = [
  {
    accessorKey: "status",
    header: "Status",
    cell: ({ row }) => {
      const status: ServiceRequestStatus = row.getValue("status")
      return <StatusBadge status={status} />
    },
  },
  {
    accessorKey: "pipeline_id",
    header: "Pipeline",
    cell: ({ row }) => {
      const pipelineId: string = row.getValue("pipeline_id")
      return (
        <Link
          href={`/service-catalog/${pipelineId}`}
          className="hover:underline hover:text-blue-500"
        >
          {pipelineId}
        </Link>
      )
    },
  },
  {
    accessorKey: "created_on",
    header: "Created Date",
    cell: ({ row }) => {
      const dateIsoString: string = row.getValue("created_on")
      const dateObject = new Date(dateIsoString)
      return formatDateString(dateObject)
    },
  },
  {
    accessorKey: "last_updated",
    header: "Last Updated",
    cell: ({ row }) => {
      const dateIsoString: string = row.getValue("last_updated")
      const dateObject = new Date(dateIsoString)
      return formatTimeDifference(dateObject)
    },
  },
  {
    id: "actions",
    header: "Actions",
    cell: ({ row }) => {
      const serviceRequest: ServiceRequest = row.original
      return (
        <ServiceRequestActions
          serviceRequest={serviceRequest}
          onCancelRequest={(pipelineId: string) => {}}
        />
      )
    },
  },
]
