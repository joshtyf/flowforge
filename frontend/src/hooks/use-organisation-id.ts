import { getCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { useMemo } from "react"

const useOrganisationId = () => {
  // TODO: Uncomment logic once org id is implemented into cookies
  // const organisationId = useMemo(
  //   () => getCookie("organisation_id") as string,
  //   []
  // )
  // const router = useRouter();
  // if (!organisationId) {
  //   router.push("/organisation")
  // }

  return { organisationId: 1 }
}

export default useOrganisationId
