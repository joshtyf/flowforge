import { ArrayFieldTemplateProps } from "@rjsf/utils"

export default function ArrayFieldTemplate(props: ArrayFieldTemplateProps) {
  return (
    <div>
      {props.items.map((element) => {
        return element.children
      })}
      {props.canAdd && (
        <button type="button" onClick={props.onAddClick}></button>
      )}
    </div>
  )
}
