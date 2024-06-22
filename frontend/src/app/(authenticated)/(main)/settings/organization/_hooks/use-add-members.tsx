import { Skeleton } from "@/components/ui/skeleton"
import { toast } from "@/components/ui/use-toast"
import { createMembershipForOrg, getAllUsers } from "@/lib/service"
import { UserInfo } from "@/types/user-profile"
import { useEffect, useMemo, useState } from "react"

interface UseAddMembersOptions {
  existingMembers: UserInfo[]
  filter: string
  organizationId: number
  refetchMembers: () => void
}

export default function useAddMembers({
  existingMembers,
  filter,
  organizationId,
  refetchMembers,
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

  const handleAddMember = () => {
    if (!selectedMember || !selectedMember.role) {
      return
    }
    createMembershipForOrg(
      selectedMember.user_id,
      organizationId,
      selectedMember.role
    )
      .then(() => {
        toast({
          title: "Add Member Successful",
          description: `${selectedMember.name} has been added to the organization.`,
          variant: "success",
        })
        refetchMembers()
        setSelectedMember(undefined)
      })
      .catch((err) => {
        toast({
          title: "Add Member Failure",
          description: `Error adding ${selectedMember.name} to the organization.`,
          variant: "destructive",
        })
        console.error(err)
      })
  }

  return {
    allUsers: filteredMembers,
    selectedMember,
    setSelectedMember,
    handleAddMember,
  }
}
