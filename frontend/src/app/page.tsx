import Link from 'next/link';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { FileText, FolderPlus, GitBranch } from 'lucide-react';

export default function HomePage() {
  return (
    <div className="space-y-8">
      <div className="text-center space-y-4">
        <h1 className="text-5xl font-extrabold bg-gradient-to-r from-purple-600 to-blue-600 bg-clip-text text-transparent">
          Template Manager
        </h1>
        <p className="text-lg text-gray-600 dark:text-gray-300 max-w-2xl mx-auto">
          Manage your project templates and create new projects with ease.
          Store your favorite repository templates and generate new projects instantly.
        </p>
      </div>

      <div className="grid md:grid-cols-2 gap-6 max-w-4xl mx-auto">
        <Card className="hover:shadow-lg transition-shadow bg-white/70 dark:bg-gray-800/50 backdrop-blur border-gray-200 dark:border-gray-700">
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <FileText className="h-5 w-5" />
              Templates
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-gray-600 dark:text-gray-300">
              Manage your repository templates. Add new templates from GitHub repositories 
              and organize them by language and tags.
            </p>
            <Link href="/templates">
              <Button className="w-full">
                Manage Templates
              </Button>
            </Link>
          </CardContent>
        </Card>

        <Card className="hover:shadow-lg transition-shadow bg-white/70 dark:bg-gray-800/50 backdrop-blur border-gray-200 dark:border-gray-700">
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <FolderPlus className="h-5 w-5" />
              Projects
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <p className="text-gray-600 dark:text-gray-300">
              Create new projects from your templates. Choose a template, give your 
              project a name, and we&apos;ll create a new repository for you.
            </p>
            <Link href="/projects">
              <Button className="w-full">
                View Projects
              </Button>
            </Link>
          </CardContent>
        </Card>
      </div>

      <div className="bg-white/70 dark:bg-gray-800/50 backdrop-blur rounded-lg p-6 max-w-4xl mx-auto border border-gray-200 dark:border-gray-700">
        <h2 className="text-2xl font-bold mb-4 flex items-center gap-2">
          <GitBranch className="h-6 w-6" />
          How it works
        </h2>
        <div className="grid md:grid-cols-3 gap-6">
          <div className="text-center">
            <div className="bg-blue-100 dark:bg-blue-900 rounded-full w-12 h-12 flex items-center justify-center mx-auto mb-3">
              <span className="text-blue-600 dark:text-blue-400 font-bold">1</span>
            </div>
            <h3 className="font-semibold mb-2">Add Templates</h3>
            <p className="text-sm text-gray-600 dark:text-gray-300">
              Register your favorite GitHub repositories as templates
            </p>
          </div>
          <div className="text-center">
            <div className="bg-green-100 dark:bg-green-900 rounded-full w-12 h-12 flex items-center justify-center mx-auto mb-3">
              <span className="text-green-600 dark:text-green-400 font-bold">2</span>
            </div>
            <h3 className="font-semibold mb-2">Create Projects</h3>
            <p className="text-sm text-gray-600 dark:text-gray-300">
              Choose a template and create a new project with a custom name
            </p>
          </div>
          <div className="text-center">
            <div className="bg-purple-100 dark:bg-purple-900 rounded-full w-12 h-12 flex items-center justify-center mx-auto mb-3">
              <span className="text-purple-600 dark:text-purple-400 font-bold">3</span>
            </div>
            <h3 className="font-semibold mb-2">Start Coding</h3>
            <p className="text-sm text-gray-600 dark:text-gray-300">
              Your new repository is ready with a clean history
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
