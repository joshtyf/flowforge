import { cn } from "@/lib/utils"
import React from "react"

interface HeaderAccessoryProp {
  className?: string
}
export default function HeaderAccessory({ className }: HeaderAccessoryProp) {
  return <div className={cn("h-[5px] w-[60px] bg-black", className)} />
}
