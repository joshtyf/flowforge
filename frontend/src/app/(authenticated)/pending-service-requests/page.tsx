"use client"

import HeaderAccessory from "@/components/ui/header-accessory"
import usePendingServiceRequest from "./_hooks/use-pending-service-requests"
import { columns } from "./columns"
import { DataTable } from "@/components/layouts/data-table"

export default function ApproveServiceRequestPage() {
  const { serviceRequests } = usePendingServiceRequest()
  return (
    <div className="flex flex-col justify-start py-10">
      <HeaderAccessory />
      <div className="flex items-baseline">
        <p className="font-bold text-3xl pt-5">Pending Service Requests</p>
      </div>
      <div className="py-10">
        <DataTable columns={columns} data={serviceRequests} />
      </div>
    </div>
  )
}
