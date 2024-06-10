import { Button, ButtonWithSpinner } from "@/components/ui/button"
import HeaderAccessory from "@/components/ui/header-accessory"
import { Pipeline } from "@/types/pipeline"
import validator from "@rjsf/validator-ajv8"
import Form, { IChangeEvent } from "@rjsf/core"
import FieldTemplate from "@/components/form/custom-templates/field-template"
import FieldErrorTemplate from "@/components/form/custom-templates/field-error-template"
import BaseInputTemplate from "@/components/form/custom-templates/base-input-template"
import ArrayFieldTemplate from "@/components/form/custom-templates/array-field-template"
import CustomCheckboxes from "@/components/form/custom-widgets/custom-checkboxes"
import CustomSelect from "@/components/form/custom-widgets/custom-select"
import { RegistryWidgetsType } from "@rjsf/utils"
import { RJSFSchema, UiSchema } from "@rjsf/utils"
import { ChevronLeft } from "lucide-react"
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime"

interface ServiceRequestViewProps {
  router?: AppRouterInstance
  pipelineName: string
  pipelineDescription?: string
  rjsfSchema: RJSFSchema
  uiSchema: UiSchema
  handleSubmit?: (event: IChangeEvent) => void
  isSubmittingRequest?: boolean
  isSubmitButtonDisabled?: boolean
  viewOnly?: boolean
  formData?: object
}
const widgets: RegistryWidgetsType = {
  CheckboxesWidget: CustomCheckboxes,
  SelectWidget: CustomSelect,
}

export default function ServiceRequestView({
  router,
  pipelineName,
  pipelineDescription,
  rjsfSchema,
  uiSchema,
  handleSubmit,
  isSubmittingRequest,
  isSubmitButtonDisabled,
  viewOnly = false,
  formData,
}: ServiceRequestViewProps) {
  const isSubmitEnabled = handleSubmit && !viewOnly
  return (
    <>
      <div className="flex flex-col justify-start py-10">
        <HeaderAccessory />
        <div className="flex items-baseline space-x-2 pt-5">
          {router && (
            <Button
              size="icon"
              variant="ghost"
              onClick={() => {
                router.back()
              }}
            >
              <ChevronLeft />
            </Button>
          )}

          <p className="font-bold text-3xl">{pipelineName}</p>
        </div>
        <p className={`text-lg pt-3 ${router && "ml-12"} text-gray-500`}>
          {pipelineDescription}
        </p>
      </div>
      <div className="w-full flex justify-center">
        <div className="w-4/5">
          <Form
            disabled={viewOnly}
            formData={formData}
            schema={rjsfSchema}
            uiSchema={uiSchema}
            validator={validator}
            onSubmit={handleSubmit}
            templates={{
              FieldTemplate,
              FieldErrorTemplate,
              BaseInputTemplate,
              ArrayFieldTemplate,
            }}
            widgets={widgets}
            showErrorList={false}
          >
            <div className="flex justify-end">
              {isSubmitEnabled && (
                <ButtonWithSpinner
                  type="submit"
                  disabled={isSubmittingRequest || isSubmitButtonDisabled}
                  isLoading={isSubmittingRequest}
                  size="lg"
                >
                  Submit
                </ButtonWithSpinner>
              )}
            </div>
          </Form>
        </div>
      </div>
    </>
  )
}
