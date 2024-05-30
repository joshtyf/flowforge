import { Button } from "@/components/ui/button"
import HeaderAccessory from "@/components/ui/header-accessory"
import { Skeleton } from "@/components/ui/skeleton"
import { ChevronLeft } from "lucide-react"
import { useRouter } from "next/navigation"

interface ServiceRequestLogsSkeletonViewProps {}

export default function ServiceRequestLogsSkeletonView({}: ServiceRequestLogsSkeletonViewProps) {
  const router = useRouter()

  return (
    <div className="flex flex-col justify-start py-10">
      <HeaderAccessory />
      <div className="flex items-baseline space-x-2 mt-5">
        <Button
          size="icon"
          variant="ghost"
          onClick={() => {
            router.back()
          }}
        >
          <ChevronLeft />
        </Button>

        <Skeleton className="w-[250px] h-[35px]" />
      </div>
    </div>
  )
}
