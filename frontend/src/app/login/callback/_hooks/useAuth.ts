import { setCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { useEffect } from "react"
const TIMEOUT_DURATION = 5

interface UseAuthenticationOptions {
  timeout?: number
}

const useAuthentication = ({ timeout = 5 }: UseAuthenticationOptions) => {
  const router = useRouter()
  useEffect(() => {
    const timeout = setTimeout(() => {
      router.replace("/")
    }, TIMEOUT_DURATION * 1000)
    const hash = window.location.hash.substring(1)
    const search = new URLSearchParams(hash)
    const accessToken = search.get("access_token")
    // Set default to 7200 seconds (2 hours)
    const expiresIn: string = search.get("expires_in") ?? "7200"
    setCookie("loggedIn", "true")
    setCookie("access_token", accessToken, {
      maxAge: Number(expiresIn),
    })

    return () => clearTimeout(timeout)
  }, [router])
  return {}
}

export default useAuthentication
