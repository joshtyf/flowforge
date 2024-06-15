"use client"

import { Skeleton } from "@/components/ui/skeleton"
import { setCookie } from "cookies-next"
import { PlusSquare } from "lucide-react"
import { useRouter } from "next/navigation"
import useOrganizations from "./_hooks/use-organizations"
import CreateOrgFormDialog from "./_components/create-org-form-dialog"
import { useState } from "react"
import { Toaster } from "@/components/ui/toaster"

export default function OrganizationsPage() {
  const [openFormDialog, setOpenFormDialog] = useState(false)

  const { organizations, orgsLoading, handleCreateOrg, createOrgloading } =
    useOrganizations({ setOpenFormDialog })

  const router = useRouter()
  return (
    <div className="mt-20 flex flex-col justify-center items-center">
      <p className="mb-8 text-2xl">Your Organizations</p>
      {orgsLoading ? (
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
            <li
              className="w-full px-8 py-4 space-x-3 cursor-pointer text-xl text-gray-400 hover:text-blue-500 flex justify-center items-center"
              onClick={() => setOpenFormDialog(true)}
            >
              <PlusSquare />
              <p>Create new Organization</p>
            </li>
          </ul>
          <CreateOrgFormDialog
            open={openFormDialog}
            setOpen={setOpenFormDialog}
            handleCreateOrg={handleCreateOrg}
            createOrgloading={createOrgloading}
          />
        </div>
      )}
      <Toaster />
    </div>
  )
}
