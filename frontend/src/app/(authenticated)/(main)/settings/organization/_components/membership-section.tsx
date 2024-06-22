import useMemberships from "../_hooks/use-memberships"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { useState } from "react"
import useDebounce from "@/hooks/use-debounce"
import AddMemberDialog from "./add-member-dialog"
import { useCurrentUserInfo } from "@/contexts/current-user-info-context"
import { MoreHorizontal, User } from "lucide-react"
import MemberActions from "./member-actions"

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
              <div>
                <span className="flex items-center">
                  <p>{member.name}</p>
                  {member.user_id === userInfo?.user_id && (
                    <User size="16" className="ml-2 text-blue-500" />
                  )}
                </span>

                <p className="text-sm text-muted-foreground">{member.role}</p>
              </div>
              <MemberActions>
                <Button variant="ghost" className="h-8 w-8 p-0 ml-auto">
                  <MoreHorizontal className="h-4 w-4" />
                </Button>
              </MemberActions>
            </li>
          ))}
        </ul>
      </div>
    </div>
  )
}
