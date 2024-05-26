"use client"

import Navbar from "@/components/layouts/navbar"
import Sidebar from "@/components/layouts/sidebar"
import { Toaster } from "@/components/ui/toaster"
import { useCurrentUserInfo } from "@/contexts/current-user-info-context"
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
  const userInfo = useCurrentUserInfo()

  return (
    render && (
      <QueryClientProvider client={queryClient}>
        <div
          className="flex flex-row w-full min-h-[100vh]"
          suppressHydrationWarning
        >
          <Sidebar
            className={` ${isSidebarOpen ? "min-w-[300px] w-[300px]" : "min-w-0 w-0"} overflow-x-hidden transition-width duration-300 ease-in-out`}
          />
          <div className="w-full">
            <Navbar
              toggleSidebar={toggleSidebar}
              username={userInfo?.name ?? ""}
            />
            <div className="w-full h-full max-h-[90vh] flex justify-center items-center flex-col relative">
              <div className="w-5/6 h-full relative">{children}</div>
            </div>
          </div>
          <Toaster />
        </div>
      </QueryClientProvider>
    )
  )
}
