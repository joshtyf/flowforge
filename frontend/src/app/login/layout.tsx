import { ReactNode } from "react"

interface LoginLayoutProps {
  children: ReactNode
}
export default function LoginLayout({ children }: LoginLayoutProps) {
  return (
    <div className="flex flex-col justify-center items-center">
      <span className="my-4 flex space-x-2 items-center">
        <img
          src={"/flowforge.png"}
          width="30"
          height="30"
          alt="flowforge icon"
        />
        <p className="text-xl">Welcome to Flowforge</p>
      </span>
      {children}
    </div>
  )
}
