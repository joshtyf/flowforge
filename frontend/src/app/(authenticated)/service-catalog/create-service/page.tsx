"use client"

import React from "react"
import HeaderAccessory from "@/components/ui/header-accessory"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"
import { ChevronLeft } from "lucide-react"
import dynamic from "next/dynamic"
import useCreateService from "./_hooks/useCreateService"

const DynamicReactJson = dynamic(() => import("@microlink/react-json-view"), {
  ssr: false,
})

export default function CreateServicePage() {
  const router = useRouter()

  const { serviceObject, setServiceObject, handleSubmitObject } =
    useCreateService()
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
        <div className="w-3/5 h-[500px] bg-secondary">
          <DynamicReactJson
            src={serviceObject}
            name={false}
            displayDataTypes={false}
            onEdit={(edit) => {
              setServiceObject(edit.updated_src)
            }}
            onAdd={(add) => setServiceObject(add.updated_src)}
            onDelete={(del) => setServiceObject(del.updated_src)}
          />
        </div>
        <div className="w-3/5 flex justify-end">
          <Button onClick={handleSubmitObject} size="lg">
            Submit
          </Button>
        </div>
      </div>
    </>
  )
}
