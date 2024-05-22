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

interface ServiceRequestActionsProps {
  serviceRequest: ServiceRequest
  onCancelRequest: (pipelineId: string) => void // TODO: I don't think this is correct
  onStartRequest: (serviceRequestId: string) => void
}

export default function ServiceRequestActions({
  serviceRequest,
  onCancelRequest,
  onStartRequest,
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
        {serviceRequest.status === ServiceRequestStatus.NOT_STARTED ? (
          <DropdownMenuItem onClick={() => onStartRequest(serviceRequest.id)}>
            <Button variant="ghost">Start Request</Button>
          </DropdownMenuItem>
        ) : (
          <DropdownMenuItem>
            <Link href={`/service-request-logs/${serviceRequest.id}`}>
              <Button variant="ghost">View Logs</Button>
            </Link>
          </DropdownMenuItem>
        )}
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
