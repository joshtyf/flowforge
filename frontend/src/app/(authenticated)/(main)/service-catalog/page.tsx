"use client"

import React from "react"
import useServices from "./_hooks/use-services"

import { Button } from "@/components/ui/button"
import HeaderAccessory from "@/components/ui/header-accessory"

import { useRouter } from "next/navigation"
import ServicesSkeletonView from "./_views/services-skeleton-view"
import ServicesView from "./_views/services-view"

export default function ServiceCatalogPage() {
  const { services, isServicesLoading } = useServices()
  const router = useRouter()

  return (
    <>
      <div className="flex flex-col justify-start py-10">
        <HeaderAccessory />
        <div className="flex items-baseline">
          <p className="font-bold text-3xl pt-5">Service Catalog</p>
          <Button
            className="ml-auto"
            onClick={() => router.push("/service-catalog/create-service")}
          >
            Create Service
          </Button>
        </div>
      </div>
      {isServicesLoading ? (
        <ServicesSkeletonView />
      ) : (
        <ServicesView services={services} router={router} />
      )}
    </>
  )
}
