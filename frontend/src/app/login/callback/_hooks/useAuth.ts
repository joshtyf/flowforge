import { setCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { useCallback, useEffect, useRef, useState } from "react"

interface UseAuthenticationOptions {
  timeoutDuration?: number
}

const useAuthentication = ({
  timeoutDuration = 5,
}: UseAuthenticationOptions) => {
  const [countdown, setCountdown] = useState(timeoutDuration)
  const router = useRouter()
  useEffect(() => {
    const hash = window.location.hash.substring(1)
    const search = new URLSearchParams(hash)
    const accessToken = search.get("access_token")
    // Set default to 7200 seconds (2 hours)
    const expiresIn: string = search.get("expires_in") ?? "7200"
    setCookie("loggedIn", "true")
    setCookie("access_token", accessToken, {
      maxAge: Number(expiresIn),
    })
    const interval = setInterval(() => {
      if (countdown > 0) {
        setCountdown(countdown - 1)
      } else {
        router.replace("/")
      }
    }, 1000)

    return () => clearInterval(interval)
  }, [countdown, router])

  return { countdown }
}

export default useAuthentication
