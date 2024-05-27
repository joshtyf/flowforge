import { useState, useEffect } from "react"

export const useCountdown = (duration: number, onComplete: VoidFunction) => {
  const [countdown, setCountdown] = useState(duration)
  useEffect(() => {
    const interval = setInterval(() => {
      setCountdown((prevCountdown) => {
        if (prevCountdown - 1 === 0) {
          clearInterval(interval)
          onComplete()
          return 0
        } else {
          return prevCountdown - 1
        }
      })
    }, 1000)
    return () => clearInterval(interval)
  }, [onComplete])
  return { countdown }
}
