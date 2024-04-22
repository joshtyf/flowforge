"use client"

import HeaderAccessory from "@/components/ui/header-accessory"
import useServiceRequests from "./_hooks/use-service-requests"
import { DataTable } from "@/components/data-table/data-table"
import { columns } from "./columns"

export default function ServiceRequestDashboardPage() {
  const { serviceRequests } = useServiceRequests()
  return (
    <div className="flex flex-col justify-start py-10">
      <HeaderAccessory />
      <div className="flex items-baseline">
        <p className="font-bold text-3xl pt-5">Your Service Requests</p>
      </div>
      <div className="py-10">
        <DataTable columns={columns} data={serviceRequests} />
      </div>
    </div>
  )
}