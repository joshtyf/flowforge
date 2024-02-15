import { Skeleton } from "@/components/ui/skeleton"

const NO_OF_SKELETONS = 10

const createSkeletons = (number: number) => {
  const skeleton = []
  for (let i = 0; i < number; i++) {
    skeleton.push(
      <div key={i} className="flex items-center justify-center">
        <Skeleton className="w-[250px] h-[100px]" />
      </div>
    )
  }
  return skeleton
}

export default function ServicesSkeletonView() {
  return (
    <>
      <div className=" grid grid-cols-auto-fill-min-20 gap-y-10 max-h-[75%] overflow-y-auto">
        {createSkeletons(NO_OF_SKELETONS)}
      </div>
      <div className="w-full flex justify-center absolute bottom-0">
        <Skeleton className="w-[300px] h-8" />
      </div>
    </>
  )
}
