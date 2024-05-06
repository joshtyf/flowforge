import { cn } from "@/lib/utils"
import { ServiceRequestStep } from "@/types/service-request"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"
import { StepStatus } from "@/types/pipeline"
import { StepStatusIcon, stepStatusBadgeVariant } from "./step-status-icon"

interface StepProps {
  name: string
  status: StepStatus
}

function Step({ name, status }: StepProps) {
  return (
    <li
      className={`w-[7.5rem] flex flex-col space-y-2 py-2 items-center
    [&:not(:last-child):after]:content-[''] 
    [&:not(:last-child):after]:relative 
    [&:not(:last-child):after]:top-[1rem] 
    [&:not(:last-child):after]:left-[4.7rem]
    [&:not(:last-child):after]:w-[6rem]
    [&:not(:last-child):after]:h-[2px] 
    ${status === StepStatus.STEP_COMPLETED ? "[&:not(:last-child):after]:bg-green-300" : "[&:not(:last-child):after]:bg-gray-300"}
    [&:not(:last-child):after]:-order-1`}
    >
      <TooltipProvider>
        <Tooltip delayDuration={300}>
          <TooltipTrigger disabled>
            <div
              className={cn(
                `w-[2rem] h-[2rem] flex justify-center items-center ${stepStatusBadgeVariant(
                  {
                    status,
                  }
                )}`,
                "rounded-full"
              )}
            >
              <StepStatusIcon status={status} />
            </div>
          </TooltipTrigger>
          <TooltipContent>
            <p>{status}</p>
          </TooltipContent>
        </Tooltip>
      </TooltipProvider>
      <p className="text-sm text-center w-full text-wrap truncate">{name}</p>
    </li>
  )
}

interface StepperProps {
  steps?: ServiceRequestStep[]
}

export default function PipelineStepper({ steps }: StepperProps) {
  return (
    <div className="flex justify-center">
      <ol className="flex flex-wrap py-2 space-x-[2rem]">
        {steps?.map((step, index) => (
          <Step key={index} name={step.name} status={step.status} />
        ))}
      </ol>
    </div>
  )
}
