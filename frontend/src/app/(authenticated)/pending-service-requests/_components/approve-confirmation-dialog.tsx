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
          <AlertDialogTitle>Are you sure?</AlertDialogTitle>
          <AlertDialogDescription>
            The action cannot be reversed and once approved, the pipeline will
            start for the user.
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
