import Spinner from "@/components/layouts/spinner"
import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogContent,
} from "@/components/ui/dialog"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"

interface CreateOrgFormDialogProps {
  handleCreateOrg: (name: string) => void
  open: boolean
  setOpen: React.Dispatch<React.SetStateAction<boolean>>
  createOrgloading: boolean
}

const createOrgformSchema = z.object({
  // Following GitHub rules for org name max characters: https://github.com/dead-claudia/github-limits?tab=readme-ov-file#organization-names
  orgName: z
    .string()
    .max(39, {
      message: "Organization name can only have a maximum of 39 characters.",
    })
    .min(1, "Organization name is required"),
})

export default function CreateOrgFormDialog({
  handleCreateOrg,
  open,
  setOpen,
  createOrgloading,
}: CreateOrgFormDialogProps) {
  const form = useForm<z.infer<typeof createOrgformSchema>>({
    resolver: zodResolver(createOrgformSchema),
    defaultValues: {
      orgName: "",
    },
  })

  const onSubmit = ({ orgName }: z.infer<typeof createOrgformSchema>) => {
    handleCreateOrg(orgName)
  }
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <div className="space-y-10">
              <DialogHeader>
                <DialogTitle>Create New Organization</DialogTitle>
              </DialogHeader>
              <FormField
                control={form.control}
                name="orgName"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Organization Name</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <DialogFooter>
                <Button onClick={() => setOpen(false)} variant="outline">
                  Cancel
                </Button>
                <Button type="submit" disabled={createOrgloading}>
                  {createOrgloading ? <Spinner /> : "Create"}
                </Button>
              </DialogFooter>
            </div>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}
