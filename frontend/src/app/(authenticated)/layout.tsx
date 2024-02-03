"use client"

import Navbar from "@/components/layouts/navbar"
import Sidebar from "@/components/layouts/sidebar"
import { Toaster } from "@/components/ui/toaster"
import React, { ReactNode, useState } from "react"

interface AuthenticatedLayoutProps {
  children: ReactNode
}

export default function AuthenticatedLayout({
  children,
}: AuthenticatedLayoutProps) {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true)

  const toggleSidebar = () => {
    setIsSidebarOpen((isSidebarOpen) => !isSidebarOpen)
  }

  return (
    <>
      <div className="flex flex-row w-full min-h-[100vh]">
        <Sidebar
          className={` ${isSidebarOpen ? "min-w-[280px] w-[280px]" : "min-w-0 w-0"} overflow-x-hidden transition-width duration-300 ease-in-out`}
        />
        <div className="w-full">
          <Navbar toggleSidebar={toggleSidebar} />
          <div className="w-full h-full max-h-[90vh] flex justify-center items-center flex-col relative">
            <div className="w-5/6 h-full relative">{children}</div>
          </div>
        </div>
        <Toaster />
      </div>
    </>
  )
}
