"use client"

import { UserContext } from "@/components/contexts/user"
import Navbar from "@/components/layouts/navbar"
import Sidebar from "@/components/layouts/sidebar"
import { Toaster } from "@/components/ui/toaster"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { ReactNode, useContext, useState } from "react"

interface MainLayoutProps {
  children: ReactNode
}

export default function MainLayout({ children }: MainLayoutProps) {
  const queryClient = new QueryClient()

  const [isSidebarOpen, setIsSidebarOpen] = useState(true)
  const toggleSidebar = () => {
    setIsSidebarOpen((isSidebarOpen) => !isSidebarOpen)
  }
  const userProfile = useContext(UserContext)

  return (
    <QueryClientProvider client={queryClient}>
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
    </QueryClientProvider>
  )
}
