import { FieldErrorProps } from "@rjsf/utils"
import { ReactElement } from "react"

export default function FieldErrorTemplate(props: FieldErrorProps) {
  const { errors } = props
  return (
    <ul>
      {errors?.map((error: string | ReactElement, i: number) => {
        return (
          <li key={i} className="error">
            {error.toString()}
          </li>
        )
      })}
    </ul>
  )
}
