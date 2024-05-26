"use client"

import NotFoundPage from "@/app/not-found"
import MainNavigationLayout from "@/components/layouts/main-navigation-layout"
import { useUserMemberships } from "@/contexts/user-memberships-context"
import { ReactNode, useState } from "react"

interface AdminLayoutProps {
  children: ReactNode
}

export default function AdminLayout({ children }: AdminLayoutProps) {
  const { isAdmin } = useUserMemberships()

  return isAdmin ? (
    <MainNavigationLayout>{children}</MainNavigationLayout>
  ) : (
    <NotFoundPage />
  )
}
