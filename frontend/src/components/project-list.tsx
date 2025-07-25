'use client';

import { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Project } from '@/types';
import { apiClient } from '@/lib/api';
import { Trash2, ExternalLink, Clock, CheckCircle, XCircle } from 'lucide-react';

interface ProjectListProps {
  refreshTrigger?: number;
}

export function ProjectList({ refreshTrigger }: ProjectListProps) {
  const [projects, setProjects] = useState<Project[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchProjects = async () => {
    try {
      setLoading(true);
      const data = await apiClient.getProjects();
      setProjects(data || []);
      setError(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch projects');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchProjects();
  }, [refreshTrigger]);

  const handleDelete = async (id: number) => {
    if (!confirm('Are you sure you want to delete this project?')) return;

    try {
      await apiClient.deleteProject(id);
      setProjects(prev => prev.filter(p => p.id !== id));
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to delete project');
    }
  };

  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'creating':
        return <Clock className="h-4 w-4 text-yellow-600 dark:text-yellow-400" />;
      case 'ready':
        return <CheckCircle className="h-4 w-4 text-green-600 dark:text-green-400" />;
      case 'error':
        return <XCircle className="h-4 w-4 text-red-600 dark:text-red-400" />;
      default:
        return null;
    }
  };

  const getStatusVariant = (status: string) => {
    switch (status) {
      case 'creating':
        return 'secondary';
      case 'ready':
        return 'default';
      case 'error':
        return 'destructive';
      default:
        return 'outline';
    }
  };

  if (loading) {
    return <div className="text-center py-8 text-gray-600 dark:text-gray-300">Loading projects...</div>;
  }

  if (error) {
    return (
      <div className="text-center py-8">
        <p className="text-red-600 dark:text-red-400 mb-4">{error}</p>
        <Button onClick={fetchProjects}>Retry</Button>
      </div>
    );
  }

  if (projects.length === 0) {
    return (
      <div className="text-center py-8">
        <p className="text-gray-600 dark:text-gray-300 mb-4">No projects found</p>
        <p className="text-sm text-gray-500 dark:text-gray-400">Create your first project from a template</p>
      </div>
    );
  }

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
      {projects.map((project) => (
        <Card key={project.id} className="hover:shadow-lg transition-shadow">
          <CardHeader className="pb-3">
            <div className="flex items-start justify-between">
              <CardTitle className="text-lg">{project.name}</CardTitle>
              <Button
                size="sm"
                variant="ghost"
                onClick={() => handleDelete(project.id)}
                className="text-red-600 hover:text-red-700 dark:text-red-400 dark:hover:text-red-300"
              >
                <Trash2 className="h-4 w-4" />
              </Button>
            </div>
          </CardHeader>
          <CardContent className="space-y-3">
            <div className="flex items-center gap-2">
              {getStatusIcon(project.status)}
              <Badge variant={getStatusVariant(project.status)}>
                {project.status}
              </Badge>
            </div>

            <div className="space-y-2">
              <div className="text-sm">
                <span className="font-medium">Template:</span> {project.template?.name}
              </div>
              
              {project.git_url && (
                <div className="flex items-center gap-2 text-sm">
                  <ExternalLink className="h-4 w-4" />
                  <a
                    href={project.git_url}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-blue-600 dark:text-blue-400 hover:underline truncate"
                  >
                    View Repository
                  </a>
                </div>
              )}
            </div>

            <div className="flex flex-wrap gap-2">
              {project.template?.language && (
                <Badge variant="secondary">{project.template.language}</Badge>
              )}
              {project.template?.tags && project.template.tags.split(',').map((tag, index) => (
                <Badge key={index} variant="outline">
                  {tag.trim()}
                </Badge>
              ))}
            </div>

            <div className="text-xs text-gray-500 dark:text-gray-400">
              Created: {new Date(project.created_at).toLocaleDateString()}
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}
