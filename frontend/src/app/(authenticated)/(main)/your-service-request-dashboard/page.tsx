"use client"

import { DataTable } from "@/components/data-table/data-table"
import HeaderAccessory from "@/components/ui/header-accessory"
import useServiceRequests from "./_hooks/use-service-requests"
import { columns } from "./columns"
import { usePagination } from "@/hooks/use-pagination"

export default function ServiceRequestDashboardPage() {
  const { onPaginationChange, pagination } = usePagination()
  const { response } = useServiceRequests({
    page: pagination.pageIndex + 1, // API is 1-based
    pageSize: pagination.pageSize,
  })

  return (
    <div className="flex flex-col justify-start py-10">
      <HeaderAccessory />
      <div className="flex items-baseline">
        <p className="font-bold text-3xl pt-5">Your Service Requests</p>
      </div>
      <div className="py-10">
        <DataTable
          columns={columns}
          data={response?.data}
          onPaginationChange={onPaginationChange}
          pagination={pagination}
        />
      </div>
    </div>
  )
}
