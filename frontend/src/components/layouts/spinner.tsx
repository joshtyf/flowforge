import { cn } from "@/lib/utils"
import { Loader2 } from "lucide-react"
import React from "react"

interface SpinnerProps extends React.InputHTMLAttributes<HTMLInputElement> {
  size?: number
}

export default function Spinner({ size, ...props }: SpinnerProps) {
  return <Loader2 size={size} className={cn("animate-spin", props.className)} />
}
