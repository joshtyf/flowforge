import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Search } from "lucide-react"

interface AddMemberDialogProps {
  children: React.ReactNode
}

export default function AddMemberDialog({ children }: AddMemberDialogProps) {
  return (
    <Dialog>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Add a member to the organization.</DialogTitle>
        </DialogHeader>
        <div className="relative">
          <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input placeholder="Search by Username" className="pl-8" />
        </div>
        <DialogFooter>
          <Button type="submit" className="w-full">
            Add
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
