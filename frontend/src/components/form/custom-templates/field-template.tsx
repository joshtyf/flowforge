import { cn } from "@/lib/utils"
import { FieldTemplateProps } from "@rjsf/utils"

export default function CustomFieldTemplate(props: FieldTemplateProps) {
  const {
    id,
    classNames,
    style,
    label,
    help,
    required,
    description,
    errors,
    children,
  } = props
  return (
    <div className={cn("my-5", classNames)} style={style}>
      <label htmlFor={id} className="text-lg">
        {label}
      </label>
      <span className="text-destructive">{required ? "*" : null}</span>
      <span className="text-sm text-gray-400">{description}</span>
      <div className="pt-5">{children}</div>
      {errors}
      {help}
    </div>
  )
}
