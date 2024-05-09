"use client"

import HeaderAccessory from "@/components/ui/header-accessory"
import useOrgServiceRequests from "./_hooks/use-org-service-requests"
import { orgServiceRequestColumns } from "./columns"
import { DataTable } from "@/components/data-table/data-table"
import { usePagination } from "@/hooks/use-pagination"

export default function ApproveServiceRequestPage() {
  const { onPaginationChange, pagination } = usePagination()
  const { orgServiceRequestsData, noOfPages } = useOrgServiceRequests({
    page: pagination.pageIndex + 1, // API is 1-based
    pageSize: pagination.pageSize,
  })
  return (
    <div className="flex flex-col justify-start py-10">
      <HeaderAccessory />
      <div className="flex items-baseline">
        <p className="font-bold text-3xl pt-5">
          Admin Service Request Dashboard
        </p>
      </div>
      <div className="py-10">
        <DataTable
          columns={orgServiceRequestColumns}
          data={orgServiceRequestsData?.data}
          pageCount={noOfPages}
          onPaginationChange={onPaginationChange}
          pagination={pagination}
        />
      </div>
    </div>
  )
}
