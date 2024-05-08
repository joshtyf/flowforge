import { Skeleton } from "@/components/ui/skeleton"
import useStepLogs from "../_hooks/use-step-logs"

interface StepLogsProps {
  serviceRequestId: string
  stepName: string
}

export default function StepLogs({
  serviceRequestId,
  stepName,
}: StepLogsProps) {
  const { logs } = useStepLogs({ serviceRequestId, stepName })

  return stepName === "" ? (
    <Skeleton className="h-[70vh]" />
  ) : (
    <div className="bg-gray-900 text-green-300 border-none rounded-lg p-3 focus:ring-2 focus:ring-offset-2 focus:ring-green-500 font-mono h-[70vh] overflow-auto">
      {logs.map((log, index) => (
        <div key={index} className="flex space-x-2">
          <p className="opacity-70">{index + 1}</p>
          <p> {log}</p>
        </div>
      ))}
    </div>
  )
}
