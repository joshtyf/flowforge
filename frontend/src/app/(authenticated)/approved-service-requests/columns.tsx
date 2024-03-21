"use client"

import { StatusBadge } from "@/components/layouts/status-badge"
import { formatDateString, formatTimeDifference } from "@/lib/utils"
import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"
import { ColumnDef } from "@tanstack/react-table"
import { ExternalLink } from "lucide-react"
import Link from "next/link"

export const approvedServiceRequestColumns: ColumnDef<ServiceRequest>[] = [
  {
    accessorKey: "status",
    header: "Status",
    cell: ({ row }) => {
      const status: ServiceRequestStatus = row.getValue("status")
      return <StatusBadge status={status} />
    },
  },
  {
    id: "service_name",
    header: "Service",
    cell: ({ row }) => {
      const serviceRequest: ServiceRequest = row.original
      return (
        <Link
          href={`/service-catalog/${serviceRequest.pipeline_id}`}
          className="hover:underline hover:text-blue-500 flex space-x-2"
        >
          <p>{serviceRequest.form_data.name}</p>
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
]
