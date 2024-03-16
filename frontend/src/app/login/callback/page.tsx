"use client"

import { useCountdown } from "@/lib/hooks"
import { useRouter } from "next/navigation"
import useAuthentication from "./_hooks/use-auth"

export default function LoginCallbackPage() {
  useAuthentication()

  const TIMEOUT_DURATION = 5
  const router = useRouter()
  const { countdown } = useCountdown(TIMEOUT_DURATION, () =>
    router.replace("/")
  )
  return <div>Successfully logged in! Redirecting in {countdown}s</div>
}
