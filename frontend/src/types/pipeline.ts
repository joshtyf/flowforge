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

type PipelineForm = {
  id?: string
  version?: number
  first_step_name?: string
  steps?: PipelineStep[]
  created_on?: string
}

type Pipeline = PipelineForm & {
  pipeline_name: string
  pipeline_description?: string
}

export type { Pipeline, PipelineForm, PipelineStep }
