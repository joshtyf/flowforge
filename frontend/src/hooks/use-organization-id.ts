import { getCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { useMemo } from "react"

const useOrganizationId = () => {
  const organizationId = useMemo(
    () => parseInt(getCookie("org_id") as string, 10),
    []
  )
  const router = useRouter()
  if (!organizationId) {
    router.push("/organization")
  }

  return { organizationId }
}

export default useOrganizationId
