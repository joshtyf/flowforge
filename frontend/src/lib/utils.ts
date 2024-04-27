import {
  ServiceRequestStep,
  ServiceRequestSteps,
} from "@/types/service-request"
import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function isJson(item: string) {
  let value = typeof item !== "string" ? JSON.stringify(item) : item
  try {
    value = JSON.parse(value)
  } catch (e) {
    return false
  }

  return typeof value === "object" && value !== null
}

export function formatDateString(date: Date) {
  const options: Intl.DateTimeFormatOptions = {
    day: "numeric",
    month: "short",
    year: "numeric",
  }

  return date.toLocaleDateString("en-UK", options)
}

export function formatTimeDifference(date: Date) {
  if (!date || isNaN(date.getTime())) {
    return "Invalid date"
  }

  const now = new Date()

  const diff = now.getTime() - date.getTime()
  if (diff < 0) {
    return "Date is ahead of current time"
  }

  if (diff < 3600000) {
    const minutes = Math.floor(diff / 60000)
    return `${minutes} min${minutes > 1 ? "s" : ""} ago`
  } else if (diff < 86400000) {
    const hours = Math.floor(diff / 3600000) // Convert to hours
    return `${hours} hr${hours > 1 ? "s" : ""} ago`
  } else if (diff < 31536000000) {
    const days = Math.floor(diff / 86400000) // Convert to days
    return `${days} day${days > 1 ? "s" : ""} ago`
  } else {
    const years = Math.floor(diff / 31536000000) // Convert to years
    return `${years} year${years > 1 ? "s" : ""} ago`
  }
}

export function isArrayValuesUnique<T>(arr: T[]): boolean {
  return new Set(arr).size === arr.length
}

export function isArrayValuesString(arr: unknown[]): arr is string[] {
  return arr.every((item): item is string => typeof item === "string")
}

export function createStepsFromObject(
  firstStep: string,
  steps?: ServiceRequestSteps
): ServiceRequestStep[] {
  if (!steps) {
    return []
  }
  const stepsArray: ServiceRequestStep[] = []
  const stepsSet = new Set()
  let currentStep = firstStep
  while (currentStep !== "") {
    if (stepsSet.has(currentStep)) {
      console.error("Cycle detected in service request steps")
      return []
    }
    stepsArray.push(steps[currentStep])
    stepsSet.add(currentStep)
    currentStep = steps[currentStep].next_step_name
  }
  return stepsArray
}
