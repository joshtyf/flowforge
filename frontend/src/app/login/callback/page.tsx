"use client"

import useAuth from "./_hooks/useAuth"

export default function LoginCallbackPage() {
  const TIMEOUT_DURATION = 5

  useAuth({ timeout: TIMEOUT_DURATION })
  return <div>Successfully logged in! Redirecting in {TIMEOUT_DURATION}s</div>
}
