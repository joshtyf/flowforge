import { useState, useEffect } from "react"

export const useCountdown = (duration: number) => {
  const [countdown, setCountdown] = useState(duration)
  useEffect(() => {
    const interval = setInterval(() => {
      if (countdown > 0) {
        setCountdown(countdown - 1)
      }
    }, 1000)
    return () => clearInterval(interval)
  }, [countdown])
  return { countdown }
}
