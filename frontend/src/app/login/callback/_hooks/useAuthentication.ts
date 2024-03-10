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
      console.log("redirecting")
      router.replace("/")
    }, TIMEOUT_DURATION * 1000)
    const hash = window.location.hash.substring(1)
    const search = new URLSearchParams(hash)
    const accessToken = search.get("access_token")
    const expiresIn = search.get("expires_in")
    const tokenType = search.get("token_type")
    setCookie("loggedIn", "true")

    return () => clearTimeout(timeout)
  }, [router])
  return {}
}

export default useAuthentication
