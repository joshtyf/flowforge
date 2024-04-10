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
import usePagination from "../_hooks/use-pagination"
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
interface ServicesViewProps {
  services: Pipeline[] | void
  router: AppRouterInstance
}

const createPaginationItems = (
  currentPage: number,
  noOfPages: number,
  handleClickPageNo: (pageNo: number) => void
) => {
  const paginationItems = []
  for (let i = 1; i <= noOfPages; i++) {
    paginationItems.push(
      <PaginationItem key={i}>
        <Button
          variant="ghost"
          className="p-0"
          onClick={() => {
            handleClickPageNo(i)
          }}
        >
          <PaginationLink isActive={i === currentPage}>{i}</PaginationLink>
        </Button>
      </PaginationItem>
    )
  }
  return paginationItems
}

export default function ServicesView({ services, router }: ServicesViewProps) {
  const {
    page,
    noOfPages,
    handleClickNextPage,
    isNextDisabled,
    handleClickPrevPage,
    isPrevDisabled,
    handleClickPageNo,
    handleSetItemsPerPage,
    servicesAtPage,
  } = usePagination({ itemsPerPage: 4, services })

  return services && services.length === 0 ? (
    <div className="flex justify-center items-center w-full h-3/5">
      <h1 className="font-bold text-3xl">No services available</h1>
    </div>
  ) : (
    <>
      <div className=" grid grid-cols-auto-fill-min-20 gap-y-10 max-h-[75%] overflow-y-auto">
        {servicesAtPage?.map((service) => (
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
      <div className="w-full flex justify-center absolute bottom-0">
        <Pagination>
          <PaginationContent>
            <PaginationItem>
              <Button
                variant="ghost"
                className="p-0"
                disabled={isPrevDisabled}
                onClick={handleClickPrevPage}
              >
                <PaginationPrevious />
              </Button>
            </PaginationItem>
            {createPaginationItems(page, noOfPages, handleClickPageNo)}
            <PaginationItem>
              <Button
                variant="ghost"
                className="p-0"
                disabled={isNextDisabled}
                onClick={handleClickNextPage}
              >
                <PaginationNext />
              </Button>
            </PaginationItem>
            <Select defaultValue="10" onValueChange={handleSetItemsPerPage}>
              <SelectTrigger className="w-[100px]">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectLabel>Items per page</SelectLabel>
                  <SelectItem value={"5"}>5</SelectItem>
                  <SelectItem value={"10"}>10</SelectItem>
                  <SelectItem value={"15"}>15</SelectItem>
                  <SelectItem value={"20"}>20</SelectItem>
                  <SelectItem value={"25"}>25</SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
          </PaginationContent>
        </Pagination>
      </div>
    </>
  )
}
