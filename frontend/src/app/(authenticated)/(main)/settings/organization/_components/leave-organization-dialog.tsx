import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import { buttonVariants } from "@/components/ui/button"
import { Loader2 } from "lucide-react"
import { useState } from "react"

interface LeaveOrganizationDialogProps {
  children: React.ReactNode
  onConfirm: () => Promise<void>
  isOwner?: boolean
}

export default function LeaveOrganizationDialog({
  children,
  onConfirm,
  isOwner,
}: LeaveOrganizationDialogProps) {
  const [isLoading, setIsLoading] = useState(false)
  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>{children}</AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>
            Are you sure you want to leave the organization?
          </AlertDialogTitle>
          <AlertDialogDescription>
            Once you leave, you will lose access to all features available
            within the organization. If you are the{" "}
            <strong>
              <i>Owner</i>
            </strong>
            , ensure that you have transferred your ownership to another user
            before you can leave the organization.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction
            onClick={() => {
              setIsLoading(true)
              onConfirm().finally(() => setIsLoading(false))
            }}
            className={buttonVariants({ variant: "destructive" })}
            disabled={isOwner || isLoading}
          >
            {isLoading && <Loader2 className={"animate-spin mr-2"} />}
            Leave
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
