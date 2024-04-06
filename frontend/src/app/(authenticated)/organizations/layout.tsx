"use client"

import Navbar from "@/components/layouts/navbar"
import { hasCookie } from "cookies-next"
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

  return (
    render && (
      <>
        <Navbar username={"joshua"} />
        {children}
      </>
    )
  )
}
