import { ServiceRequest } from "@/types/service-request"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"
import { Label } from "@/components/ui/label"
import { formatDateString, formatTimeDifference } from "@/lib/utils"
import Link from "next/link"
import Stepper from "./stepper"
import { ExternalLink } from "lucide-react"

interface ServiceRequestDetailsProps {
  serviceRequest: ServiceRequest
}

function ServiceRequestDetails({ serviceRequest }: ServiceRequestDetailsProps) {
  const {
    id: serviceRequestId,
    pipeline_version: pipelineVersion,
    created_by: createdBy = "",
    created_on: createdOn = "",
    last_updated: lastUpdated = "",
    remarks,
    steps,
  } = serviceRequest
  return (
    <div className="grid grid-cols-2 gap-5">
      <div className="col-span-2">
        <Label className="text-muted-foreground">Service Request Id</Label>
        <div>
          <TooltipProvider>
            <Tooltip delayDuration={300}>
              <TooltipTrigger>
                <Link
                  href={`/service-request-info/${serviceRequestId}`}
                  className="hover:underline hover:text-blue-500 flex space-x-1"
                >
                  <p>{serviceRequestId}</p>
                  <ExternalLink className="w-5 h-5" />
                </Link>
              </TooltipTrigger>
              <TooltipContent>
                <p>Go to Service Request Form Details</p>
              </TooltipContent>
            </Tooltip>
          </TooltipProvider>
        </div>
      </div>
      <div>
        <Label className="text-muted-foreground">Pipeline Version</Label>
        <p>{[pipelineVersion]}</p>
      </div>
      <div>
        <Label className="text-muted-foreground">Created By</Label>
        <p>{createdBy}</p>
      </div>
      {steps?.some((step) => step.name === "Approval") && (
        <div>
          <Label className="text-muted-foreground">Approved By</Label>
          <p>{"-"}</p>
        </div>
      )}
      <div>
        <Label className="text-muted-foreground">Created on</Label>
        <p>{formatDateString(new Date(createdOn))}</p>
      </div>
      <div>
        <Label className="text-muted-foreground">Last Updated</Label>
        <p>{formatTimeDifference(new Date(lastUpdated))}</p>
      </div>
      <div className="col-span-2">
        <Label className="text-muted-foreground">Remarks</Label>
        <p>{remarks}</p>
      </div>
      <div className="col-span-2">
        <Label className="text-muted-foreground">Steps</Label>
        <Stepper steps={steps} />
      </div>
    </div>
  )
}

interface ServiceRequestDetailsDialogProps {
  serviceRequest: ServiceRequest
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export default function ServiceRequestDetailsDialog({
  serviceRequest,
  open,
  setOpen,
}: ServiceRequestDetailsDialogProps) {
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{`${serviceRequest?.pipeline_name} Details`}</DialogTitle>
        </DialogHeader>
        <ServiceRequestDetails serviceRequest={serviceRequest} />
      </DialogContent>
    </Dialog>
  )
}
