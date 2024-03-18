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
    pipeline_id: pipelineId,
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
        <Label className="text-muted-foreground">Pipeline Id</Label>
        <Link
          href={`/service-catalog/${pipelineId}`}
          className="hover:underline hover:text-blue-500 flex space-x-1"
        >
          <p>{pipelineId}</p>
          <ExternalLink className="w-5 h-5" />
        </Link>
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
  //   children: React.ReactNode
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
          <DialogTitle>{`${serviceRequest.form_data.name} Details`}</DialogTitle>
        </DialogHeader>
        <ServiceRequestDetails serviceRequest={serviceRequest} />
      </DialogContent>
    </Dialog>
  )
}
