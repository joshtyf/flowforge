import { useParams } from "next/navigation"
import { useMemo } from "react"

const useOrganisationId = () => {
  // TODO: Uncomment logic once org id is implemented into URL route
  // const { organisationId } = useParams()
  // const organisationIdString = useMemo(
  //   () => (Array.isArray(organisationId) ? organisationId[0] : organisationId),
  //   [organisationId]
  // )
  return { organisationId: 1 }
}

export default useOrganisationId
