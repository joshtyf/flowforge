import React from "react"
import { Button } from "@/components/ui/button"
import { ChevronDown, LucideUser, Menu } from "lucide-react"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import Link from "next/link"
import { getAuth0LogoutLink } from "@/lib/auth0"

interface UserActionsDropdownProps {
  username: string
}

const UserActionsDropdown = ({ username }: UserActionsDropdownProps) => (
  <DropdownMenu>
    <DropdownMenuTrigger asChild>
      <Button
        data-testid="user-profile-button"
        className="hover:text-primary hover:bg-transparent"
        variant="ghost"
      >
        <LucideUser className="mr-2" />
        <p>{username}</p>
        <ChevronDown className="ml-2" />
      </Button>
    </DropdownMenuTrigger>
    <DropdownMenuContent>
      <DropdownMenuGroup>
        <DropdownMenuItem>
          <Link href={getAuth0LogoutLink()}>
            <Button
              data-testid="logout-button"
              className="hover:text-primary hover:bg-transparent"
              variant="ghost"
            >
              <p>Logout</p>
            </Button>
          </Link>
        </DropdownMenuItem>
      </DropdownMenuGroup>
    </DropdownMenuContent>
  </DropdownMenu>
)

interface NavbarProps {
  toggleSidebar: () => void
  username: string
}

export default function Navbar({ toggleSidebar, username }: NavbarProps) {
  return (
    <div className="flex-col  md:flex">
      <div className="flex h-16 border-b items-center px-4">
        <Button variant="ghost" size="icon" onClick={toggleSidebar}>
          <Menu />
        </Button>
        <div className="ml-auto w-1/10">
          <UserActionsDropdown username={username} />
        </div>
      </div>
    </div>
  )
}
