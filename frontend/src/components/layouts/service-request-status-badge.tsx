import { cn } from "@/lib/utils"
import { ServiceRequestStatus } from "@/types/service-request"
import { cva } from "class-variance-authority"
import {
  AlertCircle,
  CheckCircle2,
  CircleDotDashed,
  CircleEllipsis,
  Moon,
  XCircle,
} from "lucide-react"

const statusIconVariant = cva("", {
  variants: {
    status: {
      [ServiceRequestStatus.NOT_STARTED]: "text-slate-500",
      [ServiceRequestStatus.PENDING]: "text-yellow-500",
      [ServiceRequestStatus.RUNNING]: "text-blue-500",
      [ServiceRequestStatus.COMPLETED]: "text-green-500",
      [ServiceRequestStatus.FAILED]: "text-red-500",
      [ServiceRequestStatus.CANCELLED]: "text-orange-500",
    },
  },
  defaultVariants: {
    status: ServiceRequestStatus.NOT_STARTED,
  },
})

export const serviceRequestStatusBadgeVariant = cva(
  "rounded-lg border text-sm font-medium",
  {
    variants: {
      status: {
        [ServiceRequestStatus.NOT_STARTED]: `${statusIconVariant({ status: ServiceRequestStatus.NOT_STARTED })} border-slate-300`,
        [ServiceRequestStatus.PENDING]: `${statusIconVariant({ status: ServiceRequestStatus.PENDING })} border-yellow-300`,
        [ServiceRequestStatus.RUNNING]: `${statusIconVariant({ status: ServiceRequestStatus.RUNNING })} border-blue-300`,
        [ServiceRequestStatus.COMPLETED]: `${statusIconVariant({ status: ServiceRequestStatus.COMPLETED })} border-green-300`,
        [ServiceRequestStatus.FAILED]: `${statusIconVariant({ status: ServiceRequestStatus.FAILED })} border-red-300`,
        [ServiceRequestStatus.CANCELLED]: `${statusIconVariant({ status: ServiceRequestStatus.CANCELLED })} border-orange-300`,
      },
    },
    defaultVariants: {
      status: ServiceRequestStatus.NOT_STARTED,
    },
  }
)

type StatusBadgeProps = {
  status: ServiceRequestStatus
}

const ServiceRequestStatusIcon = ({
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
    case ServiceRequestStatus.COMPLETED:
      return (
        <CheckCircle2
          className={cn(statusIconVariant({ status }), className)}
        />
      )
    case ServiceRequestStatus.FAILED:
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

export function ServiceRequestStatusBadge({ status }: StatusBadgeProps) {
  return (
    <div
      className={cn(
        serviceRequestStatusBadgeVariant({ status }),
        "flex w-fit py-2 pl-4 pr-5 items-center space-x-2"
      )}
    >
      <ServiceRequestStatusIcon status={status} />
      <p className="w-[5rem] flex justify-center">{status}</p>
    </div>
  )
}
