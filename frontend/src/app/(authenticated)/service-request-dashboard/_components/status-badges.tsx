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
} from "lucide-react"

const statusBadgeVariant = cva("rounded-lg border text-sm font-medium", {
  variants: {
    status: {
      [ServiceRequestStatus.NOT_STARTED]: " text-slate-500 border-slate-300",
      [ServiceRequestStatus.PENDING]: "text-yellow-500 border-yellow-300",
      [ServiceRequestStatus.RUNNING]: "text-blue-500 border-blue-300",
      [ServiceRequestStatus.SUCCESS]: "text-green-500 border-green-300",
      [ServiceRequestStatus.FAILURE]: "text-red-500 border-red-300",
      [ServiceRequestStatus.CANCELLED]: "text-orange-500 border-orange-300",
    },
  },
  defaultVariants: {
    status: ServiceRequestStatus.NOT_STARTED,
  },
})

type StatusBadgeProps = {
  status: ServiceRequestStatus
}

const StatusIcon = ({ status }: { status: ServiceRequestStatus }) => {
  switch (status) {
    case ServiceRequestStatus.NOT_STARTED:
      return <Moon />
    case ServiceRequestStatus.PENDING:
      return <CircleEllipsis />
    case ServiceRequestStatus.RUNNING:
      return <CircleDotDashed />
    case ServiceRequestStatus.SUCCESS:
      return <CheckCircle2 />
    case ServiceRequestStatus.FAILURE:
      return <AlertCircle />
    case ServiceRequestStatus.CANCELLED:
      return <XCircle />
    default:
      break
  }
}

export default function StatusBadge({ status }: StatusBadgeProps) {
  return (
    <div
      className={cn(
        statusBadgeVariant({ status }),
        "flex w-[10rem] py-2 pl-4 pr-5 items-center justify-center space-x-2"
      )}
    >
      <StatusIcon status={status} />
      <p>{status}</p>
    </div>
  )
}
