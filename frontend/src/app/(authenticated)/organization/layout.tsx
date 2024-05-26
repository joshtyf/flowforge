"use client"

import Navbar from "@/components/layouts/navbar"
import { useCurrentUserInfo } from "@/contexts/current-user-info-context"
import { hasCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { ReactNode, useEffect, useState } from "react"

interface AuthenticatedLayoutProps {
  children: ReactNode
}

export default function AuthenticatedLayout({
  children,
}: AuthenticatedLayoutProps) {
  const userInfo = useCurrentUserInfo()

  return (
    <>
      <Navbar username={userInfo?.name ?? ""} enableSidebarToggle={false} />
      {children}
    </>
  )
}
