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
import { Button } from "@/components/ui/button"
import { Loader2 } from "lucide-react"
import { useState } from "react"

interface MemberActionAlertDialogProps {
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
  onConfirm: () => Promise<void>
  title: string
}

export default function MemberActionAlertDialog({
  open,
  setOpen,
  onConfirm,
  title,
}: MemberActionAlertDialogProps) {
  const [isLoading, setIsLoading] = useState(false)

  return (
    <AlertDialog open={open} onOpenChange={setOpen}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>{title}</AlertDialogTitle>
          <AlertDialogDescription>
            This is a permanent action and cannot be undone.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancel</AlertDialogCancel>
          <AlertDialogAction
            onClick={() => {
              setIsLoading(true)
              onConfirm().finally(() => setIsLoading(false))
            }}
            disabled={isLoading}
          >
            {" "}
            {isLoading && <Loader2 className={"animate-spin mr-2"} />}
            Confirm
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
