type PipelineStep = {
  step_name: string
  step_type: "API" | "WAIT_FOR_APPROVAL"
  next_step_name: string
  prev_step_name: string
  parameters: {
    method: "GET" | "POST" | "PATCH"
    url: "string"
  }
  is_terminal_step: boolean
}

type Pipeline = {
  pipeline_name: string
  version: number
  first_step_name: string
  steps: PipelineStep[]
  created_on?: string
}

export type { Pipeline, PipelineStep }
