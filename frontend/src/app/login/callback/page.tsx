"use client"

import useAuthentication from "./_hooks/useAuthentication"

export default function LoginCallbackPage() {
  const TIMEOUT_DURATION = 5

  useAuthentication({ timeout: TIMEOUT_DURATION })
  return <div>Successfully logged in! Redirecting in {TIMEOUT_DURATION}s</div>
}
