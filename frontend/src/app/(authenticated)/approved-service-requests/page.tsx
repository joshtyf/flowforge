"use client"

import HeaderAccessory from "@/components/ui/header-accessory"
import { approvedServiceRequestColumns } from "./columns"
import { DataTable } from "@/components/layouts/data-table"
import useApprovedServiceRequest from "./_hooks/use-approved-service-requests"

export default function ApproveServiceRequestPage() {
  const { serviceRequests } = useApprovedServiceRequest()
  return (
    <div className="flex flex-col justify-start py-10">
      <HeaderAccessory />
      <div className="flex items-baseline">
        <p className="font-bold text-3xl pt-5">Approved Service Requests</p>
      </div>
      <div className="py-10">
        <DataTable
          columns={approvedServiceRequestColumns}
          data={serviceRequests}
        />
      </div>
    </div>
  )
}
