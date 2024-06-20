"use client"

import MainNavigationLayout from "@/components/layouts/main-navigation-layout"
import SettingsSidebar from "./_components/settings-sidebar"
import Navbar from "@/components/layouts/navbar"
import { useCurrentUserInfo } from "@/contexts/current-user-info-context"

export default function SettingsLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <MainNavigationLayout enableOrgName={false}>
      <div className="flex justify-center pt-10">
        <div className="w-[90%] flex space-x-10">
          <SettingsSidebar className="w-[20%]" />
          <div className="w-[80%]">{children}</div>
        </div>
      </div>
    </MainNavigationLayout>
  )
}
