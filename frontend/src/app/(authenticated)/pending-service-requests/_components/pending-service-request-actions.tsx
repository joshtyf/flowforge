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

interface ApproveServiceRequestActionsProps {
  serviceRequestId: string
  approveRequest: (serviceRequestId: string) => void
  rejectRequest: (serviceRequestId: string, remarks?: string) => void
}

export default function ApproveServiceRequestActions({
  serviceRequestId,
  approveRequest,
  rejectRequest,
}: ApproveServiceRequestActionsProps) {
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
        <DropdownMenuItem>
          <Button
            variant="ghost"
            className="text-green-700 hover:text-green-500"
            onClick={() => setOpenApproveConfirmationDialog(true)}
          >
            Approve
          </Button>
        </DropdownMenuItem>
        <DropdownMenuItem>
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
      <ApproveConfirmationDialog
        open={openApproveConfirmationDialog}
        setOpen={setOpenApproveConfirmationDialog}
        onApprove={() => approveRequest(serviceRequestId)}
      />
      <RejectConfirmationDialog
        open={openRejectConfirmationDialog}
        setOpen={setOpenRejectConfirmationDialog}
        onReject={(remarks?: string) =>
          rejectRequest(serviceRequestId, remarks)
        }
      />
    </DropdownMenu>
  )
}
