import { toast } from "@/components/ui/use-toast"
import { createOrg, getAllOrgsForUser } from "@/lib/service"
import { Organization } from "@/types/organization"
import { useEffect, useState } from "react"

interface UseOrganizationsOptions {
  setOpenFormDialog: React.Dispatch<React.SetStateAction<boolean>>
}

export default function useOrganizations({
  setOpenFormDialog,
}: UseOrganizationsOptions) {
  const [organizations, setOrganizations] = useState<Organization[]>([])
  const [orgsLoading, setOrgsLoading] = useState(false)
  const [createOrgloading, setCreateOrgloading] = useState(false)

  const fetchOrgs = (refetch?: boolean) => {
    if (!refetch) {
      setOrgsLoading(true)
    }
    getAllOrgsForUser().then((orgs) => {
      setOrganizations(orgs)
      if (!refetch) {
        setOrgsLoading(false)
      }
    })
  }

  useEffect(() => {
    fetchOrgs()
  }, [])

  const handleCreateOrg = async (name: string) => {
    setCreateOrgloading(true)
    createOrg(name)
      .then(() => {
        fetchOrgs(true)
        toast({
          title: "Organization Created Successfully",
          description: (
            <p>
              Check the new organization under{" "}
              <strong>Your Organizations</strong>.
            </p>
          ),
        })
        setOpenFormDialog(false)
      })
      .catch((err) => {
        toast({
          variant: "destructive",
          title: "Create Organization Error",
          description: "Could not create organization. Please try again later.",
        })
        console.error(err)
      })
      .finally(() => {
        setCreateOrgloading(false)
      })
  }

  return {
    organizations,
    orgsLoading,
    handleCreateOrg,
    createOrgloading,
  }
}
