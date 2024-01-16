"use client"

import Navbar from "@/components/layouts/navbar"
import React, { ReactNode } from "react"

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
