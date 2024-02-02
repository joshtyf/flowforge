"use client"

import React from "react"
import HeaderAccessory from "@/components/ui/header-accessory"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"
import { ChevronLeft } from "lucide-react"
import useCreateService from "./_hooks/use-create-service"
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
export default function CreateServicePage() {
  const router = useRouter()

  const { form, handleSubmitForm } = useCreateService()

  return (
    <>
      <div className="flex flex-col justify-start py-10">
        <HeaderAccessory />
        <div className="flex items-baseline space-x-2">
          <Button size="icon" variant="ghost" onClick={() => router.back()}>
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
                  <FormLabel>Name</FormLabel>
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
                  <FormLabel>Description</FormLabel>
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
                    <FormLabel>Form</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder="Form schema"
                        className="h-[300px]"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="steps"
                render={({ field }) => (
                  <FormItem className="w-1/2">
                    <FormLabel>Pipeline Steps</FormLabel>
                    <FormControl>
                      <Textarea
                        className="h-[300px]"
                        placeholder="Pipeline steps schema"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
            {/* <div className="w-3/5 flex justify-end"> */}
            <Button type="submit" size="lg">
              Submit
            </Button>
            {/* </div> */}
          </form>
        </Form>
      </div>
    </>
  )
}
