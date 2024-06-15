import { toast } from "@/components/ui/use-toast"
import { createOrg } from "@/lib/service"
import { zodResolver } from "@hookform/resolvers/zod"
import { useState } from "react"
import { useForm } from "react-hook-form"
import { z } from "zod"

interface UseCreateOrganizationOptions {
  refetchOrgs: () => void
}

export const createOrgformSchema = z.object({
  // Following GitHub rules for org name max characters: https://github.com/dead-claudia/github-limits?tab=readme-ov-file#organization-names
  orgName: z
    .string()
    .max(39, {
      message: "Organization name can only have a maximum of 39 characters.",
    })
    .min(1, "Organization name is required"),
})

export default function useCreateOrganizationForm({
  refetchOrgs,
}: UseCreateOrganizationOptions) {
  const [createOrgLoading, setCreateOrgLoading] = useState(false)
  const [openFormDialog, setOpenFormDialog] = useState(false)

  const form = useForm<z.infer<typeof createOrgformSchema>>({
    resolver: zodResolver(createOrgformSchema),
    defaultValues: {
      orgName: "",
    },
  })

  const handleCreateOrg = (name: string) => {
    setCreateOrgLoading(true)
    createOrg(name)
      .then(() => {
        refetchOrgs()
        toast({
          variant: "success",
          title: "Organization Created Successfully",
          description: (
            <p>
              Check the new organization under{" "}
              <strong>Your Organizations</strong>.
            </p>
          ),
        })
        setOpenFormDialog(false)
        form.reset({
          orgName: "",
        })
      })
      .catch((err) => {
        toast({
          variant: "destructive",
          title: "Create Organization Error",
          description: "Could not create organization. Please try again later.",
        })
        console.error(err)
      })
      .finally(() => {
        setCreateOrgLoading(false)
      })
  }

  return {
    form,
    handleCreateOrg,
    createOrgLoading,
    openFormDialog,
    setOpenFormDialog,
  }
}
