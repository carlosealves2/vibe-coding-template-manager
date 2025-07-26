import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { ThemeProvider } from "@/components/theme-provider";
import { ThemeToggle } from "@/components/theme-toggle";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Template Manager",
  description: "Manage your project templates and create new projects",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={inter.className}>
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <div className="min-h-screen flex flex-col bg-gradient-to-br from-gray-50 to-white dark:from-gray-900 dark:to-gray-950 text-gray-900 dark:text-gray-100">
            <nav className="border-b bg-gradient-to-r from-purple-600 to-blue-600 text-white shadow">
              <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex justify-between items-center h-16">
                  <Link href="/" className="text-xl font-bold">
                    Template Manager
                  </Link>
                  <div className="flex items-center gap-4">
                    <Link href="/templates">
                      <Button variant="link" className="text-white hover:text-gray-200">Templates</Button>
                    </Link>
                    <Link href="/projects">
                      <Button variant="link" className="text-white hover:text-gray-200">Projects</Button>
                    </Link>
                    <ThemeToggle />
                  </div>
                </div>
              </div>
            </nav>
            <main className="flex-1 max-w-7xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
              {children}
            </main>
            <footer className="border-t bg-gradient-to-r from-purple-600 to-blue-600 text-white">
              <div className="max-w-7xl mx-auto px-4 py-4 text-center text-sm">
                Built with Next.js and shadcn/ui
              </div>
            </footer>
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}
