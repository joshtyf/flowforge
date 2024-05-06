import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { ServiceRequest, ServiceRequestStatus } from "@/types/service-request"
import { MoreHorizontal } from "lucide-react"
import ServiceRequestDetailsDialog from "@/components/layouts/service-request-details-dialog"
import { useState } from "react"
import Link from "next/link"

interface ServiceRequestActionsProps {
  serviceRequest: ServiceRequest
  onCancelRequest: (pipelineId: string) => void
}

export default function ServiceRequestActions({
  serviceRequest,
  onCancelRequest,
}: ServiceRequestActionsProps) {
  const [openDialog, setOpenDialog] = useState(false)

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
        <DropdownMenuItem>
          {/* TODO: Add on click logic*/}
          <Button
            variant="ghost"
            onClick={() => onCancelRequest(serviceRequest.pipeline_id)}
          >
            Cancel Request
          </Button>
        </DropdownMenuItem>
      </DropdownMenuContent>
      <ServiceRequestDetailsDialog
        open={openDialog}
        setOpen={setOpenDialog}
        serviceRequest={serviceRequest}
      />
    </DropdownMenu>
  )
}
