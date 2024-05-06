import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { MoreHorizontal } from "lucide-react"
import { useState } from "react"
import { ApproveConfirmationDialog } from "./approve-confirmation-dialog"
import { RejectConfirmationDialog } from "./reject-confirmation-dialog"
import ServiceRequestDetailsDialog from "@/components/layouts/service-request-details-dialog"
import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"

interface AdminServiceRequestActionsProps {
  serviceRequest: ServiceRequest
  approveRequest: (serviceRequestId: string) => void
  rejectRequest: (serviceRequestId: string, remarks?: string) => void
}

export default function AdminServiceRequestActions({
  serviceRequest,
  approveRequest,
  rejectRequest,
}: AdminServiceRequestActionsProps) {
  const [openDetailsDialog, setOpenDetailsDialog] = useState(false)
  const [openApproveConfirmationDialog, setOpenApproveConfirmationDialog] =
    useState(false)
  const [openRejectConfirmationDialog, setOpenRejectConfirmationDialog] =
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
            setOpenDetailsDialog(true)
          }}
        >
          <Button variant="ghost">View Details</Button>
        </DropdownMenuItem>

        <DropdownMenuItem
          disabled={serviceRequest.status !== ServiceRequestStatus.PENDING}
        >
          <Button
            variant="ghost"
            className="text-green-700 hover:text-green-500"
            onClick={() => setOpenApproveConfirmationDialog(true)}
          >
            Approve
          </Button>
        </DropdownMenuItem>
        <DropdownMenuItem
          disabled={serviceRequest.status !== ServiceRequestStatus.PENDING}
        >
          {/* TODO: Add on click logic*/}
          <Button
            variant="ghost"
            className="text-red-700 hover:text-red-500"
            onClick={() => setOpenRejectConfirmationDialog(true)}
          >
            Reject
          </Button>
        </DropdownMenuItem>
      </DropdownMenuContent>
      <ServiceRequestDetailsDialog
        serviceRequest={serviceRequest}
        open={openDetailsDialog}
        setOpen={setOpenDetailsDialog}
      />
      <ApproveConfirmationDialog
        open={openApproveConfirmationDialog}
        setOpen={setOpenApproveConfirmationDialog}
        onApprove={() => approveRequest(serviceRequest.id)}
      />
      <RejectConfirmationDialog
        open={openRejectConfirmationDialog}
        setOpen={setOpenRejectConfirmationDialog}
        onReject={(remarks?: string) =>
          rejectRequest(serviceRequest.id, remarks)
        }
      />
    </DropdownMenu>
  )
}
