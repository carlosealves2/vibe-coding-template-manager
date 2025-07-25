'use client';

import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { ProjectForm } from '@/components/project-form';
import { ProjectList } from '@/components/project-list';
import { Plus } from 'lucide-react';

export default function ProjectsPage() {
  const [showForm, setShowForm] = useState(false);
  const [refreshTrigger, setRefreshTrigger] = useState(0);

  const handleFormSuccess = () => {
    setShowForm(false);
    setRefreshTrigger(prev => prev + 1);
  };

  const handleFormCancel = () => {
    setShowForm(false);
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-foreground">Projects</h1>
          <p className="text-muted-foreground mt-1">
            View and manage your created projects
          </p>
        </div>
        <Button onClick={() => setShowForm(true)} className="flex items-center gap-2">
          <Plus className="h-4 w-4" />
          Create Project
        </Button>
      </div>

      {showForm && (
        <div className="mb-8">
          <ProjectForm
            onSuccess={handleFormSuccess}
            onCancel={handleFormCancel}
          />
        </div>
      )}

      <ProjectList refreshTrigger={refreshTrigger} />
    </div>
  );
}
