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
      <label htmlFor={id} className="text-1xl">
        {label}
      </label>
      <span className="text-destructive">{required ? "*" : null}</span>
      {description}
      <div className="pt-5">{children}</div>
      <span className="text-destructive">{errors}</span>
      {help}
    </div>
  )
}
