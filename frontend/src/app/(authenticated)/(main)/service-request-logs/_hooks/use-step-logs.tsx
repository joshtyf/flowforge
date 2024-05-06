import { toast } from "@/components/ui/use-toast"
import { getServiceRequestLogs } from "@/lib/service"
import { useQuery } from "@tanstack/react-query"
import { useEffect, useState } from "react"

export interface UseStepLogsOptions {
  serviceRequestId: string
  stepName: string
}

const useStepLogs = ({ serviceRequestId, stepName }: UseStepLogsOptions) => {
  const [offset, setOffset] = useState<number | undefined>()
  const [logs, setLogs] = useState<string[]>([])

  const { data: latestLogsData, refetch: refetchLatestLogsData } = useQuery({
    queryKey: [stepName, serviceRequestId, "logs"],
    queryFn: () =>
      getServiceRequestLogs(serviceRequestId, stepName, offset).catch((err) => {
        console.error(err)
        toast({
          title: "Fetching Service Request Logs Error",
          description: `Failed to fetch Service Request Logs for ${stepName}. Please try again later.`,
          variant: "destructive",
        })
      }),
    // 10s interval fetch time
    refetchInterval: 10000,
  })

  // To set logs and offset value when logs data is returned from API call
  useEffect(() => {
    if (latestLogsData && !latestLogsData.end_of_logs) {
      if (offset === undefined || latestLogsData.next_offset > offset) {
        // Only set new logs if there is a difference in offset
        setLogs((oldLogs) => oldLogs.concat(latestLogsData.logs))
        // Only set offset if its not EOL and it is undefined (initial value) or there is a change in offset
        setOffset(latestLogsData?.next_offset)
      }
    }
  }, [latestLogsData, offset])

  // To set logs and offset value when logs data is returned from API call
  useEffect(() => {
    if (offset && offset !== -1) {
      refetchLatestLogsData()
    }
  }, [offset, refetchLatestLogsData])

  return { logs }
}

export default useStepLogs
