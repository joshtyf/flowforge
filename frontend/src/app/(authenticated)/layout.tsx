"use client"

import { UserContext } from "@/components/contexts/user"
import { getUserProfile } from "@/lib/auth0"
import { UserProfile } from "@/types/user-profile"
import { getCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { ReactNode, createContext, useEffect, useMemo, useState } from "react"

interface AuthenticatedLayoutProps {
  children: ReactNode
}

export default function AuthenticatedLayout({
  children,
}: AuthenticatedLayoutProps) {
  const [render, setRender] = useState(false)
  const [userProfile, setUserProfile] = useState<UserProfile>({})
  const router = useRouter()
  const isLoggedIn = useMemo(() => {
    return getCookie("access_token") ? true : false
  }, [])
  useEffect(() => {
    if (!isLoggedIn) {
      router.push("/login")
    } else {
      // TODO: Extract to use context to store user profile
      getUserProfile(getCookie("access_token") as string).then(
        (userProfile) => {
          setUserProfile(userProfile)
        }
      )
      // To resolve Hydration UI mismatch issue
      setRender(true)
    }
  }, [isLoggedIn, router, setUserProfile])

  return (
    <UserContext.Provider value={userProfile}>
      {render && children}
    </UserContext.Provider>
  )
}
