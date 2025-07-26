"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import Link from "next/link";
import { apiClient } from "@/lib/api";
import { Project } from "@/types";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

export default function ProjectDetailPage() {
  const params = useParams();
  const id = Number(params.id);
  const [project, setProject] = useState<Project | null>(null);
  const [logs, setLogs] = useState<string[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchProject = async () => {
      try {
        const data = await apiClient.getProject(id);
        setProject(data);
      } catch {
        setError("Failed to fetch project");
      }
    };

    fetchProject();
    const es = apiClient.streamProjectLogs(id, (msg) => {
      setLogs((prev) => [...prev, msg]);
    });
    return () => es.close();
  }, [id]);

  if (error) return <div className="text-red-600">{error}</div>;
  if (!project) return <div>Loading...</div>;

  return (
    <div className="space-y-4">
      <Link href="/projects" className="text-sm underline">
        Back to projects
      </Link>
      <Card>
        <CardHeader>
          <CardTitle>{project.name}</CardTitle>
        </CardHeader>
        <CardContent className="space-y-2">
          <div className="flex items-center gap-2">
            <Badge>{project.status}</Badge>
            {project.git_url && (
              <a
                href={project.git_url}
                target="_blank"
                rel="noopener noreferrer"
                className="text-blue-600 hover:underline"
              >
                Repository
              </a>
            )}
          </div>
          <div className="text-sm">Template: {project.template?.name}</div>
        </CardContent>
      </Card>
      <Card>
        <CardHeader>
          <CardTitle>Logs</CardTitle>
        </CardHeader>
        <CardContent>
          <pre className="text-sm whitespace-pre-wrap">{logs.join("\n")}</pre>
        </CardContent>
      </Card>
    </div>
  );
}
