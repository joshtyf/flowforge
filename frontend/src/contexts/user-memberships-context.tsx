import useOrganization from "@/hooks/use-organization"
import { getUserMemberships } from "@/lib/service"
import { Role, UserMemberships } from "@/types/membership"
import { createContext, useContext, useEffect, useMemo, useState } from "react"

interface MembershipContextValue {
  isAdmin?: boolean
  isOwner?: boolean
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
  const { organizationId } = useOrganization()

  const isAdminOfCurrentOrg = useMemo(() => {
    return userMemberships?.memberships.some(
      (membership) =>
        membership.org_id === organizationId &&
        (membership.role === Role.Admin || membership.role === Role.Owner)
    )
  }, [userMemberships, organizationId])

  const isOwnerOfCurrentOrg = useMemo(() => {
    return userMemberships?.memberships.some(
      (membership) =>
        membership.org_id === organizationId && membership.role === Role.Owner
    )
  }, [userMemberships, organizationId])

  return (
    <MembershipContext.Provider
      value={{
        isAdmin: isAdminOfCurrentOrg,
        isOwner: isOwnerOfCurrentOrg,
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
