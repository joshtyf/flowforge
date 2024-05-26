import MainNavigationLayout from "@/components/layouts/main-navigation-layout"

export default function NonAdminLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return <MainNavigationLayout>{children}</MainNavigationLayout>
}
