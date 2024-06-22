import useMemberships from "../_hooks/use-memberships"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { useState } from "react"
import useDebounce from "@/hooks/use-debounce"
import AddMemberDialog from "./add-member-dialog"
import { useCurrentUserInfo } from "@/contexts/current-user-info-context"
import { User } from "lucide-react"

interface MembershipSectionProps {
  organizationId: number
}

export default function MembershipSection({
  organizationId,
}: MembershipSectionProps) {
  const [searchFilter, setSearchFilter] = useState("")

  // Delay filter execution by 0.5s at each filter change
  const { debouncedValue: debouncedFilter } = useDebounce(searchFilter, 500)
  const { members, refetchMembers } = useMemberships({
    orgId: organizationId,
    filter: debouncedFilter,
  })

  const userInfo = useCurrentUserInfo()

  return (
    <div className="space-y-5">
      <div>
        <h1 className="text-xl">Members</h1>
      </div>
      <div className="flex items-center">
        <Input
          placeholder="Search for member"
          className="max-w-xs"
          onChange={(e) => setSearchFilter(e.target.value)}
        />
        <AddMemberDialog
          existingMembers={members}
          organizationId={organizationId}
          refetchMembers={refetchMembers}
        >
          <Button variant={"outline"} className="ml-auto">
            Add Member
          </Button>
        </AddMemberDialog>
      </div>
      <div className="border rounded-md">
        <ul className="divide-y divide-slate-200">
          {members.map((member) => (
            <li key={member.user_id} className="px-4 py-4 flex items-center">
              <p>{member.name}</p>
              {member.user_id === userInfo?.user_id && (
                <User size="16" className="ml-2 text-blue-500" />
              )}
              <p className="text-sm text-muted-foreground ml-auto">
                {member.role}
              </p>
            </li>
          ))}
        </ul>
      </div>
    </div>
  )
}
