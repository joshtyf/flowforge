import useOrganizationId from "@/hooks/use-organization-id"
import { getUserMemberships } from "@/lib/service"
import { Role, UserMemberships } from "@/types/membership"
import { createContext, useContext, useEffect, useMemo, useState } from "react"

interface MembershipContextValue {
  isAdmin?: boolean
}

const MembershipContext = createContext<MembershipContextValue | null>(null)

export function UserMembershipsProvider({
  children,
}: {
  children: React.ReactNode
}) {
  const [userMemberships, setUserMemberships] = useState<UserMemberships>()

  useEffect(() => {
    getUserMemberships()
      .then(setUserMemberships)
      .catch((err) => {
        console.error(err)
      })
  }, [])
  const { organizationId } = useOrganizationId()

  const isAdminOfCurrentOrg = useMemo(() => {
    return userMemberships?.memberships.some(
      (membership) =>
        membership.org_id === organizationId &&
        (membership.role === Role.Admin || membership.role === Role.Owner)
    )
  }, [userMemberships, organizationId])

  return (
    <MembershipContext.Provider
      value={{
        isAdmin: isAdminOfCurrentOrg,
      }}
    >
      {children}
    </MembershipContext.Provider>
  )
}

export function useUserMemberships() {
  const context = useContext(MembershipContext)
  if (!context) {
    throw new Error(
      "useUserMemberships must be used within a UserMembershipsProvider"
    )
  }

  return context
}
