"use client"

import { setCookie } from "cookies-next"
import { useRouter } from "next/navigation"

export default function OrganisationsPage() {
  // TODO: Replace with actual organisations
  const orgs = [
    {
      id: 1,
      name: "Organisation 1",
    },
    {
      id: 2,
      name: "Organisation 2",
    },
  ]
  const router = useRouter()
  return (
    <div className="mt-20 flex flex-col justify-center items-center">
      <p className="mb-8 text-2xl">Your Organisations</p>
      <div className="border rounded-md w-2/5">
        <ul className="divide-y divide-slate-200">
          {orgs.map((org) => (
            <li
              key={org.id}
              className="px-8 py-4 cursor-pointer text-xl hover:text-blue-500"
              onClick={(e) => {
                setCookie("org_id", org.id)
                router.push("/")
              }}
            >
              {org.name}
            </li>
          ))}
        </ul>
      </div>
    </div>
  )
}
