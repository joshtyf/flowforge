"use client"

import Navbar from "@/components/layouts/navbar"
import Sidebar from "@/components/layouts/sidebar"
import { getUserProfile } from "@/lib/auth0"
import { Auth0UserProfile } from "@/types/user-profile"
import { getCookie } from "cookies-next"
import { ReactNode, useEffect, useState } from "react"

interface MainNavigationLayoutProps {
  children: ReactNode
}

export default function MainNavigationLayout({
  children,
}: MainNavigationLayoutProps) {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true)
  const toggleSidebar = () => {
    setIsSidebarOpen((isSidebarOpen) => !isSidebarOpen)
  }

  const [userProfile, setUserProfile] = useState<Auth0UserProfile>()
  useEffect(() => {
    getUserProfile(getCookie("access_token") as string)
      .then((userProfile) => setUserProfile(userProfile))
      .catch((error) => {
        console.log(error)
      })
  }, [])
  return (
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
          username={userProfile?.nickname ?? ""}
        />
        <div className="w-full h-full max-h-[90vh] flex justify-center items-center flex-col relative">
          <div className="w-5/6 h-full relative">{children}</div>
        </div>
      </div>
    </div>
  )
}
