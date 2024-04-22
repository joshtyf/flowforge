"use client"

import { Button } from "@/components/ui/button"
import HeaderAccessory from "@/components/ui/header-accessory"
import { ChevronLeft } from "lucide-react"
import { useRouter } from "next/navigation"

interface ServiceRequestLogsViewProps {}

export default function ServiceRequestLogsView({}: ServiceRequestLogsViewProps) {
  const router = useRouter()
  return (
    <>
      <div className="flex flex-col justify-start py-10">
        <HeaderAccessory />
        <div className="flex items-baseline space-x-2 pt-5">
          <Button
            size="icon"
            variant="ghost"
            onClick={() => {
              router.back()
            }}
          >
            <ChevronLeft />
          </Button>

          <p className="font-bold text-3xl">Service Request Logs</p>
        </div>
      </div>
      <div>
        <div className="bg-gray-900 text-green-300 border-none rounded-lg p-3 focus:ring-2 focus:ring-offset-2 focus:ring-green-500 font-mono h-[70vh] overflow-auto">
          Test logs
        </div>
        {/* <Textarea
          disabled
          className="bg-black disabled:opacity-100 text-green-500 border border-black p-4 font-mono"
        ></Textarea> */}
      </div>
    </>
  )
}
