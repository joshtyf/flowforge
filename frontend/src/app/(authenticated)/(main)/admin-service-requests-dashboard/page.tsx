"use client"

import HeaderAccessory from "@/components/ui/header-accessory"
import useOrgServiceRequests from "./_hooks/use-org-service-requests"
import { orgServiceRequestColumns } from "./columns"
import { DataTable } from "@/components/data-table/data-table"

export default function ApproveServiceRequestPage() {
  const { serviceRequests } = useOrgServiceRequests()
  return (
    <div className="flex flex-col justify-start py-10">
      <HeaderAccessory />
      <div className="flex items-baseline">
        <p className="font-bold text-3xl pt-5">
          Admin Service Request Dashboard
        </p>
      </div>
      <div className="py-10">
        <DataTable columns={orgServiceRequestColumns} data={serviceRequests} />
      </div>
    </div>
  )
}
