import { getMembersForOrg } from "@/lib/service"
import { UserInfo } from "@/types/user-profile"
import { useEffect, useMemo, useState } from "react"

interface UseMembershipOptions {
  orgId: number
  filter: string
}

export default function useMemberships({
  orgId,
  filter,
}: UseMembershipOptions) {
  const [members, setMembers] = useState<UserInfo[]>([])
  const [isLoadingMembers, setIsLoadingMembers] = useState(false)

  const filteredMembers = useMemo(() => {
    return members.filter((member) => member.name.includes(filter))
  }, [members, filter])

  useEffect(() => {
    setIsLoadingMembers(true)
    getMembersForOrg(orgId)
      .then((res) => setMembers(res.members))
      .catch((err) => console.error(err))
      .finally(() => setIsLoadingMembers(false))
  }, [orgId])

  const refetchMembers = () => {
    getMembersForOrg(orgId)
      .then((res) => setMembers(res.members))
      .catch((err) => console.error(err))
  }

  return { members: filteredMembers, isLoadingMembers, refetchMembers }
}
