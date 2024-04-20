import React from "react"
import { Column } from "@tanstack/react-table"
import { ChevronsUpDown, EyeOff, SortAsc, SortDesc } from "lucide-react"

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

interface DataTableColumnHeaderProps<TData, TValue>
  extends React.HTMLAttributes<HTMLDivElement> {
  column: Column<TData, TValue>
  title: string
  filterableValues?: string[]
}

export function DataTableColumnHeaderFilterableValue<TData, TValue>({
  column,
  title,
  className,
  filterableValues,
}: DataTableColumnHeaderProps<TData, TValue>) {
  const [position, setPosition] = React.useState("all")
  if (!column.getCanSort()) {
    return <div className={cn(className)}>{title}</div>
  }

  function renderLockIcon(): React.JSX.Element {
    const sorted = column.getIsSorted()
    if (sorted === "desc") {
      return <SortDesc className="ml-2 h-4 w-4" />
    } else if (sorted === "asc") {
      return <SortAsc className=" ml-2 h-4 w-4" />
    }
    return <ChevronsUpDown className="ml-2 h-4 w-4" />
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
            {renderLockIcon()}
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
            {filterableValues?.map((value) => (
              <DropdownMenuRadioItem
                key={value}
                onClick={() => column.setFilterValue(value)}
                value={value.toLowerCase()}
              >
                {value}
              </DropdownMenuRadioItem>
            ))}
          </DropdownMenuRadioGroup>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}
