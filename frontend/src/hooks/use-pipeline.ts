import { getPipeline } from "@/lib/service"
import { Pipeline } from "@/types/pipeline"
import { useEffect, useState } from "react"
import useOrganizationId from "./use-organization-id"

interface UsePipelineOptions {
  pipelineId: string
}

const usePipeline = ({ pipelineId }: UsePipelineOptions) => {
  const [pipeline, setPipeline] = useState<Pipeline>()
  const [loading, setLoading] = useState(false)
  const { organizationId } = useOrganizationId()
  useEffect(() => {
    setLoading(true)
    getPipeline(pipelineId, organizationId)
      .then(setPipeline)
      .finally(() => setLoading(false))
  }, [pipelineId, organizationId])

  return {
    pipeline,
    isPipelineLoading: loading,
  }
}

export default usePipeline
