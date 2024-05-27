import { useCountdown } from "@/hooks/use-countdown"
import { useRouter } from "next/navigation"
import React from "react"

interface RequestCreatedTextWithCountdownProps {
  countdownStart?: number
}

export default function PipelineCreatedTextWithCountdown({
  countdownStart = 3,
}: RequestCreatedTextWithCountdownProps) {
  const router = useRouter()
  const { countdown } = useCountdown(countdownStart, () => {
    router.push("/service-catalog")
  })
  return (
    <p>
      You will be redirected to <strong>Service Catalog</strong> in {countdown}
      s.
    </p>
  )
}
