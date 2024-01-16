import React from "react"
import type { Metadata } from "next"
import { Nunito_Sans } from "next/font/google"
import "@/app/styles/globals.css"

const font = Nunito_Sans({ weight: "500", subsets: ["latin"] })

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
