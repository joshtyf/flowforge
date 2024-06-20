"use client"

import { Toaster } from "@/components/ui/toaster"
import SettingsSidebar from "./_components/settings-sidebar"
import Navbar from "@/components/layouts/navbar"
import { useEffect, useState } from "react"
import { Auth0UserProfile } from "@/types/user-profile"
import { getUserProfile } from "@/lib/auth0"
import { getCookie } from "cookies-next"

export default function SettingsLayout({
  children,
}: {
  children: React.ReactNode
}) {
  const [userProfile, setUserProfile] = useState<Auth0UserProfile>()
  useEffect(() => {
    getUserProfile(getCookie("access_token") as string)
      .then((userProfile) => setUserProfile(userProfile))
      .catch((error) => {
        console.log(error)
      })
  }, [])

  return (
    <>
      <Navbar
        enableSidebarToggle={false}
        enableOrgName={false}
        username={userProfile?.nickname ?? ""}
      />
      <div className="flex justify-center pt-10">
        <div className="w-[80%] flex space-x-10">
          <SettingsSidebar className="w-[30%]" />
          <div className="w-[70%]">{children}</div>
        </div>
      </div>
    </>
  )
}
