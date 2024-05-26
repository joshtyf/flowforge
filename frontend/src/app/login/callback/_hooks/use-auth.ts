import { setCookie } from "cookies-next"
import { useEffect } from "react"

interface UseAuthenticationOptions {}

const useAuthentication = () => {
  useEffect(() => {
    const hash = window.location.hash.substring(1)
    const search = new URLSearchParams(hash)
    const accessToken = search.get("access_token")
    // Set default to 7200 seconds (2 hours)
    const expiresIn: string = search.get("expires_in") ?? "7200"
    setCookie("logged_in", "true")
    setCookie("access_token", accessToken, {
      maxAge: Number(expiresIn),
    })
  }, [])
}

export default useAuthentication
