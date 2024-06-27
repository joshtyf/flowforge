"use client"

import { Skeleton } from "@/components/ui/skeleton"
import { setCookie } from "cookies-next"
import { PlusSquare } from "lucide-react"
import { useRouter } from "next/navigation"
import useOrganizations from "./_hooks/use-organizations"
import CreateOrgFormDialog from "./_components/create-org-form-dialog"
import { Toaster } from "@/components/ui/toaster"
import useCreateOrganizationForm from "./_hooks/use-create-organization-form"

export default function OrganizationsPage() {
  const { organizations, orgsLoading, refetchOrgs } = useOrganizations()
  const {
    form,
    openFormDialog,
    setOpenFormDialog,
    createOrgLoading,
    handleCreateOrg,
  } = useCreateOrganizationForm({ refetchOrgs })
  const router = useRouter()
  return (
    <div className="mt-20 flex flex-col justify-center items-center">
      <p className="mb-4 text-2xl">Your Organizations</p>
      <p className="mb-4 text-gray-400">
        Please select an organization to access Flowforge features.
      </p>
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
                  setCookie("org_name", org.name)
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
            form={form}
            openFormDialog={openFormDialog}
            setOpenFormDialog={setOpenFormDialog}
            createOrgLoading={createOrgLoading}
            handleCreateOrg={handleCreateOrg}
          />
        </div>
      )}
      <Toaster />
    </div>
  )
}
