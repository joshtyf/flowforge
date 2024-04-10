import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { MoreHorizontal } from "lucide-react"
import { useState } from "react"
import ServiceRequestDetailsDialog from "@/components/layouts/service-request-details-dialog"
import { ServiceRequest } from "@/types/service-request"

interface RejectedServiceRequestActionsProps {
  serviceRequest: ServiceRequest
}

export default function RejectedServiceRequestActions({
  serviceRequest,
}: RejectedServiceRequestActionsProps) {
  const [openDetailsDialog, setOpenDetailsDialog] = useState(false)
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
      </DropdownMenuContent>
      <ServiceRequestDetailsDialog
        serviceRequest={serviceRequest}
        open={openDetailsDialog}
        setOpen={setOpenDetailsDialog}
      />
    </DropdownMenu>
  )
}
