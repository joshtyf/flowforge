import { Button } from "@/components/ui/button"
import { getAuth0AuthorizeLink } from "@/lib/auth0"
import Link from "next/link"

export default function LoginPage() {
  return (
    <div className="space-x-4">
      <Link href={getAuth0AuthorizeLink()}>
        <Button>Login</Button>
      </Link>
    </div>
  )
}
