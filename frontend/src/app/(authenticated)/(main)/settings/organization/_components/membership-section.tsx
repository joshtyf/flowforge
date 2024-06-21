import { PlusSquare } from "lucide-react"
import useMemberships from "../_hooks/use-memberships"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"

interface MembershipSectionProps {
  organizationId: number
}

export default function MembershipSection({
  organizationId,
}: MembershipSectionProps) {
  const { members } = useMemberships({ orgId: organizationId })
  return (
    <div className="space-y-5">
      <div>
        <h1 className="text-xl">Members</h1>
      </div>
      <div className="flex items-center">
        <Input placeholder="Search for member" className="max-w-xs" />
        <Button variant={"outline"} className="ml-auto">
          Add Member
        </Button>
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
