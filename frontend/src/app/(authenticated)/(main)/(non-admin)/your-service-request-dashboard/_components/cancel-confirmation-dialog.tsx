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

interface CancelConfirmationDialogProps {
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
  onCancel: () => void
}

export function CancelConfirmationDialog({
  open,
  setOpen,
  onCancel,
}: CancelConfirmationDialogProps) {
  return (
    <AlertDialog open={open} onOpenChange={setOpen}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Cancel Service Request</AlertDialogTitle>
          <AlertDialogDescription>
            Are you sure? This action is <strong>irreversible</strong>. Do note
            that any step that is currently running will finish executing first
            before the request is cancelled.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction onClick={onCancel}>
            Cancel Request
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
