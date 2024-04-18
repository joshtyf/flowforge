type UserProfile = {
  email?: string
  name?: string
  preferred_username?: string
  nickname?: string
}

// TODO: This type should be moved to a shared location
type UserFromBackend = {
  user_id: string
  name: string
  identity_provider: string
  created_on: Date
  deleted: boolean
}

export type { UserProfile, UserFromBackend }
