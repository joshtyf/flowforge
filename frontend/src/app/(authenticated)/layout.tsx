// components/layouts/authenticated-layout.tsx

import Navbar from "@/components/layouts/navbar";
import Sidebar from "@/components/layouts/sidebar";
import { Toaster } from "@/components/ui/toaster";
import React, { ReactNode, useState } from "react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

interface AuthenticatedLayoutProps {
  children: ReactNode;
}

export default function AuthenticatedLayout({
  children,
}: AuthenticatedLayoutProps) {
  const queryClient = new QueryClient();

  const [isSidebarOpen, setIsSidebarOpen] = useState(true);

  const toggleSidebar = () => {
    setIsSidebarOpen((isSidebarOpen) => !isSidebarOpen);
  };

  return (
    <QueryClientProvider client={queryClient}>
      <div className="flex flex-row w-full min-h-[100vh]">
        {/* Sidebar */}
        <Sidebar
          className={`${
            isSidebarOpen ? "min-w-[280px] w-[280px]" : "min-w-0 w-0"
          } overflow-x-hidden transition-width duration-300 ease-in-out`}
        />

        {/* Main Content */}
        <div className="w-full">
          {/* Navbar */}
          <Navbar toggleSidebar={toggleSidebar} />

          {/* Main Content Wrapper */}
          <div className="w-full h-full max-h-[90vh] flex justify-center items-center flex-col relative">
            <div className="w-5/6 h-full relative">{children}</div>
          </div>
        </div>

        {/* Toaster */}
        <Toaster />
      </div>
    </QueryClientProvider>
  );
}
