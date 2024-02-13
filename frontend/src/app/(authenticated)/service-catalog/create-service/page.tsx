"use client"

import React from "react"
import HeaderAccessory from "@/components/ui/header-accessory"
import { useRouter } from "next/navigation"
import { Button, ButtonWithSpinner } from "@/components/ui/button"
import { ChevronLeft } from "lucide-react"
import useCreateService from "./_hooks/use-create-service"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
export default function CreateServicePage() {
  const router = useRouter()
  const { form, handleTextAreaTabKeyDown, handleSubmitForm, isSubmitting } =
    useCreateService({ router })

  return (
    <>
      <div className="flex flex-col justify-start py-10">
        <HeaderAccessory />
        <div className="flex items-baseline space-x-2">
          <Button
            size="icon"
            variant="ghost"
            onClick={() => router.push("/service-catalog")}
          >
            <ChevronLeft />
          </Button>
          <p className="font-bold text-3xl pt-5">Create Service</p>
        </div>
      </div>
      <div className="flex flex-col justify-center items-center w-full">
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(handleSubmitForm)}
            className="w-full space-y-8"
          >
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="font-bold text-lg">Name</FormLabel>
                  <FormControl>
                    <Input placeholder="Name of service" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="font-bold text-lg">
                    Description
                  </FormLabel>
                  <FormControl>
                    <Input placeholder="Description of service" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <div className="flex w-full justify-center space-x-5">
              <FormField
                control={form.control}
                name="form"
                render={({ field }) => (
                  <FormItem className="w-1/2">
                    <FormLabel className="font-bold text-lg">Form</FormLabel>
                    <FormControl>
                      <Textarea
                        id="textarea"
                        placeholder="Form schema"
                        className="h-[300px]"
                        onKeyDown={handleTextAreaTabKeyDown}
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="pipeline"
                render={({ field }) => (
                  <FormItem className="w-1/2">
                    <FormLabel className="font-bold text-lg">
                      Pipeline
                    </FormLabel>
                    <FormControl>
                      <Textarea
                        id="textarea"
                        className="h-[300px]"
                        placeholder="Pipeline schema"
                        onKeyDown={handleTextAreaTabKeyDown}
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
            <div className="flex justify-end">
              <ButtonWithSpinner
                type="submit"
                disabled={isSubmitting}
                isLoading={isSubmitting}
                size="lg"
              >
                Submit
              </ButtonWithSpinner>
            </div>
          </form>
        </Form>
      </div>
    </>
  )
}
