enum Role {
  Owner = "Owner",
  Admin = "Admin",
  Member = "Member",
}

type Membership = {
  org_id: number
  role: Role
  joined_on: string
}

type UserMemberships = {
  user_id: string
  memberships: Membership[]
}
export { Role }
export type { Membership, UserMemberships }
