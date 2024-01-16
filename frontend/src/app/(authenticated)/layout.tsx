"use client"

import React, { ReactNode } from "react"
import Navbar from "@/app/components/layouts/navbar"

interface AuthenticatedLayoutProps {
  children: ReactNode
}

export default function AuthenticatedLayout({
  children,
}: AuthenticatedLayoutProps) {
  return (
    <>
      <Navbar />
      {children}
    </>
  )
}
