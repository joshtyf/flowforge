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

export default function ServiceCatalogPage() {
  const { services } = useServices()

  return (
    <div className="w-full flex justify-center items-center flex-col">
      <div className="w-5/6">
        <div className=" flex flex-col justify-start pt-10">
          <HeaderAccessory />
          <div className="flex items-baseline">
            <p className="font-bold text-3xl pt-5">Service Catalog</p>
            <Button className="ml-auto">Create Service</Button>
          </div>
        </div>
        <div className="pt-10 grid grid-cols-auto-fill-min-20 gap-y-10 ">
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
      </div>
    </div>
  )
}
