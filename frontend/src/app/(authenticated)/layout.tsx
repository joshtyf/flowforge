"use client"

import { CurrentUserInfoContextProvider } from "@/contexts/current-user-info-context"
import apiClient from "@/lib/apiClient"
import { getCookie, hasCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { ReactNode, useEffect, useState } from "react"

interface AuthenticatedLayoutProps {
  children: ReactNode
}

export default function AuthenticatedLayout({
  children,
}: AuthenticatedLayoutProps) {
  const router = useRouter()
  const [render, setRender] = useState(false)
  useEffect(() => {
    if (!getCookie("logged_in") || !hasCookie("access_token")) {
      router.push("/login")
    } else {
      apiClient.defaults.headers.Authorization = `Bearer ${getCookie("access_token") as string}`
      setRender(true)
    }
    apiClient.defaults.headers.Authorization = `Bearer ${getCookie("access_token") as string}`
  }, [router])
  return (
    render && (
      <CurrentUserInfoContextProvider>
        {children}
      </CurrentUserInfoContextProvider>
    )
  )
}
