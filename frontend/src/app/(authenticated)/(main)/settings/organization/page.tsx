"use client"

import { Separator } from "@/components/ui/separator"
import ChangeOrgNameForm from "./_components/change-org-name-form"
import useOrganization from "@/hooks/use-organization"
import MembershipSection from "./_components/membership-section"

export default function OrganizationSettingsPage() {
  const { organizationId, organizationName, updateOrgNameInCookie } =
    useOrganization()
  return (
    <div className="space-y-8">
      {organizationId ? (
        <>
          <h1 className="text-2xl font-bold">
            Organization Settings for <i>{organizationName}</i>
          </h1>
          <Separator className="w-full my-2" />
          <div className="w-3/4">
            <ChangeOrgNameForm
              organizationId={organizationId}
              updateOrgNameInCookie={updateOrgNameInCookie}
            />
          </div>
          <Separator className="w-full my-4" />
          <MembershipSection organizationId={organizationId} />
          <Separator className="w-full my-2" />
        </>
      ) : (
        <div>
          <h1 className="text-2xl font-bold">
            Please select an organization to continue with organization
            settings.
          </h1>
        </div>
      )}
    </div>
  )
}
