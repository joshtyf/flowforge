"use client"

import React from "react"
import { useParams, useRouter } from "next/navigation"
import HeaderAccessory from "@/components/ui/header-accessory"
import { Button } from "@/components/ui/button"
import { ChevronLeft } from "lucide-react"
import useServiceRequest from "./_hooks/useServiceRequest"

export default function ServiceRequestPage() {
  const { serviceRequestId } = useParams()
  const serviceRequestIdString = Array.isArray(serviceRequestId)
    ? serviceRequestId[0]
    : serviceRequestId
  const router = useRouter()
  const { serviceRequest } = useServiceRequest({
    serviceRequestId: serviceRequestIdString,
  })
  const { name, description } = serviceRequest
  return (
    <>
      <div className="flex flex-col justify-start py-10">
        <HeaderAccessory />
        <div className="flex items-baseline space-x-2">
          <Button size="icon" variant="ghost" onClick={() => router.back()}>
            <ChevronLeft />
          </Button>
          <p className="font-bold text-3xl pt-5">{name}</p>
        </div>
        <p className="text-lg pt-3 ml-12 text-gray-500">{description}</p>
      </div>
    </>
  )
}
