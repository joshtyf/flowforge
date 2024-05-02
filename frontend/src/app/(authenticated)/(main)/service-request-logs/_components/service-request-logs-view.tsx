"use client"

import PipelineStepper from "@/components/layouts/pipeline-stepper"
import { Button } from "@/components/ui/button"
import HeaderAccessory from "@/components/ui/header-accessory"
import { ServiceRequestStep } from "@/types/service-request"
import { ChevronLeft } from "lucide-react"
import { useRouter } from "next/navigation"
import StepLogs from "./step-logs"

interface ServiceRequestLogsViewProps {
  serviceRequestId: string
  serviceRequestSteps?: ServiceRequestStep[]
  currentStep: string
}

export default function ServiceRequestLogsView({
  serviceRequestId,
  serviceRequestSteps,
  currentStep,
}: ServiceRequestLogsViewProps) {
  const router = useRouter()
  return (
    <>
      <div className="flex flex-col justify-start py-10">
        <HeaderAccessory />
        <div className="flex items-baseline space-x-2 pt-5">
          <Button
            size="icon"
            variant="ghost"
            onClick={() => {
              router.back()
            }}
          >
            <ChevronLeft />
          </Button>

          <p className="font-bold text-3xl">Service Request Logs</p>
        </div>
      </div>
      <div>
        <PipelineStepper steps={serviceRequestSteps} />
        <StepLogs
          serviceRequestId={serviceRequestId}
          currentStep={currentStep}
        />
      </div>
    </>
  )
}
