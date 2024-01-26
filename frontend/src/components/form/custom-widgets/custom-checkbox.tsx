import { Checkbox } from "@/components/ui/checkbox"
import { WidgetProps } from "@rjsf/utils"

export default function CustomCheckbox(props: WidgetProps) {
  return (
    <>
      <Checkbox
        id="custom"
        className={props.value ? "checked" : "unchecked"}
        onClick={() => props.onChange(!props.value)}
      />
      <label
        htmlFor="terms"
        className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
      >
        {String(props.value)}
      </label>
    </>
  )
}
