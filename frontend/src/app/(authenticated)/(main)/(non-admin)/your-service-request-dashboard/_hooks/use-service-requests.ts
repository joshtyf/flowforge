import { toast } from "@/components/ui/use-toast"
import useOrganization from "@/hooks/use-organization"
import { getAllServiceRequest } from "@/lib/service"
import { useQuery } from "@tanstack/react-query"
import { useMemo } from "react"

interface UseServiceRequestProps {
  page: number
  pageSize: number
}

const useServiceRequests = ({ page, pageSize }: UseServiceRequestProps) => {
  const { organizationId } = useOrganization()
  const { isLoading, data } = useQuery({
    queryKey: ["user_service_requests", page, pageSize],
    queryFn: () => {
      return getAllServiceRequest(organizationId, page, pageSize).catch(
        (err) => {
          console.error(err)
          toast({
            title: "Fetching Service Requests Error",
            description:
              "Failed to fetch Service Requests for user. Please try again later.",
            variant: "destructive",
          })
        }
      )
    },
    refetchInterval: 2000,
  })

  const noOfPages = useMemo(
    () =>
      data?.metadata.total_count
        ? Math.ceil(data?.metadata.total_count / pageSize)
        : undefined,
    [data, pageSize]
  )

  return {
    response: data,
    isLoading,
    noOfPages,
  }
}

export default useServiceRequests
