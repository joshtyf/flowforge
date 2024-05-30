"use client"

import { Toaster } from "@/components/ui/toaster"
import { UserMembershipsProvider } from "@/contexts/user-memberships-context"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { hasCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { ReactNode, useEffect, useState } from "react"

interface MainLayoutProps {
  children: ReactNode
}

export default function MainLayout({ children }: MainLayoutProps) {
  const router = useRouter()
  const [render, setRender] = useState(false)
  useEffect(() => {
    if (!hasCookie("org_id")) {
      router.push("/organization")
    } else {
      setRender(true)
    }
  }, [router])

  const queryClient = new QueryClient()

  return (
    render && (
      <QueryClientProvider client={queryClient}>
        <UserMembershipsProvider>
          {children}
          <Toaster />
        </UserMembershipsProvider>
      </QueryClientProvider>
    )
  )
}
