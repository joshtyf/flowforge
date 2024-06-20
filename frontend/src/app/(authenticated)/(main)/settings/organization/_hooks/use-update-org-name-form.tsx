import { toast } from "@/components/ui/use-toast"
import useOrganization from "@/hooks/use-organization"
import { updateOrgName } from "@/lib/service"
import { zodResolver } from "@hookform/resolvers/zod"
import { setCookie } from "cookies-next"
import { useState } from "react"
import { useForm } from "react-hook-form"
import { z } from "zod"

export const changeOrgNameFormSchema = z.object({
  orgName: z
    .string()
    .min(1, "Organization name is required")
    .max(39, "Organization name can only have a max of 39 characters."),
})

interface UseUpdateOrgNameFormOptions {
  organizationId: number
  updateOrgNameInCookie: (name: string) => void
}

export default function useUpdateOrgNameForm({
  organizationId,
  updateOrgNameInCookie,
}: UseUpdateOrgNameFormOptions) {
  const [updateOrgNameLoading, setUpdateOrgNameLoading] = useState(false)
  const form = useForm<z.infer<typeof changeOrgNameFormSchema>>({
    resolver: zodResolver(changeOrgNameFormSchema),
    defaultValues: {
      orgName: "",
    },
  })

  const handleUpdateOrgName = (name: string) => {
    setUpdateOrgNameLoading(true)
    updateOrgName(organizationId, name)
      .then(() => {
        toast({
          variant: "success",
          title: "Organization Name Update Successful",
          description: (
            <p>
              Your organization's name has been updated to{" "}
              <strong>{name}</strong>.
            </p>
          ),
        })
        form.reset({
          orgName: "",
        })
        updateOrgNameInCookie(name)
      })
      .catch((err) => {
        toast({
          variant: "destructive",
          title: "Update Organization Name Error",
          description: "Could not update organization. Please try again later.",
        })
        console.error(err)
      })
      .finally(() => {
        setUpdateOrgNameLoading(false)
      })
  }

  const onFormSubmit = ({
    orgName,
  }: z.infer<typeof changeOrgNameFormSchema>) => {
    handleUpdateOrgName(orgName)
  }

  return { updateOrgNameLoading, form, onFormSubmit }
}
