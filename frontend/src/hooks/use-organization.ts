import { getCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { useMemo } from "react"

const useOrganization = () => {
  const organizationId = useMemo(
    () => parseInt(getCookie("org_id") as string, 10),
    []
  )

  const organizationName = useMemo(() => getCookie("org_name"), [])
  const router = useRouter()
  if (!organizationId) {
    router.push("/organization")
  }

  return { organizationId, organizationName }
}

export default useOrganization
