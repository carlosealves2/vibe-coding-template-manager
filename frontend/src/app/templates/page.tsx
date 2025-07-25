'use client';

import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { TemplateForm } from '@/components/template-form';
import { TemplateList } from '@/components/template-list';
import { Template } from '@/types';
import { Plus } from 'lucide-react';

export default function TemplatesPage() {
  const [showForm, setShowForm] = useState(false);
  const [editingTemplate, setEditingTemplate] = useState<Template | undefined>();
  const [refreshTrigger, setRefreshTrigger] = useState(0);

  const handleFormSuccess = () => {
    setShowForm(false);
    setEditingTemplate(undefined);
    setRefreshTrigger(prev => prev + 1);
  };

  const handleFormCancel = () => {
    setShowForm(false);
    setEditingTemplate(undefined);
  };

  const handleEdit = (template: Template) => {
    setEditingTemplate(template);
    setShowForm(true);
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Templates</h1>
          <p className="text-gray-600 mt-1">
            Manage your repository templates
          </p>
        </div>
        <Button onClick={() => setShowForm(true)} className="flex items-center gap-2">
          <Plus className="h-4 w-4" />
          Add Template
        </Button>
      </div>

      {showForm && (
        <div className="mb-8">
          <TemplateForm
            template={editingTemplate}
            onSuccess={handleFormSuccess}
            onCancel={handleFormCancel}
          />
        </div>
      )}

      <TemplateList onEdit={handleEdit} refreshTrigger={refreshTrigger} />
    </div>
  );
}
