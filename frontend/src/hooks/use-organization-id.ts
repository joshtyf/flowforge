import { getCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { useMemo } from "react"

const useOrganizationId = () => {
  // TODO: Uncomment logic once org id is implemented into cookies
  // const organizationId = useMemo(
  //   () => getCookie("organization_id") as string,
  //   []
  // )
  // const router = useRouter();
  // if (!organizationId) {
  //   router.push("/organization")
  // }

  return { organizationId: 1 }
}

export default useOrganizationId
