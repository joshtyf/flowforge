import { toast } from "@/components/ui/use-toast"
import { getUserById } from "@/lib/service"
import { UserInfo } from "@/types/user-profile"
import { useEffect, useState } from "react"

const useUserInfo = (userId: string) => {
  const [user, setUser] = useState<UserInfo>()
  const [isLoading, setIsLoading] = useState(true)
  useEffect(() => {
    setIsLoading(true)
    getUserById(userId)
      .then((user) => setUser(user))
      .catch((err) => {
        console.error(err)
        toast({
          title: "Fetching User Info Error",
          description: "Failed to fetch user info. Please try again later.",
          variant: "destructive",
        })
      })
      .finally(() => setIsLoading(false))
  }, [userId])

  return { user, isUserInfoLoading: isLoading }
}

export default useUserInfo
