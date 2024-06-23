import useMemberships from "../_hooks/use-memberships"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { useState } from "react"
import useDebounce from "@/hooks/use-debounce"
import AddMemberDialog from "./add-member-dialog"
import { useCurrentUserInfo } from "@/contexts/current-user-info-context"
import { MoreHorizontal, User } from "lucide-react"
import MemberActions from "./member-actions"
import { useUserMemberships } from "@/contexts/user-memberships-context"
import { Role } from "@/types/membership"
import { cn } from "@/lib/utils"
import LeaveOrganizationDialog from "./leave-organization-dialog"

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
  const { isOwner, isAdmin } = useUserMemberships()

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
          <Button
            variant={"outline"}
            className={cn("ml-auto", !isAdmin && "hidden")}
          >
            Add Member
          </Button>
        </AddMemberDialog>
        <LeaveOrganizationDialog onConfirm={() => {}} isOwner={isOwner}>
          <Button
            className={!isAdmin ? "ml-auto" : "ml-3"}
            variant={"destructive"}
          >
            Leave Organization
          </Button>
        </LeaveOrganizationDialog>
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
              <MemberActions member={member} isOwner={isOwner}>
                <Button
                  variant="ghost"
                  className="h-8 w-8 p-0 ml-auto"
                  disabled={
                    // Disable action button if the target member is current user or the owner
                    member.user_id === userInfo?.user_id ||
                    member.role === Role.Owner
                  }
                >
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
