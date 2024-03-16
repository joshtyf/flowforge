import { ServiceRequest } from "@/types/service-request"
import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useState } from "react"

interface ServiceRequestDetailsProps {
  serviceRequest: ServiceRequest
}

function ServiceRequestDetails({ serviceRequest }: ServiceRequestDetailsProps) {
  return (
    <div className="flex w-full h-full justify-center items-center">
      Service Request Details
    </div>
  )
}

interface ServiceRequestDetailsDialogProps {
  serviceRequest: ServiceRequest
  children: React.ReactNode
}

export default function ServiceRequestDetailsDialog({
  serviceRequest,
  children,
}: ServiceRequestDetailsDialogProps) {
  const [open, setOpen] = useState(false)

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>{`${serviceRequest.form_data.name} Details`}</DialogTitle>
          <DialogDescription>
            More details regarding Service Request
          </DialogDescription>
        </DialogHeader>
        <ServiceRequestDetails serviceRequest={serviceRequest} />
      </DialogContent>
    </Dialog>
  )
}
