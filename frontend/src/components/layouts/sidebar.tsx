import { cn } from "@/lib/utils"
import Link from "next/link"
import React from "react"
import { buttonVariants } from "../ui/button"
import { GitBranchPlus, LibraryBig, Workflow } from "lucide-react"

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
}

const links: LinkType[] = [
  {
    title: "Service Catalog",
    icon: LibraryBig,
    href: "/service-catalog",
    variant: "ghost",
  },
  {
    title: "Service Request Dashboard",
    icon: Workflow,
    href: "#",
    variant: "ghost",
  },
  {
    title: "Service Creation Dashboard",
    icon: GitBranchPlus,
    href: "#",
    variant: "ghost",
  },
]

interface SidebarProps {
  className?: string
}

export default function Sidebar({ className }: SidebarProps) {
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
        {links.map((link, index) => (
          <Link
            key={index}
            href={link.href}
            className={cn(
              buttonVariants({ variant: link.variant, size: "sm" }),
              link.variant === "default" &&
                "dark:bg-muted dark:text-white dark:hover:bg-muted dark:hover:text-white",
              "justify-start"
            )}
          >
            <link.icon className="h-5 w-5 mr-2" />
            <span className="inline">{link.title}</span>
          </Link>
        ))}
      </nav>
    </div>
  )
}
