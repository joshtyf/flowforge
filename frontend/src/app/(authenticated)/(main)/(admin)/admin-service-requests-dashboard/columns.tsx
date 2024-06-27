"use client"

import { formatDateString, formatTimeDifference } from "@/lib/utils"
import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"
import { ColumnDef } from "@tanstack/react-table"
import Link from "next/link"
import AdminServiceRequestActions from "./_components/admin-service-request-actions"
import { ServiceRequestStatusBadge } from "@/components/layouts/service-request-status-badge"
import { ExternalLink } from "lucide-react"
import { DataTableColumnHeaderFilterableValue } from "@/components/data-table/data-table-column-header-filterable-value"
import CreatedByInfo from "@/components/layouts/created-by-info"
import { approveServiceRequest, rejectServiceRequest } from "@/lib/service"
import { toast } from "@/components/ui/use-toast"

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
      return <ServiceRequestStatusBadge status={status} />
    },
  },
  {
    accessorKey: "id",
    header: "Service Request ID",
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
    header: "Created By",
    cell: ({ row }) => {
      const serviceRequest: ServiceRequest = row.original
      const userId: string = serviceRequest.user_id
      return <CreatedByInfo userId={userId} />
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
        <AdminServiceRequestActions
          serviceRequest={serviceRequest}
          approveRequest={(serviceRequestId: string) => {
            approveServiceRequest(serviceRequestId)
              .then(() => {
                toast({
                  title: "Approve Service Request Successful",
                  description:
                    "Please check the dashboard for the updated status of the Service Request.",
                  variant: "success",
                })
              })
              .catch((error) => {
                toast({
                  title: "Approve Service Request Error",
                  description:
                    "Failed to approve Service Request. Please try again later.",
                  variant: "destructive",
                })
                console.error(error)
              })
          }}
          rejectRequest={(serviceRequestId: string, remarks?: string) => {
            rejectServiceRequest(serviceRequestId, remarks)
              .then(() => {
                toast({
                  title: "Reject Service Request Successful",
                  description:
                    "Please check the dashboard for the updated status of the Service Request.",
                  variant: "success",
                })
              })
              .catch((error) => {
                toast({
                  title: "Reject Service Request Error",
                  description:
                    "Failed to reject Service Request. Please try again later.",
                  variant: "destructive",
                })
                console.error(error)
              })
          }}
        />
      )
    },
  },
]
