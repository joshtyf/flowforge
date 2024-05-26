"use client"

import MainNavigationLayout from "@/components/layouts/main-navigation-layout"
import Navbar from "@/components/layouts/navbar"
import Sidebar from "@/components/layouts/sidebar"
import { Toaster } from "@/components/ui/toaster"
import { UserMembershipsProvider } from "@/context/user-memberships-context"
import { getUserProfile } from "@/lib/auth0"
import { Auth0UserProfile } from "@/types/user-profile"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { getCookie, hasCookie } from "cookies-next"
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

  const [isSidebarOpen, setIsSidebarOpen] = useState(true)
  const toggleSidebar = () => {
    setIsSidebarOpen((isSidebarOpen) => !isSidebarOpen)
  }

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
