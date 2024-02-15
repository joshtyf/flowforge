import { Button } from "@/components/ui/button"
import { Card, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Pipeline } from "@/types/pipeline"
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination"
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime"
interface ServicesViewProps {
  services: Pipeline[] | void
  router: AppRouterInstance
}

export default function ServicesView({ services, router }: ServicesViewProps) {
  return services && services.length === 0 ? (
    <div className="flex justify-center items-center w-full h-3/5">
      <h1 className="font-bold text-3xl">No services available</h1>
    </div>
  ) : (
    <>
      <div className=" grid grid-cols-auto-fill-min-20 gap-y-10 max-h-[75%] overflow-y-auto">
        {services?.map((service) => (
          <div key={service.id} className="flex items-center justify-center">
            <Card className="w-[250px] shadow">
              <CardHeader>
                <CardTitle>{service.pipeline_name}</CardTitle>
                {/* TODO: Add description once available */}
                {/* <CardDescription>{service.description}</CardDescription> */}
              </CardHeader>
              <CardFooter className="flex justify-end">
                <Button
                  variant="outline"
                  onClick={() => router.push(`/service-catalog/${service.id}`)}
                >
                  Request
                </Button>
              </CardFooter>
            </Card>
          </div>
        ))}
      </div>
      {/* TODO: Implement pagination logic */}
      <div className="w-full flex justify-center absolute bottom-0">
        <Pagination>
          <PaginationContent>
            <PaginationItem>
              <PaginationPrevious />
            </PaginationItem>
            <PaginationItem>
              <PaginationLink isActive>1</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink>2</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink>3</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationNext />
            </PaginationItem>
          </PaginationContent>
        </Pagination>
      </div>
    </>
  )
}
