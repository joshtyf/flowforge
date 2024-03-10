import { getUserProfile } from "@/lib/auth0"
import { UserProfile } from "@/types/user-profile"
import { useEffect, useState } from "react"

interface UseUserProfileOptions {
  accessToken: string
}

const useUserProfile = ({ accessToken }: UseUserProfileOptions) => {
  const [userProfile, setUserProfile] = useState<UserProfile>({})

  useEffect(() => {
    getUserProfile(accessToken).then((userProfile) =>
      setUserProfile(userProfile)
    )
  }, [accessToken])
  return { userProfile }
}

export default useUserProfile
