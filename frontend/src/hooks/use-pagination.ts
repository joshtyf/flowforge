import { useState } from "react"

export function usePagination() {
  const [pagination, setPagination] = useState({
    pageSize: 10,
    pageIndex: 0,
  })
  const { pageSize, pageIndex } = pagination

  return {
    limit: pageSize,
    onPaginationChange: setPagination,
    pagination,
    skip: pageSize * pageIndex,
  }
}
