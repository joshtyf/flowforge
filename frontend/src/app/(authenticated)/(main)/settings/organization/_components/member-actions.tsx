import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

interface MemberActionsProps {
  children: React.ReactNode
}

export default function MemberActions({ children }: MemberActionsProps) {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>{children}</DropdownMenuTrigger>
      <DropdownMenuContent align="start">
        <DropdownMenuItem>
          <Button variant="ghost">Promote to Admin</Button>
        </DropdownMenuItem>
        <DropdownMenuItem>
          <Button variant="ghost">Demote to Member</Button>
        </DropdownMenuItem>
        <DropdownMenuItem>
          <Button variant="ghost">Transfer Ownership</Button>
        </DropdownMenuItem>
        <DropdownMenuItem>
          <Button variant="ghost">Remove from Organization</Button>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
