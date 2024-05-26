import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import apiClient from "@/lib/apiClient"
import { getAuth0LogoutLink } from "@/lib/auth0"
import { deleteCookie } from "cookies-next"
import { ChevronDown, LucideUser, Menu } from "lucide-react"
import Link from "next/link"
import { useRouter } from "next/navigation"

interface UserActionsDropdownProps {
  username: string
}

const UserActionsDropdown = ({ username }: UserActionsDropdownProps) => {
  const router = useRouter()
  const logout = () => {
    // Reset auth token upon logout
    deleteCookie("logged_in")
    deleteCookie("access_token")
    deleteCookie("org_id")
    delete apiClient.defaults.headers.Authorization
    router.push(getAuth0LogoutLink())
  }
  return (
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
          <DropdownMenuItem className={"cursor-pointer hover:bg-muted"}>
            <Link href="/organization">
              <Button
                variant="ghost"
                className="hover:text-primary hover:bg-transparent"
              >
                Select Organization
              </Button>
            </Link>
          </DropdownMenuItem>
          <DropdownMenuItem
            onClick={logout}
            className={"cursor-pointer hover:bg-muted"}
          >
            <Button
              data-testid="logout-button"
              className="hover:text-primary hover:bg-transparent"
              variant="ghost"
            >
              <p>Logout</p>
            </Button>
          </DropdownMenuItem>
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}

interface NavbarProps {
  toggleSidebar?: () => void
  username: string
  enableSidebarToggle?: boolean
}

export default function Navbar({
  toggleSidebar,
  username,
  enableSidebarToggle = true,
}: NavbarProps) {
  return (
    <div className="flex-col  md:flex">
      <div className="flex h-16 border-b items-center px-4">
        {enableSidebarToggle && (
          <Button variant="ghost" size="icon" onClick={toggleSidebar}>
            <Menu />
          </Button>
        )}
        <div className="ml-auto w-1/10">
          <UserActionsDropdown username={username} />
        </div>
      </div>
    </div>
  )
}
