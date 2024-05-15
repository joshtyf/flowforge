import { Skeleton } from "@/components/ui/skeleton"
import useUserInfo from "@/hooks/use-user-info"

interface CreatedByInfoProps {
  userId: string
}

export default function CreatedByInfo({ userId }: CreatedByInfoProps) {
  const { user, isUserInfoLoading } = useUserInfo(userId)

  return isUserInfoLoading ? (
    <Skeleton className="w-28 h-5" />
  ) : (
    <p>{user?.name ?? "N.A."}</p>
  )
}
