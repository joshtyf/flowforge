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
import { leaveOrganization } from "@/lib/service"
import { toast } from "@/components/ui/use-toast"
import { useRouter } from "next/navigation"

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
  const { isOwner, isAdmin, refetchMemberships } = useUserMemberships()

  const router = useRouter()

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
        <div className="flex space-x-2 ml-auto pl-2">
          <AddMemberDialog
            existingMembers={members}
            organizationId={organizationId}
            refetchMembers={refetchMembers}
          >
            <Button variant={"outline"} className={`${!isAdmin && "hidden"}`}>
              Add Member
            </Button>
          </AddMemberDialog>
          <LeaveOrganizationDialog
            onConfirm={async () => {
              await leaveOrganization(organizationId)
                .then(() => {
                  router.push("/organization")
                })
                .catch((e) => {
                  toast({
                    title: "Leave Organization Error",
                    description:
                      "Unable to leave the organization. Please try again later.",
                    variant: "destructive",
                  })
                  console.error(e)
                })
              return
            }}
            isOwner={isOwner}
          >
            <Button variant={"destructive"}>Leave Organization</Button>
          </LeaveOrganizationDialog>
        </div>
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
              <MemberActions
                member={member}
                isOwner={isOwner}
                organizationId={organizationId}
                refetchMembers={refetchMembers}
                refetchMemberships={refetchMemberships}
              >
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
