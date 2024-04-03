import { UserProfile } from "@/types/user-profile"
import { createContext } from "react"

export const UserContext = createContext<UserProfile>({})
