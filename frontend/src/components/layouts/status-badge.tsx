import { cn } from "@/lib/utils"
import { ServiceRequestStatus } from "@/types/service-request"
import { cva } from "class-variance-authority"
import {
  AlertCircle,
  CheckCircle2,
  CircleDotDashed,
  CircleEllipsis,
  CircleOff,
  Moon,
  XCircle,
  XOctagon,
} from "lucide-react"

export const statusIconVariant = cva("", {
  variants: {
    status: {
      [ServiceRequestStatus.NOT_STARTED]: "text-slate-500",
      [ServiceRequestStatus.PENDING]: "text-yellow-500",
      [ServiceRequestStatus.REJECTED]: "text-red-800",
      [ServiceRequestStatus.RUNNING]: "text-blue-500",
      [ServiceRequestStatus.SUCCESS]: "text-green-500",
      [ServiceRequestStatus.COMPLETED]: "text-green-500",
      [ServiceRequestStatus.FAILURE]: "text-red-500",
      [ServiceRequestStatus.CANCELLED]: "text-orange-500",
    },
  },
  defaultVariants: {
    status: ServiceRequestStatus.NOT_STARTED,
  },
})

export const statusBadgeVariant = cva("rounded-lg border text-sm font-medium", {
  variants: {
    status: {
      [ServiceRequestStatus.NOT_STARTED]: `${statusIconVariant({ status: ServiceRequestStatus.NOT_STARTED })} border-slate-300`,
      [ServiceRequestStatus.PENDING]: `${statusIconVariant({ status: ServiceRequestStatus.PENDING })} border-yellow-300`,
      [ServiceRequestStatus.REJECTED]: `${statusIconVariant({ status: ServiceRequestStatus.REJECTED })} border-yellow-800`,
      [ServiceRequestStatus.RUNNING]: `${statusIconVariant({ status: ServiceRequestStatus.RUNNING })} border-blue-300`,
      [ServiceRequestStatus.SUCCESS]: `${statusIconVariant({ status: ServiceRequestStatus.SUCCESS })} border-green-300`,
      [ServiceRequestStatus.COMPLETED]: `${statusIconVariant({ status: ServiceRequestStatus.COMPLETED })} border-green-300`,
      [ServiceRequestStatus.FAILURE]: `${statusIconVariant({ status: ServiceRequestStatus.FAILURE })} border-red-300`,
      [ServiceRequestStatus.CANCELLED]: `${statusIconVariant({ status: ServiceRequestStatus.CANCELLED })} border-orange-300`,
    },
  },
  defaultVariants: {
    status: ServiceRequestStatus.NOT_STARTED,
  },
})

type StatusBadgeProps = {
  status: ServiceRequestStatus
}

export const StatusIcon = ({
  status,
  className,
}: {
  status: ServiceRequestStatus
  className?: string
}) => {
  switch (status) {
    case ServiceRequestStatus.NOT_STARTED:
      return <Moon className={cn(statusIconVariant({ status }), className)} />
    case ServiceRequestStatus.PENDING:
      return (
        <CircleEllipsis
          className={cn(statusIconVariant({ status }), className)}
        />
      )
    case ServiceRequestStatus.REJECTED:
      return (
        <XOctagon className={cn(statusIconVariant({ status }), className)} />
      )
    case ServiceRequestStatus.RUNNING:
      return (
        <CircleDotDashed
          className={cn(
            statusIconVariant({ status }),
            "animate-spin-slow",
            className
          )}
        />
      )
    case ServiceRequestStatus.SUCCESS:
    case ServiceRequestStatus.COMPLETED:
      return (
        <CheckCircle2
          className={cn(statusIconVariant({ status }), className)}
        />
      )
    case ServiceRequestStatus.FAILURE:
      return (
        <AlertCircle className={cn(statusIconVariant({ status }), className)} />
      )
    case ServiceRequestStatus.CANCELLED:
      return (
        <XCircle className={cn(statusIconVariant({ status }), className)} />
      )
    default:
      break
  }
}

export function StatusBadge({ status }: StatusBadgeProps) {
  return (
    <div
      className={cn(
        statusBadgeVariant({ status }),
        "flex w-fit py-2 pl-4 pr-5 items-center space-x-2"
      )}
    >
      <StatusIcon status={status} />
      <p className="w-[5rem] flex justify-center">{status}</p>
    </div>
  )
}
