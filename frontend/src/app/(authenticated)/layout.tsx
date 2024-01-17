"use client"

import Navbar from "@/components/layouts/navbar"
import Sidebar from "@/components/layouts/sidebar"
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
          className={` ${isSidebarOpen ? "w-96" : "w-0"} overflow-x-hidden transition-width duration-300 ease-in-out`}
        />
        <div className="w-full">
          <Navbar toggleSidebar={toggleSidebar} />
          {children}
        </div>
      </div>
    </>
  )
}
