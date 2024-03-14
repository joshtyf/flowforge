"use client"

import { useCountdown } from "@/lib/hooks"
import { useRouter } from "next/navigation"
import { useEffect } from "react"
import useAuthentication from "./_hooks/use-auth"

export default function LoginCallbackPage() {
  useAuthentication()

  const TIMEOUT_DURATION = 5
  const { countdown } = useCountdown(TIMEOUT_DURATION)
  const router = useRouter()
  useEffect(() => {
    if (countdown === 0) {
      router.replace("/")
    }
  }, [router, countdown])
  return <div>Successfully logged in! Redirecting in {countdown}s</div>
}
