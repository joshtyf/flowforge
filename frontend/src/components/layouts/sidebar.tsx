import { cn } from "@/lib/utils"
import Link from "next/link"
import React from "react"
import { buttonVariants } from "../ui/button"
import { LibraryBig, Workflow, LockKeyhole } from "lucide-react"
import { useUserMemberships } from "@/contexts/user-memberships-context"
import { usePathname } from "next/navigation"

type LinkType = {
  title: string
  icon: React.ElementType
  href: string
  variant:
    | "link"
    | "ghost"
    | "default"
    | "destructive"
    | "outline"
    | "secondary"
    | null
    | undefined
  isAdminFeature?: boolean
  key: string
}

const links: LinkType[] = [
  {
    title: "Service Catalog",
    key: "service-catalog",
    icon: LibraryBig,
    href: "/service-catalog",
    variant: "ghost",
  },
  {
    title: "Your Service Requests Dashboard",
    key: "your-service-request-dashboard",
    icon: Workflow,
    href: "/your-service-request-dashboard",
    variant: "ghost",
  },
  {
    title: "Admin Service Request Dashboard",
    key: "admin-service-request-dashboard",
    icon: LockKeyhole,
    href: "/admin-service-requests-dashboard",
    variant: "ghost",
    isAdminFeature: true,
  },
]

interface SidebarProps {
  className?: string
}

export default function Sidebar({ className }: SidebarProps) {
  const { isAdmin } = useUserMemberships()
  const pathname = usePathname()
  return (
    <div className={cn("group flex flex-col gap-4 border-r", className)}>
      <nav className="grid gap-y-2">
        <div className="flex items-center justify-center h-16 border-b">
          <Link href="/">
            <span className="flex items-center space-x-2">
              <img
                src={"/flowforge.png"}
                width="30"
                height="30"
                alt="flowforge icon"
              />
              <p className="text-2xl font-bold">Flowforge</p>
            </span>
          </Link>
        </div>
        {links.map((link) =>
          link.isAdminFeature && !isAdmin ? (
            <></>
          ) : (
            <Link
              key={link.key}
              href={link.href}
              className={cn(
                buttonVariants({ variant: link.variant, size: "sm" }),
                link.variant === "default" &&
                  "dark:bg-muted dark:text-white dark:hover:bg-muted dark:hover:text-white",
                pathname === link.href && "font-semibold",
                "justify-start"
              )}
            >
              <link.icon className="h-5 w-5 mr-2" />
              <span className="inline">{link.title}</span>
            </Link>
          )
        )}
      </nav>
    </div>
  )
}
