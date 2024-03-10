"use client"

import useAuth from "./_hooks/use-auth"

export default function LoginCallbackPage() {
  const TIMEOUT_DURATION = 5
  const { countdown } = useAuth({ timeoutDuration: TIMEOUT_DURATION })
  return <div>Successfully logged in! Redirecting in {countdown}s</div>
}
