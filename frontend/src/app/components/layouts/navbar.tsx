import React from "react"
import { Button } from "../ui/button"
import { ChevronDown, LucideUser, Menu } from "lucide-react"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@radix-ui/react-dropdown-menu"

const UserActionsDropdown = () => (
  <DropdownMenu>
    <DropdownMenuTrigger asChild>
      <Button
        data-testid="user-profile-button"
        className="hover:text-primary hover:bg-transparent"
        variant="ghost"
      >
        <LucideUser className="mr-2" />
        User
        <ChevronDown className="ml-2" />
      </Button>
    </DropdownMenuTrigger>
    <DropdownMenuContent>
      <DropdownMenuGroup>
        <DropdownMenuItem>
          <Button
            data-testid="logout-button"
            className="hover:text-primary hover:bg-transparent"
            variant="ghost"
          >
            Logout
          </Button>
        </DropdownMenuItem>
      </DropdownMenuGroup>
    </DropdownMenuContent>
  </DropdownMenu>
)

export default function Navbar() {
  return (
    <div className="flex-col border-b md:flex">
      <div className="flex h-16 items-center px-4">
        <Button variant="outline" size="icon">
          <Menu />
        </Button>
        <div className="ml-auto w-1/10">
          <UserActionsDropdown />
        </div>
      </div>
    </div>
  )
}
