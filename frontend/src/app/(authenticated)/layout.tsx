"use client"

import Navbar from "@/components/layouts/navbar"
import Sidebar from "@/components/layouts/sidebar"
import { Toaster } from "@/components/ui/toaster"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { getCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { ReactNode, useEffect, useMemo, useState } from "react"
import useUserProfile from "./_hooks/use-user-profile"

interface AuthenticatedLayoutProps {
  children: ReactNode
}

export default function AuthenticatedLayout({
  children,
}: AuthenticatedLayoutProps) {
  const queryClient = new QueryClient()

  const [isSidebarOpen, setIsSidebarOpen] = useState(true)
  const [render, setRender] = useState(false)
  const toggleSidebar = () => {
    setIsSidebarOpen((isSidebarOpen) => !isSidebarOpen)
  }

  const router = useRouter()
  const isLoggedIn = useMemo(() => {
    return getCookie("access_token") ? true : false
  }, [])
  useEffect(() => {
    if (!isLoggedIn) {
      router.push("/login")
    } else {
      // To resolve Hydration UI mismatch issue
      setRender(true)
    }
  }, [isLoggedIn, router])

  const { userProfile } = useUserProfile({
    accessToken: getCookie("access_token") as string,
  })

  return (
    <QueryClientProvider client={queryClient}>
      {render && (
        <div
          className="flex flex-row w-full min-h-[100vh]"
          suppressHydrationWarning
        >
          <Sidebar
            className={` ${isSidebarOpen ? "min-w-[280px] w-[280px]" : "min-w-0 w-0"} overflow-x-hidden transition-width duration-300 ease-in-out`}
          />
          <div className="w-full">
            <Navbar
              toggleSidebar={toggleSidebar}
              username={userProfile.nickname ?? ""}
            />
            <div className="w-full h-full max-h-[90vh] flex justify-center items-center flex-col relative">
              <div className="w-5/6 h-full relative">{children}</div>
            </div>
          </div>
          <Toaster />
        </div>
      )}
    </QueryClientProvider>
  )
}
