import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import MemberActionAlertDialog from "./member-action-alert-dialog"
import { UserInfo } from "@/types/user-profile"
import { useState } from "react"
import { Role } from "@/types/membership"
import { removeMember, transferOwnership } from "@/lib/service"
import { toast } from "@/components/ui/use-toast"

interface MemberActionsProps {
  children: React.ReactNode
  member: UserInfo
  isOwner?: boolean
  organizationId: number
  refetchMembers: () => void
  refetchMemberships: () => void
}

export default function MemberActions({
  children,
  member,
  isOwner = false,
  organizationId,
  refetchMembers,
  refetchMemberships,
}: MemberActionsProps) {
  const [openPromoteToAdminDialog, setOpenPromoteToAdminDialog] =
    useState(false)

  const [openDemoteToMemberDialog, setOpenDemoteToMemberDialog] =
    useState(false)

  const [openRemoveFromOrgDialog, setOpenRemoveFromOrgDialog] = useState(false)
  const [openTransferOwnershipDialog, setOpenTransferOwnershipDialog] =
    useState(false)

  const isTargetMember = member.role === Role.Member
  const isTargetAdmin = member.role === Role.Admin

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>{children}</DropdownMenuTrigger>
      <DropdownMenuContent align="start">
        {isTargetMember && (
          <DropdownMenuItem>
            <Button
              onClick={() => setOpenPromoteToAdminDialog(true)}
              variant="ghost"
            >
              Promote to Admin
            </Button>
          </DropdownMenuItem>
        )}
        {isTargetAdmin && (
          <DropdownMenuItem>
            <Button
              onClick={() => setOpenDemoteToMemberDialog(true)}
              variant="ghost"
            >
              Demote to Member
            </Button>
          </DropdownMenuItem>
        )}
        {(isTargetMember || isTargetAdmin) && (
          <DropdownMenuItem>
            <Button
              onClick={() => setOpenRemoveFromOrgDialog(true)}
              variant="ghost"
            >
              Remove from Organization
            </Button>
          </DropdownMenuItem>
        )}
        {isOwner && isTargetAdmin && (
          <DropdownMenuItem>
            <Button
              onClick={() => setOpenTransferOwnershipDialog(true)}
              variant="ghost"
            >
              Transfer Ownership
            </Button>
          </DropdownMenuItem>
        )}
      </DropdownMenuContent>
      <MemberActionAlertDialog
        open={openPromoteToAdminDialog}
        setOpen={setOpenPromoteToAdminDialog}
        onConfirm={async () => {}}
        title={`Promote ${member.name} to Admin?`}
      />
      <MemberActionAlertDialog
        open={openDemoteToMemberDialog}
        setOpen={setOpenDemoteToMemberDialog}
        onConfirm={async () => {}}
        title={`Demote ${member.name} to Member?`}
      />
      <MemberActionAlertDialog
        open={openRemoveFromOrgDialog}
        setOpen={setOpenRemoveFromOrgDialog}
        onConfirm={async () => {
          await removeMember(member.user_id, organizationId, member.role!)
            .then(() => {
              toast({
                title: "Remove Member Successful",
                description: `${member.name} has been removed from the organization.`,
                variant: "success",
              })
              refetchMembers()
            })
            .catch((err) => {
              toast({
                title: "Remove Member Failure",
                description: `Error removing ${member.name} from the organization. Please try again later.`,
                variant: "destructive",
              })
              console.error(err)
            })
          return
        }}
        title={`Remove ${member.name} from organization?`}
      />
      <MemberActionAlertDialog
        open={openTransferOwnershipDialog}
        setOpen={setOpenTransferOwnershipDialog}
        onConfirm={async () => {
          await transferOwnership(member.user_id, organizationId)
            .then(() => {
              toast({
                title: "Transfer Ownership Successful",
                description: `${member.name} has been promoted to Owner and you have been demoted to Admin.`,
                variant: "success",
              })
              refetchMembers()
              refetchMemberships()
            })
            .catch((err) => {
              toast({
                title: "Transfer Ownership Failure",
                description: `Error transferring ownership. Please try again later.`,
                variant: "destructive",
              })
              console.error(err)
            })
          return
        }}
        title={`Transfer ownership of organization to ${member.name}?`}
      />
    </DropdownMenu>
  )
}
