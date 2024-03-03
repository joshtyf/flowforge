"use client"

import HeaderAccessory from "@/components/ui/header-accessory"

export default function ApproveServiceRequestPage() {
  return (
    <>
      <div className="flex flex-col justify-start py-10">
        <HeaderAccessory />
        <div className="flex items-baseline">
          <p className="font-bold text-3xl pt-5">Pending Service Requests</p>
        </div>
      </div>
    </>
  )
}
