import {
  Select,
  SelectValue,
  SelectTrigger,
  SelectContent,
  SelectGroup,
  SelectItem,
} from "@/components/ui/select"
import {
  FormContextType,
  GenericObjectType,
  WidgetProps,
  enumOptionsIndexForValue,
  enumOptionsValueForIndex,
} from "@rjsf/utils"

export default function CustomSelect<F extends FormContextType = object>(
  props: WidgetProps
) {
  const {
    autofocus,
    disabled,
    formContext = {} as F,
    id,
    multiple,
    onBlur,
    onChange,
    onFocus,
    options,
    placeholder,
    readonly,
    value,
  } = props
  const { readonlyAsDisabled = true } = formContext as GenericObjectType

  const { enumOptions, enumDisabled, emptyValue } = options

  const handleChange = (nextValue: string | number | (string | number)[]) =>
    onChange(enumOptionsValueForIndex(nextValue, enumOptions, emptyValue))

  const handleBlur = () =>
    onBlur(id, enumOptionsValueForIndex(value, enumOptions, emptyValue))

  const handleFocus = () =>
    onFocus(id, enumOptionsValueForIndex(value, enumOptions, emptyValue))

  const selectedIndexes = enumOptionsIndexForValue(value, enumOptions, multiple)

  return (
    <Select
      name={id}
      disabled={disabled || (readonlyAsDisabled && readonly)}
      value={
        typeof selectedIndexes === "undefined" ? emptyValue : selectedIndexes
      }
      onValueChange={!readonly ? handleChange : undefined}
    >
      <SelectTrigger
        autoFocus={autofocus}
        onFocus={!readonly ? handleFocus : undefined}
        onBlur={!readonly ? handleBlur : undefined}
        className="w-[180px]"
      >
        <SelectValue placeholder={placeholder} />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          {enumOptions?.map(({ value, label }, i: number) => {
            const disabled =
              Array.isArray(enumDisabled) && enumDisabled.indexOf(value) != -1
            return (
              <SelectItem
                key={i}
                id={label}
                value={String(i)}
                disabled={disabled}
              >
                {label}
              </SelectItem>
            )
          })}
        </SelectGroup>
      </SelectContent>
    </Select>
  )
}
