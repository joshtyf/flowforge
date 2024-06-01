import { useCountdown } from "@/hooks/use-countdown"
import { useRouter } from "next/navigation"
import React from "react"

interface RequestCreatedTextWithCountdownProps {
  countdownStart?: number
}

export default function RequestCreatedTextWithCountdown({
  countdownStart = 3,
}: RequestCreatedTextWithCountdownProps) {
  const router = useRouter()
  const { countdown } = useCountdown(countdownStart, () => {
    router.push("/your-service-request-dashboard")
  })
  return (
    <p>
      You will be redirected to <strong>Your Service Requests Dashboard</strong>{" "}
      in {countdown}s.
    </p>
  )
}
