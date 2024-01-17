"use client"

import Navbar from "@/components/layouts/navbar"
import Sidebar from "@/components/layouts/sidebar"
import React, { ReactNode } from "react"

interface AuthenticatedLayoutProps {
  children: ReactNode
}

export default function AuthenticatedLayout({
  children,
}: AuthenticatedLayoutProps) {
  return (
    <>
      <div className="flex flex-row w-full min-h-[100vh]">
        <Sidebar />
        <div className="w-full">
          <Navbar />
          {children}
        </div>
      </div>
    </>
  )
}
