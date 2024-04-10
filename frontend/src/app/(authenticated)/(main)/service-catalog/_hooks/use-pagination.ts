import { Pipeline } from "@/types/pipeline"
import { useCallback, useEffect, useState } from "react"

interface UsePaginationOptions {
  itemsPerPage?: number
  services?: Pipeline[] | void
}

type ItemsPerRowType = "5" | "10" | "15" | "20" | "25"

const ITEMS_PER_PAGE = {
  "5": 5,
  "10": 10,
  "15": 15,
  "20": 20,
  "25": 25,
}

const usePagination = ({ services }: UsePaginationOptions) => {
  const [page, setPage] = useState<number>(1)
  const [itemsPerPage, setItemsPerPage] = useState<number>(ITEMS_PER_PAGE["10"])
  const [servicesAtPage, setServicesAtPage] = useState<Pipeline[]>(
    services ? services.slice(0, itemsPerPage) : []
  )
  const noOfPages = services ? Math.ceil(services.length / itemsPerPage) : 0

  const handleSetServicesAtPage = useCallback(
    (pageNo: number, items: number = itemsPerPage) => {
      const startIndex = (pageNo - 1) * items
      const endIndex = pageNo * items
      setServicesAtPage(services ? services.slice(startIndex, endIndex) : [])
    },
    [services, itemsPerPage]
  )

  useEffect(() => {
    if (page > noOfPages) {
      // Set page to last page
      const newPageNo = noOfPages
      setPage(newPageNo)
      handleSetServicesAtPage(newPageNo)
    } else {
      handleSetServicesAtPage(page)
    }
  }, [itemsPerPage, page, services, handleSetServicesAtPage, noOfPages])

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

  const handleSetItemsPerPage = (itemsPerPageValue: ItemsPerRowType) => {
    const itemsPerPage = ITEMS_PER_PAGE[itemsPerPageValue]
    setItemsPerPage(itemsPerPage)
  }
  return {
    page,
    noOfPages,
    handleClickNextPage,
    isNextDisabled: page === noOfPages,
    handleClickPrevPage,
    isPrevDisabled: page === 1,
    handleClickPageNo,
    handleSetItemsPerPage,
    servicesAtPage,
  }
}

export default usePagination
