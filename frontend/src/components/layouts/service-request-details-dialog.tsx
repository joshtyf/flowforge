import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Label } from "@/components/ui/label"
import { Skeleton } from "@/components/ui/skeleton"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"
import { toast } from "@/components/ui/use-toast"
import { getServiceRequestSteps, getUserById } from "@/lib/service"
import {
  createStepsFromObject,
  formatDateString,
  formatTimeDifference,
} from "@/lib/utils"
import { ServiceRequest, ServiceRequestSteps } from "@/types/service-request"
import { UserInfo } from "@/types/user-profile"
import { ExternalLink } from "lucide-react"
import Link from "next/link"
import { useEffect, useMemo, useState } from "react"
import PipelineStepper from "./pipeline-stepper"

interface ServiceRequestDetailsProps {
  serviceRequest: ServiceRequest
}

function ServiceRequestDetails({ serviceRequest }: ServiceRequestDetailsProps) {
  const [user, setUser] = useState<UserInfo>()
  useEffect(() => {
    getUserById(serviceRequest.user_id)
      .then((user) => setUser(user))
      .catch((err) => {
        console.error(err)
        toast({
          title: "Fetching Service Requests Error",
          description:
            "Failed to fetch Service Requests for user. Please try again later.",
          variant: "destructive",
        })
      })
  }, [serviceRequest.user_id])
  const {
    id: serviceRequestId,
    pipeline_version: pipelineVersion,
    created_by: createdBy = "",
    created_on: createdOn = "",
    last_updated: lastUpdated = "",
    remarks,
  } = serviceRequest

  const [steps, setSteps] = useState<ServiceRequestSteps>()
  const [firstStepName, setFirstStepName] = useState<string>()
  useEffect(() => {
    getServiceRequestSteps(serviceRequest.id)
      .then((resp) => {
        setSteps(resp.steps)
        setFirstStepName(resp.first_step_name)
      })
      .catch((err) => {
        console.error(err)
        toast({
          title: "Fetching Service Request Steps Error",
          description:
            "Failed to fetch Service Request Steps. Please try again later.",
          variant: "destructive",
        })
      })
  }, [serviceRequest])

  const stepsList = useMemo(
    () => createStepsFromObject(firstStepName, steps),
    [firstStepName, steps]
  )
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
        {user ? <p>{user.name}</p> : <Skeleton className="w-28 h-5" />}
      </div>
      {stepsList.some((step) => step.name === "Approval") && (
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
        <PipelineStepper steps={stepsList} />
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
      <DialogContent onOpenAutoFocus={(e) => e.preventDefault()}>
        <DialogHeader>
          <DialogTitle>{`${serviceRequest.pipeline_name} Details`}</DialogTitle>
        </DialogHeader>
        <ServiceRequestDetails serviceRequest={serviceRequest} />
      </DialogContent>
    </Dialog>
  )
}
