"use client"

import HeaderAccessory from "@/components/ui/header-accessory"
import { Textarea } from "@/components/ui/textarea"
import React from "react"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"
import { ChevronLeft } from "lucide-react"

export default function CreateServicePage() {
  const router = useRouter()

  return (
    <>
      <div className="flex flex-col justify-start py-10">
        <HeaderAccessory />
        <div className="flex items-baseline space-x-2">
          <Button size="icon" variant="ghost" onClick={() => router.back()}>
            <ChevronLeft />
          </Button>
          <p className="font-bold text-3xl pt-5">Create Service</p>
        </div>
      </div>
      <div className="flex flex-col justify-center items-center w-full space-y-10">
        <Textarea
          placeholder="Insert Service JSON Object..."
          className="w-3/5 h-[500px] bg-secondary"
        />
        <div className="w-3/5 flex justify-end">
          <Button size="lg">Submit</Button>
        </div>
      </div>
    </>
  )
}
