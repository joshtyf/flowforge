import { getCookie, setCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { useMemo, useState } from "react"

const useOrganization = () => {
  const [organizationName, setOrganizationName] = useState(
    getCookie("org_name")
  )
  const organizationId = useMemo(
    () => parseInt(getCookie("org_id") as string, 10),
    []
  )

  const updateOrgNameInCookie = (name: string) => {
    setCookie("org_name", name)
    setOrganizationName(name)
  }

  const router = useRouter()
  if (!organizationId) {
    router.push("/organization")
  }

  return { organizationId, organizationName, updateOrgNameInCookie }
}

export default useOrganization
