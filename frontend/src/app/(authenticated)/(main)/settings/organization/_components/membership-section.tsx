import { PlusSquare } from "lucide-react"
import useMemberships from "../_hooks/use-memberships"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { useState } from "react"
import useDebounce from "@/hooks/use-debounce"
import AddMemberDialog from "./add-member-dialog"

interface MembershipSectionProps {
  organizationId: number
}

export default function MembershipSection({
  organizationId,
}: MembershipSectionProps) {
  const [searchFilter, setSearchFilter] = useState("")

  // Delay filter execution by 0.5s at each filter change
  const { debouncedValue: debouncedFilter } = useDebounce(searchFilter, 500)
  const { members } = useMemberships({
    orgId: organizationId,
    filter: debouncedFilter,
  })
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
        <AddMemberDialog existingMembers={members}>
          <Button variant={"outline"} className="ml-auto">
            Add Member
          </Button>
        </AddMemberDialog>
      </div>
      <div className="border rounded-md">
        <ul className="divide-y divide-slate-200">
          {members.map((member) => (
            <li
              key={member.user_id}
              className="px-4 py-4 flex place-content-between"
            >
              <p>{member.name}</p>
              <p className="text-sm text-muted-foreground">{member.role}</p>
            </li>
          ))}
        </ul>
      </div>
    </div>
  )
}
