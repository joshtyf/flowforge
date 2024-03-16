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

function Step({ name, status }: StepProps) {
  return (
    <li
      className={`flex flex-col flex-1 items-center space-y-2 
    [&:not(:last-child):after]:content-[''] 
    [&:not(:last-child):after]:relative 
    [&:not(:last-child):after]:top-[1.5rem] 
    [&:not(:last-child):after]:left-[50%]
    [&:not(:last-child):after]:w-[73%]
    [&:not(:last-child):after]:h-[2px] 
    ${status === ServiceRequestStatus.COMPLETED ? "[&:not(:last-child):after]:bg-green-300" : "[&:not(:last-child):after]:bg-gray-300"}
    [&:not(:last-child):after]:-order-1`}
    >
      <div
        className={cn(
          `w-[48px] h-[48px] flex justify-center items-center ${statusBadgeVariant(
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
    <ol className="flex flex-wrap py-2">
      {steps?.map((step, index) => (
        <Step key={index} name={step.name} status={step.status} />
      ))}
    </ol>
  )
}
