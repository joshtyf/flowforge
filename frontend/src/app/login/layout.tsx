import { ReactNode } from "react"

interface LoginLayoutProps {
  children: ReactNode
}
export default function LoginLayout({ children }: LoginLayoutProps) {
  return (
    <div className="flex flex-col justify-center items-center min-h-[100vh] space-y-4">
      <span className="my-4 flex space-x-4 items-center">
        <img
          src={"/flowforge.png"}
          width="30"
          height="30"
          alt="flowforge icon"
        />
        <h1 className="text-3xl">Welcome to Flowforge</h1>
      </span>
      {children}
    </div>
  )
}
