"use client"

import { getCookie, hasCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { ReactNode, useEffect } from "react"

type AuthenticatedLayoutProps = {
  children: ReactNode
}
export default function AuthenticatedLayout({
  children,
}: AuthenticatedLayoutProps) {
  const router = useRouter()
  useEffect(() => {
    if (!getCookie("loggedIn") || !hasCookie("access_token")) {
      router.push("/login")
    }
  }, [router])
  return children
}
