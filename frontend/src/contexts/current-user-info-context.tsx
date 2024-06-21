import {
  AlertDialog,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"
import { ButtonWithSpinner } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { toast } from "@/components/ui/use-toast"
import { createUser, login } from "@/lib/service"
import { UserInfo } from "@/types/user-profile"
import { zodResolver } from "@hookform/resolvers/zod"
import { createContext, useContext, useEffect, useState } from "react"
import { useForm } from "react-hook-form"
import { z } from "zod"

const CurrentUserInfoContext = createContext<UserInfo | undefined>(undefined)

const createUserFormSchema = z.object({
  username: z
    .string()
    .min(1, "Username is required.")
    .max(39, { message: "Username can only have a max of 39 characters." }),
})

export function CurrentUserInfoContextProvider({
  children,
}: {
  children: React.ReactNode
}) {
  const [userInfo, setUserInfo] = useState<UserInfo>()
  const [openAlertDialog, setOpenAlertDialog] = useState(false)
  const [loadingCreateUser, setLoadingCreateUser] = useState(false)

  useEffect(() => {
    login()
      .then((res) => setUserInfo(res))
      .catch((err) => {
        if (err.response?.status === 404) {
          // New user detected, prompt to create user in Flowforge
          setOpenAlertDialog(true)
        } else {
          // Unexpected error
          console.error(err)
        }
      })
  }, [])

  const form = useForm<z.infer<typeof createUserFormSchema>>({
    resolver: zodResolver(createUserFormSchema),
    defaultValues: {
      username: "",
    },
  })

  const handleCreateUser = (username: string) => {
    setLoadingCreateUser(true)
    createUser(username)
      .then((res) => {
        toast({
          variant: "success",
          title: "Flowforge Account Registration Successful",
          description: (
            <p>
              You are able to access Flowforge features now. Welcome{" "}
              <strong>name</strong>!
            </p>
          ),
        })
        setUserInfo(res)
        setOpenAlertDialog(false)
        form.reset({
          username: "",
        })
      })
      .catch((err) => {
        toast({
          variant: "destructive",
          title: "Flowforge Account Registration Error",
          description:
            "Could not create new user for your account. Please try again later.",
        })
        console.error(err)
      })
      .finally(() => {
        setLoadingCreateUser(false)
      })
  }

  const onFormSubmit = ({ username }: z.infer<typeof createUserFormSchema>) => {
    handleCreateUser(username)
  }

  return (
    <CurrentUserInfoContext.Provider value={userInfo}>
      {children}
      <AlertDialog open={openAlertDialog} onOpenChange={setOpenAlertDialog}>
        <AlertDialogContent>
          <Form {...form}>
            <form
              onSubmit={form.handleSubmit(onFormSubmit)}
              className="space-y-5"
            >
              <AlertDialogHeader>
                <AlertDialogTitle>Welcome to Flowforge!</AlertDialogTitle>
                <AlertDialogDescription>
                  Since its your first time using Flowforge, we will need a
                  username for your account.
                </AlertDialogDescription>
              </AlertDialogHeader>
              <FormField
                control={form.control}
                name="username"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Your Username</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <AlertDialogFooter>
                <ButtonWithSpinner
                  isLoading={loadingCreateUser}
                  disabled={loadingCreateUser}
                  type="submit"
                >
                  Create
                </ButtonWithSpinner>
              </AlertDialogFooter>
            </form>
          </Form>
        </AlertDialogContent>
      </AlertDialog>
    </CurrentUserInfoContext.Provider>
  )
}

export function useCurrentUserInfo() {
  const context = useContext(CurrentUserInfoContext)

  return context
}
