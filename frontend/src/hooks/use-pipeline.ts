import { getPipeline } from "@/lib/service"
import { Pipeline } from "@/types/pipeline"
import { useEffect, useState } from "react"

interface UsePipelineOptions {
  pipelineId: string
}

const usePipeline = ({ pipelineId }: UsePipelineOptions) => {
  const [pipeline, setPipeline] = useState<Pipeline>()
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    setLoading(true)
    getPipeline(pipelineId)
      .then(setPipeline)
      .finally(() => setLoading(false))
  }, [pipelineId])

  return {
    pipeline,
    isPipelineLoading: loading,
  }
}

export default usePipeline
