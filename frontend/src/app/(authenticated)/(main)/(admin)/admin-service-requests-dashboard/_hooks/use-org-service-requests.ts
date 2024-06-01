import { toast } from "@/components/ui/use-toast"
import useOrganizationId from "@/hooks/use-organization-id"
import { getAllServiceRequestForAdmin } from "@/lib/service"
import { useQuery } from "@tanstack/react-query"
import { useMemo } from "react"

interface UseOrgServiceRequestsOptions {
  page: number
  pageSize: number
}

const useOrgServiceRequests = ({
  page,
  pageSize,
}: UseOrgServiceRequestsOptions) => {
  const { organizationId } = useOrganizationId()
  const { isLoading, data } = useQuery({
    queryKey: ["org_service_requests", organizationId, page, pageSize],
    queryFn: () =>
      getAllServiceRequestForAdmin(organizationId, page, pageSize).catch(
        (err) => {
          console.error(err)
          toast({
            title: "Fetching Service Requests Error",
            description:
              "Failed to fetch Service Requests for organization. Please try again later.",
            variant: "destructive",
          })
        }
      ),
    refetchInterval: 2000,
  })

  const noOfPages = useMemo(
    () =>
      data?.metadata.total_count
        ? Math.ceil(data?.metadata.total_count / pageSize)
        : undefined,
    [data, pageSize]
  )

  return { orgServiceRequestsData: data, noOfPages }
}

export default useOrgServiceRequests
