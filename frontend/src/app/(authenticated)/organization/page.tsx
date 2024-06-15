"use client"

import { Skeleton } from "@/components/ui/skeleton"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"
import { getAllOrgsForUser } from "@/lib/service"
import { Organization } from "@/types/organization"
import { setCookie } from "cookies-next"
import { Plus } from "lucide-react"
import { useRouter } from "next/navigation"
import { useEffect, useState } from "react"

export default function OrganizationsPage() {
  const [organizations, setOrganizations] = useState<Organization[]>([])
  const [loading, setLoading] = useState(true)

  const fetchOrgs = () => {
    getAllOrgsForUser().then((orgs) => {
      setOrganizations(orgs)
      setLoading(false)
    })
  }

  useEffect(() => {
    fetchOrgs()
  }, [])

  const router = useRouter()
  return (
    <div className="mt-20 flex flex-col justify-center items-center">
      <p className="mb-8 text-2xl">Your Organizations</p>
      {loading ? (
        <div className="space-y-4 w-2/5">
          <Skeleton className={"h-12 rounded-md"} />
          <Skeleton className={"h-12 rounded-md"} />
          <Skeleton className={"h-12 rounded-md"} />
        </div>
      ) : (
        <div className="border rounded-md w-2/5">
          <ul className="divide-y divide-slate-200">
            {organizations.map((org) => (
              <li
                key={org.org_id}
                className="px-8 py-4 cursor-pointer text-xl hover:text-blue-500"
                onClick={() => {
                  setCookie("org_id", org.org_id)
                  router.push("/")
                }}
              >
                {org.name}
              </li>
            ))}
            <TooltipProvider>
              <Tooltip delayDuration={300}>
                <TooltipTrigger
                  className="w-full px-8 py-4 cursor-pointer text-xl hover:text-blue-500 flex justify-center items-center"
                  onClick={() => {}}
                >
                  <Plus />
                </TooltipTrigger>
                <TooltipContent>
                  <p>Create New Organization</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </ul>
        </div>
      )}
    </div>
  )
}
