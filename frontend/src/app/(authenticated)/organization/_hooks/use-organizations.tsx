import { toast } from "@/components/ui/use-toast"
import { createOrg, getAllOrgsForUser } from "@/lib/service"
import { Organization } from "@/types/organization"
import { useEffect, useState } from "react"

export default function useOrganizations() {
  const [organizations, setOrganizations] = useState<Organization[]>([])
  const [orgsLoading, setOrgsLoading] = useState(false)

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

  return {
    organizations,
    orgsLoading,
    refetchOrgs: () => fetchOrgs(true),
  }
}
