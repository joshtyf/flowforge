"use client"

import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import useUpdateOrgNameForm from "../_hooks/use-update-org-name-form"
import { Input } from "@/components/ui/input"
import { useUserMemberships } from "@/contexts/user-memberships-context"
import { ButtonWithSpinner } from "@/components/ui/button"

interface ChangeOrgNameFormProps {
  organizationId: number
  updateOrgNameInCookie: (name: string) => void
}

export default function ChangeOrgNameForm({
  organizationId,
  updateOrgNameInCookie,
}: ChangeOrgNameFormProps) {
  const { updateOrgNameLoading, form, onFormSubmit } = useUpdateOrgNameForm({
    organizationId,
    updateOrgNameInCookie,
  })

  const { isOwner } = useUserMemberships()
  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onFormSubmit)} className="space-y-5">
        <FormField
          control={form.control}
          name="orgName"
          render={({ field }) => (
            <FormItem className="space-y-2">
              <h2 className="text-lg">Change Organization Name</h2>
              <FormDescription>
                Only the owner of the organization can change the name.
              </FormDescription>
              <FormControl>
                <Input {...field} disabled={!isOwner} className="max-w-md" />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <ButtonWithSpinner
          isLoading={updateOrgNameLoading}
          disabled={updateOrgNameLoading}
        >
          Change
        </ButtonWithSpinner>
      </form>
    </Form>
  )
}
