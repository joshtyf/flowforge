import { Skeleton } from "@/components/ui/skeleton"

interface StepLogsProps {
  serviceRequestId: string
  stepName: string
}

export default function StepLogs({
  serviceRequestId,
  stepName,
}: StepLogsProps) {
  return stepName === "" ? (
    <Skeleton className="h-[70vh]" />
  ) : (
    <div className="bg-gray-900 text-green-300 border-none rounded-lg p-3 focus:ring-2 focus:ring-offset-2 focus:ring-green-500 font-mono h-[70vh] overflow-auto">
      {stepName}
    </div>
  )
}
