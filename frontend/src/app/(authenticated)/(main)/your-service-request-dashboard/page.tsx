"use client"

import { DataTable } from "@/components/data-table/data-table"
import HeaderAccessory from "@/components/ui/header-accessory"
import { getPipeline } from "@/lib/service"
import { useEffect, useState } from "react"
import useServiceRequests from "./_hooks/use-service-requests"
import { ColumnsType, columns } from "./columns"

export default function ServiceRequestDashboardPage() {
  const { serviceRequests } = useServiceRequests()
  const [data, setData] = useState<ColumnsType[]>()
  useEffect(() => {
    async function fetchData() {
      if (serviceRequests) {
        const data = await Promise.all(
          serviceRequests.map(async (serviceRequest) => {
            const pipeline = await getPipeline(serviceRequest.pipeline_id)
            return { ...serviceRequest, pipelineName: pipeline.pipeline_name }
          })
        )
        setData(data)
      }
    }
    fetchData()
  }, [serviceRequests])

  return (
    <div className="flex flex-col justify-start py-10">
      <HeaderAccessory />
      <div className="flex items-baseline">
        <p className="font-bold text-3xl pt-5">Your Service Requests</p>
      </div>
      <div className="py-10">
        <DataTable columns={columns} data={data} />
      </div>
    </div>
  )
}
