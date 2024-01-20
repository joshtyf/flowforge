import React from "react"
import type { Metadata } from "next"
import { Poppins } from "next/font/google"
import "@/styles/globals.css"

const font = Poppins({ weight: "400", subsets: ["latin"] })

export const metadata: Metadata = {
  title: "FlowForge",
  description: "Service Request Management App",
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" style={{ height: "100%" }}>
      <body className={font.className}>{children}</body>
    </html>
  )
}
