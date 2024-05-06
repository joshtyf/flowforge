import { toast } from "@/components/ui/use-toast"
import { getServiceRequestLogs } from "@/lib/service"
import { ServiceRequestLogs } from "@/types/service-request"
import { useQuery } from "@tanstack/react-query"
import { useCallback, useEffect, useState } from "react"

export interface UseStepLogsOptions {
  serviceRequestId: string
  stepName: string
}

const useStepLogs = ({ serviceRequestId, stepName }: UseStepLogsOptions) => {
  const [offset, setOffset] = useState<number | undefined>()
  const [logs, setLogs] = useState<string[]>([])
  const [latestLogsData, setLatestLogsData] = useState<ServiceRequestLogs>()
  const [currentIntervalId, setCurrentIntervalId] =
    useState<ReturnType<typeof setInterval>>()

  const fetchData = useCallback(() => {
    getServiceRequestLogs(
      serviceRequestId,
      stepName,
      offset === -1 ? undefined : offset
    )
      .then((data) => {
        setLatestLogsData(data)
      })
      .catch((err) => {
        console.error(err)
        toast({
          title: "Fetching Service Request Logs Error",
          description: `Failed to fetch Service Request Logs for ${stepName}. Please try again later.`,
          variant: "destructive",
        })
      })
  }, [offset, serviceRequestId, stepName])

  useEffect(() => {
    fetchData()

    //  Clear existing interval
    clearInterval(currentIntervalId)
    // Call every 10s
    const intervalId = setInterval(() => fetchData(), 10000)
    setCurrentIntervalId(intervalId)

    // Cleanup function to clear the interval when the component unmounts
    return () => clearInterval(intervalId)
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [fetchData, offset])

  // To set logs and offset value when logs data is returned from API call
  useEffect(() => {
    if (latestLogsData) {
      const isOffsetSame = latestLogsData.next_offset === offset
      // Ignore if offset remains the same
      if (!isOffsetSame) {
        if (latestLogsData.end_of_logs) {
          // To prevent all logs for being appended to the current logs when EOL is reached
          setLogs(latestLogsData.logs)
        } else {
          setLogs((oldLogs) => oldLogs.concat(latestLogsData.logs))
        }
      }

      setOffset(latestLogsData?.next_offset)
    }
  }, [latestLogsData, offset])

  return { logs }
}

export default useStepLogs
