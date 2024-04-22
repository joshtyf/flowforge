"use client"

import { formatDateString, formatTimeDifference } from "@/lib/utils"
import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"
import { ColumnDef } from "@tanstack/react-table"
import Link from "next/link"
import PendingServiceRequestActions from "./_components/pending-service-request-actions"
import { StatusBadge } from "@/components/layouts/status-badge"
import { ExternalLink } from "lucide-react"
import { DataTableColumnHeaderFilterableValue } from "@/components/data-table/data-table-column-header-filterable-value"

export const orgServiceRequestColumns: ColumnDef<ServiceRequest>[] = [
  {
    accessorKey: "status",
    header: ({ column }) => (
      <DataTableColumnHeaderFilterableValue
        column={column}
        title="Status"
        filterableOptions={Object.values(ServiceRequestStatus).map((value) => ({
          value,
          name: value,
        }))}
      />
    ),
    cell: ({ row }) => {
      const status: ServiceRequestStatus = row.getValue("status")
      return <StatusBadge status={status} />
    },
  },
  {
    id: "service_name",
    header: "Service Name",
    cell: ({ row }) => {
      const serviceRequest: ServiceRequest = row.original
      return (
        <Link
          href={`/service-catalog/${serviceRequest.pipeline_id}`}
          className="hover:underline hover:text-blue-500 flex space-x-2"
        >
          <p>{serviceRequest.pipeline_name}</p>
          <ExternalLink className="w-5 h-5" />
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
    accessorKey: "created_by",
    header: "Created By",
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
        <PendingServiceRequestActions
          serviceRequest={serviceRequest}
          approveRequest={(serviceRequestId: string) => {
            // TODO: Replace with actual approval action
            console.log("Approve service request for:", serviceRequestId)
          }}
          rejectRequest={(serviceRequestId: string, remarks?: string) => {
            // TODO: Replace with actual rejection action
            console.log("Reject service request for: ", serviceRequestId)
            console.log("Remarks: ", remarks)
          }}
        />
      )
    },
  },
]
