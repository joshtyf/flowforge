import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"

import { z } from "zod"

interface UseRejectFormOptions {
  onReject: (remarks?: string) => void
  closeDialog: () => void
}

const RejectFormSchema = z.object({
  remarks: z.string().optional(),
})

const useRejectForm = ({ onReject, closeDialog }: UseRejectFormOptions) => {
  const rejectForm = useForm<z.infer<typeof RejectFormSchema>>({
    resolver: zodResolver(RejectFormSchema),
    mode: "onChange",
    resetOptions: { keepDefaultValues: true },
  })
  function onSubmit(values: z.infer<typeof RejectFormSchema>) {
    onReject(values.remarks)
    rejectForm.reset()
    closeDialog()
  }

  return {
    form: rejectForm,
    onSubmit,
  }
}

export default useRejectForm
