"use client"

import Navbar from "@/components/layouts/navbar"
import { getUserProfile } from "@/lib/auth0"
import { Auth0UserProfile } from "@/types/user-profile"
import { getCookie, hasCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { ReactNode, useEffect, useState } from "react"

interface AuthenticatedLayoutProps {
  children: ReactNode
}

export default function AuthenticatedLayout({
  children,
}: AuthenticatedLayoutProps) {
  const [render, setRender] = useState(false)
  const router = useRouter()
  useEffect(() => {
    if (hasCookie("org_id")) {
      router.push("/")
    } else {
      setRender(true)
    }
  }, [router])

  const [userProfile, setUserProfile] = useState<Auth0UserProfile>()
  useEffect(() => {
    getUserProfile(getCookie("access_token") as string)
      .then((userProfile) => setUserProfile(userProfile))
      .catch((error) => {
        console.log(error)
      })
  })

  return (
    render && (
      <>
        <Navbar
          username={userProfile?.nickname ?? ""}
          enableSidebarToggle={false}
        />
        {children}
      </>
    )
  )
}
