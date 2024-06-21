"use client"

import { cn } from "@/lib/utils"
import Link from "next/link"
import { usePathname } from "next/navigation"

interface SettingsSidebarProps {
  className?: string
}

const linkBaseStyle = "transition-colors w-full flex pl-2 py-1"
const linkInactiveStyle = `${linkBaseStyle} hover:rounded-md hover:bg-muted hover:text-muted-foreground`
const linkActiveStyle = `${linkInactiveStyle} font-bold`

const links = [{ name: "Organization", href: "/settings/organization" }]

export default function SettingsSidebar({ className }: SettingsSidebarProps) {
  const pathname = usePathname()

  return (
    <div className={cn("flex flex-col gap-4 py-2", className)}>
      <p className="font-bold">Settings</p>
      <nav>
        {links.map((link) => (
          <Link
            href={link.href}
            className={cn(
              `${linkBaseStyle}, ${pathname === link.href ? linkActiveStyle : linkInactiveStyle}`
            )}
          >
            {link.name}
          </Link>
        ))}
      </nav>
    </div>
  )
}
