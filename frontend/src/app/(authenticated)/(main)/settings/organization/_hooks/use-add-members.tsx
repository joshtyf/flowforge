import { getAllUsers } from "@/lib/service"
import { UserInfo } from "@/types/user-profile"
import { useEffect, useMemo, useState } from "react"

interface UseAddMembersOptions {
  existingMembers: UserInfo[]
  filter: string
}

export default function useAddMembers({
  existingMembers,
  filter,
}: UseAddMembersOptions) {
  const [allUsers, setAllUsers] = useState<UserInfo[]>()

  const [selectedMember, setSelectedMember] = useState<UserInfo>()

  useEffect(() => {
    getAllUsers()
      .then((users) => setAllUsers(users))
      .catch((err) => console.error(err))
  })

  const allUsersOutsideOrg = useMemo(() => {
    const existingMemberIdSet = new Set(existingMembers.map((u) => u.user_id))
    return allUsers?.filter((user) => !existingMemberIdSet.has(user.user_id))
  }, [allUsers, existingMembers])

  const filteredMembers = useMemo(() => {
    return allUsersOutsideOrg?.filter((member) => member.name.includes(filter))
  }, [allUsersOutsideOrg, filter])

  return { allUsers: filteredMembers, selectedMember, setSelectedMember }
}
