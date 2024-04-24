import { Button } from "@/components/ui/button"
import { Frown } from "lucide-react"
import Link from "next/link"

export default function NotFoundPage() {
  return (
    <div className="flex flex-col justify-center items-center min-h-[100vh] space-y-5">
      <img src={"/flowforge.png"} width="60" height="60" alt="flowforge icon" />
      <h1 className="text-4xl text-destructive">404 Error</h1>
      <span className="flex space-x-2 items-center">
        <p className="text-lg">
          Sorry, we couldn't find the page you were looking for.
        </p>
        <Frown />
      </span>
      <Button asChild>
        <Link href={"."}>Back to Home</Link>
      </Button>
    </div>
  )
}
