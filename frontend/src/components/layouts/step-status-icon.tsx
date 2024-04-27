import { cn } from "@/lib/utils"
import { StepStatus } from "@/types/pipeline"
import { cva } from "class-variance-authority"
import {
  AlertCircle,
  CheckCircle2,
  CircleDotDashed,
  Moon,
  XCircle,
} from "lucide-react"

const stepStatusIconVariant = cva("", {
  variants: {
    status: {
      [StepStatus.STEP_NOT_STARTED]: "text-slate-500",
      [StepStatus.STEP_RUNNING]: "text-blue-500",
      [StepStatus.STEP_COMPLETED]: "text-green-500",
      [StepStatus.STEP_FAILURE]: "text-red-500",
      [StepStatus.STEP_CANCELLED]: "text-orange-500",
    },
  },
  defaultVariants: {
    status: StepStatus.STEP_NOT_STARTED,
  },
})

export const stepStatusBadgeVariant = cva(
  "rounded-lg border text-sm font-medium",
  {
    variants: {
      status: {
        [StepStatus.STEP_NOT_STARTED]: `${stepStatusIconVariant({ status: StepStatus.STEP_NOT_STARTED })} border-slate-300`,
        [StepStatus.STEP_RUNNING]: `${stepStatusIconVariant({ status: StepStatus.STEP_RUNNING })} border-blue-300`,
        [StepStatus.STEP_COMPLETED]: `${stepStatusIconVariant({ status: StepStatus.STEP_COMPLETED })} border-green-300`,
        [StepStatus.STEP_FAILURE]: `${stepStatusIconVariant({ status: StepStatus.STEP_FAILURE })} border-red-300`,
        [StepStatus.STEP_CANCELLED]: `${stepStatusIconVariant({ status: StepStatus.STEP_CANCELLED })} border-orange-300`,
      },
    },
    defaultVariants: {
      status: StepStatus.STEP_NOT_STARTED,
    },
  }
)

export const StepStatusIcon = ({
  status,
  className,
}: {
  status: StepStatus
  className?: string
}) => {
  switch (status) {
    case StepStatus.STEP_NOT_STARTED:
      return (
        <Moon className={cn(stepStatusIconVariant({ status }), className)} />
      )

    case StepStatus.STEP_RUNNING:
      return (
        <CircleDotDashed
          className={cn(
            stepStatusIconVariant({ status }),
            "animate-spin-slow",
            className
          )}
        />
      )
    case StepStatus.STEP_COMPLETED:
      return (
        <CheckCircle2
          className={cn(stepStatusIconVariant({ status }), className)}
        />
      )
    case StepStatus.STEP_FAILURE:
      return (
        <AlertCircle
          className={cn(stepStatusIconVariant({ status }), className)}
        />
      )
    case StepStatus.STEP_CANCELLED:
      return (
        <XCircle className={cn(stepStatusIconVariant({ status }), className)} />
      )
    default:
      break
  }
}
