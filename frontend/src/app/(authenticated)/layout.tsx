"use client"

import { UserContext } from "@/components/contexts/user"
import Navbar from "@/components/layouts/navbar"
import { getUserProfile } from "@/lib/auth0"
import { UserProfile } from "@/types/user-profile"
import { getCookie, hasCookie, setCookie } from "cookies-next"
import { useRouter } from "next/navigation"
import { ReactNode, useEffect, useMemo, useState } from "react"

interface AuthenticatedLayoutProps {
  children: ReactNode
}

export default function AuthenticatedLayout({
  children,
}: AuthenticatedLayoutProps) {
  const [render, setRender] = useState(false)
  const [userProfile, setUserProfile] = useState<UserProfile>({})
  const router = useRouter()
  const isLoggedIn = useMemo(() => {
    return getCookie("access_token") ? true : false
  }, [])
  useEffect(() => {
    if (!isLoggedIn) {
      router.push("/login")
    } else {
      // TODO: Extract to use context to store user profile
      getUserProfile(getCookie("access_token") as string).then(
        (userProfile) => {
          setUserProfile(userProfile)
        }
      )
      // To resolve Hydration UI mismatch issue
      setRender(true)
    }
  }, [isLoggedIn, router, setUserProfile])

  // TODO: Replace with actual organisations
  const orgs = [
    {
      id: 1,
      name: "Organisation 1",
    },
    {
      id: 2,
      name: "Organisation 2",
    },
  ]

  return (
    <UserContext.Provider value={userProfile}>
      {/* TODO: refactor into separate component */}
      {!hasCookie("org_id") ? (
        <div className="">
          <Navbar username={userProfile.nickname ?? ""} />
          <div className="mt-20 flex flex-col justify-center items-center">
            <p className="mb-8 text-2xl">Your Organisations</p>
            <div className="border rounded-md w-2/5">
              <ul className="divide-y divide-slate-200">
                {orgs.map((org) => (
                  <li
                    key={org.id}
                    className="px-8 py-4 cursor-pointer text-xl hover:text-blue-500"
                    onClick={(e) => {
                      setCookie("org_id", org.id)
                      router.refresh() // Refresh the page to re-render
                    }}
                  >
                    {org.name}
                  </li>
                ))}
              </ul>
            </div>
          </div>
        </div>
      ) : (
        render && children
      )}
    </UserContext.Provider>
  )
}
