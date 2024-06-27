import { Button, ButtonWithSpinner } from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Search, X } from "lucide-react"
import useAddMembers from "../_hooks/use-add-members"
import { UserInfo } from "@/types/user-profile"
import { useState } from "react"
import useDebounce from "@/hooks/use-debounce"
import { Role } from "@/types/membership"
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { Separator } from "@/components/ui/separator"

interface AddMemberDialogProps {
  children: React.ReactNode
  existingMembers: UserInfo[]
  organizationId: number
  refetchMembers: () => void
}

export default function AddMemberDialog({
  children,
  existingMembers,
  organizationId,
  refetchMembers,
}: AddMemberDialogProps) {
  const [searchFilter, setSearchFilter] = useState("")
  // Delay filter execution by 0.5s at each filter change
  const { debouncedValue: debouncedFilter } = useDebounce(searchFilter, 500)
  const {
    allUsers,
    selectedMember,
    setSelectedMember,
    handleAddMember,
    isAddingMember,
  } = useAddMembers({
    existingMembers,
    filter: debouncedFilter,
    organizationId,
    refetchMembers,
  })

  return (
    <Dialog>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add member to the organization</DialogTitle>
        </DialogHeader>

        {selectedMember ? (
          <>
            <div className="border rounded-md flex items-center">
              <p className="px-4 py-2">{selectedMember.name}</p>
              <X
                size="20"
                className="ml-auto mr-2 cursor-pointer"
                onClick={() => setSelectedMember(undefined)}
              />
            </div>
            <p className="text-sm px-1">
              Select a role for {selectedMember.name}
            </p>
            <Select
              defaultValue={selectedMember.role}
              onValueChange={(value) =>
                setSelectedMember({ ...selectedMember, role: value as Role })
              }
            >
              <SelectTrigger value={selectedMember.role} className="">
                <SelectValue placeholder="Select role for member" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectItem value={Role.Member}>{Role.Member}</SelectItem>
                  <SelectItem value={Role.Admin}>{Role.Admin}</SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
            <Separator className="w-full" />
          </>
        ) : (
          <>
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
                      onClick={() => {
                        // Set default role as Member
                        setSelectedMember({ ...user, role: Role.Member })
                        setSearchFilter("")
                      }}
                    >
                      <p>{user.name}</p>
                    </li>
                  ))}
                </ul>
              </div>
            </div>
          </>
        )}
        <DialogFooter>
          <ButtonWithSpinner
            type="submit"
            className="w-full"
            disabled={!selectedMember || isAddingMember}
            onClick={() => handleAddMember()}
            isLoading={isAddingMember}
          >
            {!selectedMember
              ? "Select a user"
              : `Add ${selectedMember.name} to organization`}
          </ButtonWithSpinner>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
