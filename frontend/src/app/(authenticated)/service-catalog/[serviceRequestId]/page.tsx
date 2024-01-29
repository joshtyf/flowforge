"use client"

import React from "react"
import { useParams, useRouter } from "next/navigation"
import HeaderAccessory from "@/components/ui/header-accessory"
import { Button } from "@/components/ui/button"
import { ChevronLeft } from "lucide-react"
import useServiceRequest from "./_hooks/useServiceRequest"
import validator from "@rjsf/validator-ajv8"
import Form from "@rjsf/core"
import FieldTemplate from "@/components/form/custom-templates/field-template"
import FieldErrorTemplate from "@/components/form/custom-templates/field-error-template"
import BaseInputTemplate from "@/components/form/custom-templates/base-input-template"
import ArrayFieldTemplate from "@/components/form/custom-templates/array-field-template"
import {
  convertServiceRequestFormToRJSFSchema,
  generateUiSchema,
} from "@/lib/utils"
import { RegistryWidgetsType } from "@rjsf/utils"
import CustomCheckboxes from "@/components/form/custom-widgets/custom-checkboxes"
import CustomSelect from "@/components/form/custom-widgets/custom-select"

const widgets: RegistryWidgetsType = {
  CheckboxesWidget: CustomCheckboxes,
  SelectWidget: CustomSelect,
}

export default function ServiceRequestPage() {
  const { serviceRequestId } = useParams()
  const serviceRequestIdString = Array.isArray(serviceRequestId)
    ? serviceRequestId[0]
    : serviceRequestId
  const router = useRouter()
  const { serviceRequest, handleSubmit } = useServiceRequest({
    serviceRequestId: serviceRequestIdString,
  })

  const { name, description, form } = serviceRequest

  const uiSchema = generateUiSchema(serviceRequest)
  const rjsfSchema = convertServiceRequestFormToRJSFSchema(form)

  return (
    <>
      <div className="flex flex-col justify-start py-10">
        <HeaderAccessory />
        <div className="flex items-baseline space-x-2">
          <Button size="icon" variant="ghost" onClick={() => router.back()}>
            <ChevronLeft />
          </Button>
          <p className="font-bold text-3xl pt-5">{name}</p>
        </div>
        <p className="text-lg pt-3 ml-12 text-gray-500">{description}</p>
      </div>
      <div className="w-full flex justify-center">
        <div className="w-4/5">
          <Form
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
            <Button className="mt-auto" type="submit">
              Submit
            </Button>
          </Form>
        </div>
      </div>
    </>
  )
}
