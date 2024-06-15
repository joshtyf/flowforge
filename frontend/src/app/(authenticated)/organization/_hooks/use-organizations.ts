import { getAllOrgsForUser } from "@/lib/service"
import { Organization } from "@/types/organization"
import { useEffect, useState } from "react"

export default function useOrganizations() {
  const [organizations, setOrganizations] = useState<Organization[]>([])
  const [loading, setLoading] = useState(false)

  const fetchOrgs = () => {
    setLoading(true)
    getAllOrgsForUser().then((orgs) => {
      setOrganizations(orgs)
      setLoading(false)
    })
  }

  useEffect(() => {
    fetchOrgs()
  }, [])

  return {
    organizations,
    loadingOrgs: loading,
    refetchOrgs: fetchOrgs,
  }
}
