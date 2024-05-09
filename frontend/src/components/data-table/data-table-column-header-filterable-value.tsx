import React from "react"
import { Column } from "@tanstack/react-table"
import { EyeOff, Filter, ListFilter } from "lucide-react"

import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuRadioGroup,
  DropdownMenuRadioItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { cn } from "@/lib/utils"
export type FilterableOption<T> = {
  value: T
  name: string
}

interface DataTableColumnHeaderProps<TData, TValue>
  extends React.HTMLAttributes<HTMLDivElement> {
  column: Column<TData, TValue>
  title: string
  filterableOptions?: FilterableOption<TValue>[]
}

export function DataTableColumnHeaderFilterableValue<TData, TValue>({
  column,
  title,
  className,
  filterableOptions,
}: DataTableColumnHeaderProps<TData, TValue>) {
  const [position, setPosition] = React.useState("all")
  if (!column.getCanSort()) {
    return <div className={cn(className)}>{title}</div>
  }

  function renderFilterIcon(): React.JSX.Element {
    const isFiltered = column.getIsFiltered()
    if (isFiltered) {
      return <ListFilter className="ml-2 h-4 w-4" />
    } else {
      return <Filter className=" ml-2 h-4 w-4" />
    }
  }

  return (
    <div className={cn("flex items-center space-x-2", className)}>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
            variant="ghost"
            size="sm"
            className="-ml-3 h-8 data-[state=open]:bg-accent"
          >
            <span>{title}</span>
            {renderFilterIcon()}
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="start">
          <DropdownMenuRadioGroup value={position} onValueChange={setPosition}>
            <DropdownMenuRadioItem
              onClick={() => column.setFilterValue(null)}
              value="all"
            >
              All
            </DropdownMenuRadioItem>
            {filterableOptions?.map((option) => (
              <DropdownMenuRadioItem
                key={option.name}
                onClick={() => column.setFilterValue(option.value)}
                value={option.name}
              >
                {option.name}
              </DropdownMenuRadioItem>
            ))}
          </DropdownMenuRadioGroup>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
