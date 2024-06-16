import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"
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
  const [openAlertDialog, setOpenAlertDialog] = useState(false)

  useEffect(() => {
    login()
      .then((res) => setUserInfo(res))
      .catch((err) => {
        if (err.response?.status === 404) {
          // New user detected, prompt to create user in Flowforge
          setOpenAlertDialog(true)
          console.log("New user detected")
        } else {
          // Unexpected error
          console.error(err)
        }
      })
  }, [])

  return (
    <CurrentUserInfoContext.Provider value={userInfo}>
      {children}
      <AlertDialog open={openAlertDialog}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Welcome to Flowforge!</AlertDialogTitle>
            <AlertDialogDescription>
              Since its your first time using Flowforge, we need you to create a
              username for your account.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogAction>Create</AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </CurrentUserInfoContext.Provider>
  )
}

export function useCurrentUserInfo() {
  const context = useContext(CurrentUserInfoContext)

  return context
}
