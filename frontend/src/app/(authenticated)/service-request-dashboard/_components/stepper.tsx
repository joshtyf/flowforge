import {
  StatusIcon,
  statusBadgeVariant,
} from "@/components/layouts/status-badge"
import { cn } from "@/lib/utils"
import {
  ServiceRequestStatus,
  ServiceRequestStep,
} from "@/types/service-request"

interface StepProps {
  name: string
  status: ServiceRequestStatus
}

const STEP_SIZE = 3
const STEP_SPACING = 0.5

function Step({ name, status }: StepProps) {
  return (
    <li
      className={`flex flex-col space-y-2 items-center
    [&:not(:last-child):after]:content-[''] 
    [&:not(:last-child):after]:relative 
    [&:not(:last-child):after]:top-[1rem] 
    [&:not(:last-child):after]:left-[3.2rem]
    [&:not(:last-child):after]:w-[3rem]
    [&:not(:last-child):after]:h-[2px] 
    ${status === ServiceRequestStatus.COMPLETED ? "[&:not(:last-child):after]:bg-green-300" : "[&:not(:last-child):after]:bg-gray-300"}
    [&:not(:last-child):after]:-order-1`}
    >
      {/* TODO: Add tooltip to show status*/}
      <div
        className={cn(
          `w-[2rem] h-[2rem] flex justify-center items-center ${statusBadgeVariant(
            {
              status,
            }
          )}`,
          "rounded-full"
        )}
      >
        <StatusIcon status={status} />
      </div>
      <p className="text-sm">{name}</p>
    </li>
  )
}

interface StepperProps {
  steps?: ServiceRequestStep[]
}

export default function Stepper({ steps }: StepperProps) {
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
