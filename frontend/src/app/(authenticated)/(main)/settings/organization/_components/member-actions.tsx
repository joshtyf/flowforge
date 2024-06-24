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

interface MemberActionsProps {
  children: React.ReactNode
  member: UserInfo
  isOwner?: boolean
}

export default function MemberActions({
  children,
  member,
  isOwner = false,
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
        onConfirm={async () => {}}
        title={`Remove ${member.name} from organization?`}
      />
      <MemberActionAlertDialog
        open={openTransferOwnershipDialog}
        setOpen={setOpenTransferOwnershipDialog}
        onConfirm={async () => {}}
        title={`Transfer ownership of organization to ${member.name}?`}
      />
    </DropdownMenu>
  )
}
