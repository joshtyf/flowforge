type Auth0UserProfile = {
  email?: string
  name?: string
  preferred_username?: string
  nickname?: string
}

type UserInfo = {
  user_id: string
  name: string
  identity_provider: string
  created_on: Date
  deleted: boolean
  email: string
}

export type { Auth0UserProfile, UserInfo }
