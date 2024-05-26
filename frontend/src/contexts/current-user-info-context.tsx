import { login } from "@/lib/service"
import { UserInfo } from "@/types/user-profile"
import { createContext, useContext, useEffect, useState } from "react"

const CurrentUserInfoContext = createContext<UserInfo | undefined>(undefined)

export function CurrentUserInfoContextProvider({
  children,
}: {
  children: React.ReactNode
}) {
  const [userInfo, setUserInfo] = useState<UserInfo>()

  useEffect(() => {
    login()
      .then((res) => setUserInfo(res))
      .catch((err) => console.error(err))
  }, [])

  return (
    <CurrentUserInfoContext.Provider value={userInfo}>
      {children}
    </CurrentUserInfoContext.Provider>
  )
}

export function useCurrentUserInfo() {
  const context = useContext(CurrentUserInfoContext)

  return context
}
