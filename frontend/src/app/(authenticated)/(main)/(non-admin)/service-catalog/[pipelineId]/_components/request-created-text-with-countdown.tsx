import { useCountdown } from "@/hooks/use-countdown"
import React from "react"

interface RequestCreatedTextWithCountdownProps {
  countdownStart: number
  onCountdownComplete: () => void
}

export default function RequestCreatedTextWithCountdown({
  countdownStart,
  onCountdownComplete,
}: RequestCreatedTextWithCountdownProps) {
  const { countdown } = useCountdown(countdownStart, onCountdownComplete)
  return (
    <p>
      You will be redirected to <strong>Your Service Requests</strong> dashboard
      in {countdown}s.
    </p>
  )
}
