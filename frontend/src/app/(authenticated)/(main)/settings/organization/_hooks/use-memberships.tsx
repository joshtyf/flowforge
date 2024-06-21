import { getMembersForOrg } from "@/lib/service"
import { UserInfo } from "@/types/user-profile"
import { useEffect, useMemo, useState } from "react"

export default function useMemberships({ orgId }: { orgId: number }) {
  const [members, setMembers] = useState<UserInfo[]>([])
  const [isLoadingMembers, setIsLoadingMembers] = useState(false)

  const [filter, setFilter] = useState("")

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

  return { members: filteredMembers, isLoadingMembers, setFilter }
}
