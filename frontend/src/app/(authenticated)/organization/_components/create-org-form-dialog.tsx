import Spinner from "@/components/layouts/spinner"
import { Button, ButtonWithSpinner } from "@/components/ui/button"
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
import { FieldValues, UseFormReturn } from "react-hook-form"
import { z } from "zod"
import { createOrgformSchema } from "../_hooks/use-create-organization-form"

interface CreateOrgFormDialogProps {
  form: UseFormReturn<{ orgName: string }, object, undefined>
  openFormDialog: boolean
  setOpenFormDialog: (open: boolean) => void
  createOrgLoading: boolean
  handleCreateOrg: (name: string) => void
}

export default function CreateOrgFormDialog({
  form,
  openFormDialog,
  setOpenFormDialog,
  createOrgLoading,
  handleCreateOrg,
}: CreateOrgFormDialogProps) {
  const onFormSubmit = ({ orgName }: z.infer<typeof createOrgformSchema>) => {
    handleCreateOrg(orgName)
  }

  return (
    <Dialog open={openFormDialog} onOpenChange={setOpenFormDialog}>
      <DialogContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onFormSubmit)}>
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
                <Button
                  onClick={() => setOpenFormDialog(false)}
                  variant="outline"
                >
                  Cancel
                </Button>
                <ButtonWithSpinner type="submit" disabled={createOrgLoading}>
                  {createOrgLoading ? <Spinner /> : "Create"}
                </ButtonWithSpinner>
              </DialogFooter>
            </div>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}
