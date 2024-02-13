import { Button } from "@/components/ui/button"
import HeaderAccessory from "@/components/ui/header-accessory"
import { Skeleton } from "@/components/ui/skeleton"
import { ChevronLeft } from "lucide-react"
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime"

interface ServiceRequestSkeletonViewProps {
  router: AppRouterInstance
}

export default function ServiceRequestSkeletonView({
  router,
}: ServiceRequestSkeletonViewProps) {
  return (
    <>
      <div className="flex flex-col justify-start py-10">
        <HeaderAccessory />
        <div className="flex items-baseline space-x-2">
          <Button
            size="icon"
            variant="ghost"
            onClick={() => router.push("/service-catalog")}
          >
            <ChevronLeft />
          </Button>
          <Skeleton className="w-[250px] h-[35px] mt-5" />
        </div>
        <Skeleton className="w-[400px] h-[20px] mt-3 ml-12" />
      </div>
      <div className="w-full flex justify-center pt-5">
        <div className="w-4/5 space-y-8">
          <div>
            <Skeleton className="w-[120px] h-[30px]" />
            <Skeleton className="w-[250px] h-[15px] mt-2" />
            <Skeleton className="w-full h-[30px] mt-5" />
          </div>
          <div>
            <Skeleton className="w-[120px] h-[30px]" />
            <Skeleton className="w-[250px] h-[15px] mt-2" />
            <Skeleton className="w-[180px] h-[30px] mt-5" />
          </div>
          <div>
            <Skeleton className="w-[120px] h-[30px]" />
            <Skeleton className="w-[250px] h-[15px] mt-2" />
            <div className="mt-5 space-y-2">
              <span className="flex items-center space-x-2">
                <Skeleton className="w-[20px] h-[20px]" />
                <Skeleton className="w-[120px] h-[20px]" />
              </span>
              <span className="flex items-center space-x-2">
                <Skeleton className="w-[20px] h-[20px]" />
                <Skeleton className="w-[120px] h-[20px]" />
              </span>
              <span className="flex items-center space-x-2">
                <Skeleton className="w-[20px] h-[20px]" />
                <Skeleton className="w-[120px] h-[20px]" />
              </span>
            </div>
          </div>
          <div className="flex justify-end">
            <Skeleton className="w-[120px] h-[45px]" />
          </div>
        </div>
      </div>
    </>
  )
}
