import { Checkbox } from "@/components/ui/checkbox"
import {
  FormContextType,
  RJSFSchema,
  StrictRJSFSchema,
  WidgetProps,
  enumOptionsDeselectValue,
  enumOptionsIsSelected,
  enumOptionsSelectValue,
  optionId,
} from "@rjsf/utils"

export default function CheckboxesWidget<
  T,
  S extends StrictRJSFSchema = RJSFSchema,
  F extends FormContextType = object,
>({
  id,
  disabled,
  options,
  value,
  autofocus,
  readonly,
  required,
  onChange,
}: WidgetProps<T, S, F>) {
  const { enumOptions, enumDisabled } = options
  const checkboxesValues = Array.isArray(value) ? value : [value]
  const handleOnChange = (index: number) => (checked: boolean) => {
    if (checked) {
      onChange(enumOptionsSelectValue<S>(index, checkboxesValues, enumOptions))
    } else {
      onChange(
        enumOptionsDeselectValue<S>(index, checkboxesValues, enumOptions)
      )
    }
  }

  return (
    <>
      {Array.isArray(enumOptions) &&
        enumOptions.map((option, index: number) => {
          const checked = enumOptionsIsSelected<S>(
            option.value,
            checkboxesValues
          )
          const itemDisabled =
            Array.isArray(enumDisabled) &&
            enumDisabled.indexOf(option.value) !== -1

          return (
            <div
              key={option.value + index}
              className="flex items-center space-x-2 py-1"
            >
              <Checkbox
                required={required}
                checked={checked}
                id={optionId(id, index)}
                name={id}
                className={option.value ? "checked" : "unchecked"}
                autoFocus={autofocus && index === 0}
                onCheckedChange={handleOnChange(index)}
                disabled={disabled || itemDisabled || readonly}
              />
              <label key={option.value + index}>{option.label}</label>
            </div>
          )
        })}
    </>
  )
}
