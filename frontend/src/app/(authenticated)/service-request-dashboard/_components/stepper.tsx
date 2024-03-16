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
    <div className="w-full flex flex-col justify-center items-center">
      <div
        className={cn(
          `w-10 h-10 flex justify-center items-center ${statusBadgeVariant({
            status,
          })}`,
          "rounded-full"
        )}
      >
        <StatusIcon status={status} />
      </div>
      <p>{name}</p>
    </div>
  )
}

interface StepperProps {
  steps?: ServiceRequestStep[]
}

export default function Stepper({ steps }: StepperProps) {
  return (
    <ol className="flex flex-wrap">
      {steps?.map((step, index) => (
        <Step key={index} name={step.name} status={step.status} />
      ))}
    </ol>
  )
}
