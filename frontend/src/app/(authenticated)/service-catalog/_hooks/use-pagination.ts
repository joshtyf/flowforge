import { Pipeline } from "@/types/pipeline"
import { useState } from "react"

interface UsePaginationOptions {
  itemsPerPage?: number
  services?: Pipeline[] | void
}

const usePagination = ({
  itemsPerPage = 12,
  services,
}: UsePaginationOptions) => {
  const [page, setPage] = useState<number>(1)
  const [servicesAtPage, setServicesAtPage] = useState<Pipeline[]>(
    services ? services.slice(0, itemsPerPage) : []
  )
  const noOfPages = services ? Math.ceil(services.length / itemsPerPage) : 0

  const handleSetServicesAtPage = (pageNo: number) => {
    const startIndex = (pageNo - 1) * itemsPerPage
    const endIndex = pageNo * itemsPerPage
    setServicesAtPage(services ? services.slice(startIndex, endIndex) : [])
  }

  const handleClickNextPage = () => {
    if (page < noOfPages) {
      handleSetServicesAtPage(page + 1)
      setPage(page + 1)
    }
  }

  const handleClickPrevPage = () => {
    if (page > 1) {
      handleSetServicesAtPage(page - 1)
      setPage(page - 1)
    }
  }

  const handleClickPageNo = (pageNo: number) => {
    handleSetServicesAtPage(pageNo)
    setPage(pageNo)
  }

  return {
    page,
    noOfPages,
    handleClickNextPage,
    isNextDisabled: page === noOfPages,
    handleClickPrevPage,
    isPrevDisabled: page === 1,
    handleClickPageNo,
    servicesAtPage,
  }
}

export default usePagination
