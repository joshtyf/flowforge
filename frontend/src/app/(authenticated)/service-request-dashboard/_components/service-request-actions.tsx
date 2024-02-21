import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { MoreHorizontal } from "lucide-react"

interface ServiceRequestActionsProps {
  pipelineId: string
  onCancelRequest: (pipelineId: string) => void
}

export default function ServiceRequestActions({
  pipelineId,
  onCancelRequest,
}: ServiceRequestActionsProps) {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" className="h-8 w-8 p-0">
          <MoreHorizontal className="h-4 w-4" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="start">
        {/* TODO: Add more actions for requests*/}
        {/* TODO: Add on click logic*/}
        <DropdownMenuItem>
          <Button variant="ghost">View Details</Button>
        </DropdownMenuItem>
        <DropdownMenuItem>
          <Button variant="ghost" onClick={() => onCancelRequest(pipelineId)}>
            Cancel Request
          </Button>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
