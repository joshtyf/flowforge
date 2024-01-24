"use client"

import React from "react"
import useServices from "./_hooks/useServices"
import {
  Card,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import HeaderAccessory from "@/components/ui/header-accessory"
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination"
import { useRouter } from "next/navigation"

export default function ServiceCatalogPage() {
  const { services } = useServices()
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
      <div className=" grid grid-cols-auto-fill-min-20 gap-y-10 max-h-[75%] overflow-y-auto">
        {services.map((service) => (
          <div key={service.id} className="flex items-center justify-center">
            <Card className="w-[250px] shadow">
              <CardHeader>
                <CardTitle>{service.name}</CardTitle>
                <CardDescription>{service.description}</CardDescription>
              </CardHeader>
              <CardFooter className="flex justify-end">
                <Button variant="outline">Request</Button>
              </CardFooter>
            </Card>
          </div>
        ))}
      </div>
      <div className="w-full flex justify-center absolute bottom-0">
        <Pagination>
          <PaginationContent>
            <PaginationItem>
              <PaginationPrevious />
            </PaginationItem>
            <PaginationItem>
              <PaginationLink isActive>1</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink>2</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink>3</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationNext />
            </PaginationItem>
          </PaginationContent>
        </Pagination>
      </div>
    </>
  )
}
