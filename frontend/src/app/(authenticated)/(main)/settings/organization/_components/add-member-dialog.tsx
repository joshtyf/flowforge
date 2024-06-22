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
import useAddMembers from "../_hooks/use-add-members"
import { UserInfo } from "@/types/user-profile"
import { useState } from "react"
import useDebounce from "@/hooks/use-debounce"

interface AddMemberDialogProps {
  children: React.ReactNode
  existingMembers: UserInfo[]
}

export default function AddMemberDialog({
  children,
  existingMembers,
}: AddMemberDialogProps) {
  const [searchFilter, setSearchFilter] = useState("")
  // Delay filter execution by 0.5s at each filter change
  const { debouncedValue: debouncedFilter } = useDebounce(searchFilter, 500)
  const { allUsers, selectedMember, setSelectedMember } = useAddMembers({
    existingMembers,
    filter: debouncedFilter,
  })

  return (
    <Dialog>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add member to the organization</DialogTitle>
        </DialogHeader>
        <div className="relative">
          <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input
            placeholder="Search User by Username"
            className="pl-8"
            onChange={(e) => setSearchFilter(e.target.value)}
          />
        </div>
        <div className="max-h-[20rem] overflow-y-auto">
          <div className="border rounded-md">
            <ul className="divide-y divide-slate-200">
              {allUsers?.map((user) => (
                <li
                  key={user.user_id}
                  className="px-4 py-2 cursor-pointer hover:text-blue-500"
                >
                  <p>{user.name}</p>
                </li>
              ))}
            </ul>
          </div>
        </div>
        <DialogFooter>
          <Button type="submit" className="w-full">
            Add User
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
