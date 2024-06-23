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

interface LeaveOrganizationDialogProps {
  children: React.ReactNode
  onConfirm: () => void
  isOwner?: boolean
}

export default function LeaveOrganizationDialog({
  children,
  onConfirm,
  isOwner,
}: LeaveOrganizationDialogProps) {
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
            onClick={onConfirm}
            className={buttonVariants({ variant: "destructive" })}
            disabled={isOwner}
          >
            Leave
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
