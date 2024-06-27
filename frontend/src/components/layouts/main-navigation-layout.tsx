"use client"

import Navbar from "@/components/layouts/navbar"
import Sidebar from "@/components/layouts/sidebar"
import { useCurrentUserInfo } from "@/contexts/current-user-info-context"
import { ReactNode, useState } from "react"

interface MainNavigationLayoutProps {
  children: ReactNode
  enableOrgName?: boolean
}

export default function MainNavigationLayout({
  children,
  enableOrgName = true,
}: MainNavigationLayoutProps) {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true)
  const toggleSidebar = () => {
    setIsSidebarOpen((isSidebarOpen) => !isSidebarOpen)
  }

  const userInfo = useCurrentUserInfo()

  return (
    <div
      className="flex flex-row w-full min-h-[100vh]"
      suppressHydrationWarning
    >
      <Sidebar
        className={` ${isSidebarOpen ? "min-w-[300px] w-[300px]" : "min-w-0 w-0"} overflow-x-hidden transition-width duration-300 ease-in-out`}
      />
      <div className="w-full h-full">
        <Navbar
          toggleSidebar={toggleSidebar}
          username={userInfo?.name ?? ""}
          enableOrgName={enableOrgName}
        />
        <div className="w-full h-full flex justify-center items-center flex-col relative">
          <div className="w-5/6 h-full relative">{children}</div>
        </div>
      </div>
    </div>
  )
}
