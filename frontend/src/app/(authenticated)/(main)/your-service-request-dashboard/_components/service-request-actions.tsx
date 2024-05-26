import ServiceRequestDetailsDialog from "@/components/layouts/service-request-details-dialog"
import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"
import { MoreHorizontal } from "lucide-react"
import Link from "next/link"
import { useState } from "react"
import { CancelConfirmationDialog } from "./cancel-confirmation-dialog"

interface ServiceRequestActionsProps {
  serviceRequest: ServiceRequest
  onCancelRequest: (serviceRequestId: string) => void
  onStartRequest: (serviceRequestId: string) => void
}

export default function ServiceRequestActions({
  serviceRequest,
  onCancelRequest,
  onStartRequest,
}: ServiceRequestActionsProps) {
  const [openDialog, setOpenDialog] = useState(false)
  const [openCancelConfirmationDialog, setOpenCancelConfirmationDialog] =
    useState(false)
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" className="h-8 w-8 p-0">
          <MoreHorizontal className="h-4 w-4" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="start">
        <DropdownMenuItem
          onClick={() => {
            setOpenDialog(true)
          }}
        >
          <Button variant="ghost">View Details</Button>
        </DropdownMenuItem>
        <DropdownMenuItem
          disabled={serviceRequest.status === ServiceRequestStatus.NOT_STARTED}
        >
          <Link href={`/service-request-logs/${serviceRequest.id}`}>
            <Button variant="ghost">View Logs</Button>
          </Link>
        </DropdownMenuItem>
        <DropdownMenuItem
          disabled={serviceRequest.status !== ServiceRequestStatus.NOT_STARTED}
          onClick={() => onStartRequest(serviceRequest.id)}
        >
          <Button variant="ghost">Start Request</Button>
        </DropdownMenuItem>
        <DropdownMenuItem
          // Only allow cancelling if request is running or pending
          disabled={
            serviceRequest.status !== ServiceRequestStatus.RUNNING &&
            serviceRequest.status !== ServiceRequestStatus.PENDING
          }
          onClick={() => setOpenCancelConfirmationDialog(true)}
        >
          <Button variant="ghost">Cancel Request</Button>
        </DropdownMenuItem>
      </DropdownMenuContent>
      <ServiceRequestDetailsDialog
        open={openDialog}
        setOpen={setOpenDialog}
        serviceRequest={serviceRequest}
      />
      <CancelConfirmationDialog
        open={openCancelConfirmationDialog}
        setOpen={setOpenCancelConfirmationDialog}
        onCancel={() => onCancelRequest(serviceRequest.id)}
      />
    </DropdownMenu>
  )
}
