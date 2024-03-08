"use client"

import { setCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { useEffect } from "react"

export default function LoginCallbackPage() {
  const router = useRouter()
  const TIMEOUT_DURATION = 5

  useEffect(() => {
    const interval = setTimeout(() => {
      console.log("redirecting")
      router.replace("/")
    }, TIMEOUT_DURATION * 1000)
    const hash = window.location.hash.substring(1)
    const search = new URLSearchParams(hash)
    const accessToken = search.get("access_token")
    const expiresIn = search.get("expires_in")
    const tokenType = search.get("token_type")
    setCookie("loggedIn", "true")

    return () => clearTimeout(interval)
  }, [router])
  return <div>Successfully logged in! Redirecting in {TIMEOUT_DURATION}s</div>
}
