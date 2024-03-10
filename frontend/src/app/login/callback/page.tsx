"use client"

import { setCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { useEffect } from "react"
import useAuthentication from "./_hooks/useAuthentication"

export default function LoginCallbackPage() {
  const router = useRouter()
  const TIMEOUT_DURATION = 5

  useAuthentication({ timeout: TIMEOUT_DURATION })
  return <div>Successfully logged in! Redirecting in {TIMEOUT_DURATION}s</div>
}
