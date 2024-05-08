"use client"

import PipelineStepper from "@/components/layouts/pipeline-stepper"
import { Button } from "@/components/ui/button"
import HeaderAccessory from "@/components/ui/header-accessory"
import { ServiceRequestStep } from "@/types/service-request"
import { ChevronLeft } from "lucide-react"
import { useRouter } from "next/navigation"
import StepLogs from "./step-logs"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { StepStatusIcon } from "@/components/layouts/step-status-icon"
import { useState } from "react"

interface ServiceRequestLogsViewProps {
  serviceRequestId: string
  serviceRequestSteps?: ServiceRequestStep[]
  currentStep: string
  handleCurrentStepChange: (step: string) => void
}

export default function ServiceRequestLogsView({
  serviceRequestId,
  serviceRequestSteps,
  currentStep,
  handleCurrentStepChange,
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
        {/* Cannot use default value prop in this case as currentStep is initially empty */}
        <Tabs value={currentStep} onValueChange={handleCurrentStepChange}>
          <TabsList className="grid w-full grid-cols-2 h-full">
            {serviceRequestSteps?.map((step) => (
              <TabsTrigger key={step.name} value={step.name}>
                <div className="flex items-center space-x-1">
                  <StepStatusIcon status={step.status} />
                  <p>{step.name}</p>
                </div>
              </TabsTrigger>
            ))}
          </TabsList>
          <p className="w-full text-center opacity-30 py-2">
            Click on a tab to view logs for the step
          </p>
          {serviceRequestSteps?.map((step) => (
            <TabsContent key={step.name} value={step.name}>
              <StepLogs
                serviceRequestId={serviceRequestId}
                stepName={step.name}
              />
            </TabsContent>
          ))}
        </Tabs>
      </div>
    </>
  )
}
