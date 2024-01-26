import { FieldErrorProps } from "@rjsf/utils"
import { ReactElement } from "react"

export default function FieldErrorTemplate(props: FieldErrorProps) {
  const { errors } = props
  return (
    <ul className="list-disc ml-5">
      {errors?.map((error: string | ReactElement, i: number) => {
        return (
          <>
            <li key={i} className="text-destructive">
              {error.toString().charAt(0).toUpperCase() +
                error.toString().slice(1)}
            </li>
            <li key={i} className="text-destructive">
              {error.toString().charAt(0).toUpperCase() +
                error.toString().slice(1)}
            </li>
          </>
        )
      })}
    </ul>
  )
}
