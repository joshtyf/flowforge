import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"

interface ApproveConfirmationDialogProps {
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
  onApprove: () => void
}

export function ApproveConfirmationDialog({
  open,
  setOpen,
  onApprove,
}: ApproveConfirmationDialogProps) {
  return (
    <AlertDialog open={open} onOpenChange={setOpen}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Approve Service Request</AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure? This action is irreversible and once approved, the
            pipeline will start for the user.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction
            className={
              "bg-success text-success-foreground hover:bg-success hover:opacity-90"
            }
            onClick={onApprove}
          >
            Approve
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
