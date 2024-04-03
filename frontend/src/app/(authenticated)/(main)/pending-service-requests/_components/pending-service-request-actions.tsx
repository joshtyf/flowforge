import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { MoreHorizontal } from "lucide-react"

interface ApproveServiceRequestActionsProps {
  pipelineId: string
  approveRequest: (pipelineId: string) => void
  rejectRequest: (pipelineId: string) => void
}

export default function ApproveServiceRequestActions({
  pipelineId,
  approveRequest,
  rejectRequest,
}: ApproveServiceRequestActionsProps) {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" className="h-8 w-8 p-0">
          <MoreHorizontal className="h-4 w-4" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="start">
        <DropdownMenuItem>
          <Button
            variant="ghost"
            className="text-green-700 hover:text-green-500"
            onClick={() => approveRequest(pipelineId)}
          >
            Approve
          </Button>
        </DropdownMenuItem>
        <DropdownMenuItem>
          {/* TODO: Add on click logic*/}
          <Button
            variant="ghost"
            className="text-red-700 hover:text-red-500"
            onClick={() => rejectRequest(pipelineId)}
          >
            Reject
          </Button>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
