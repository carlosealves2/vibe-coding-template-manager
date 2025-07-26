import {
  Template,
  CreateTemplateRequest,
  Project,
  CreateProjectRequest,
} from "@/types";

const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1";

class ApiClient {
  private async request<T>(
    endpoint: string,
    options?: RequestInit,
  ): Promise<T> {
    const url = `${API_BASE_URL}${endpoint}`;

    const response = await fetch(url, {
      headers: {
        "Content-Type": "application/json",
        ...options?.headers,
      },
      ...options,
    });

    if (!response.ok) {
      const error = await response
        .json()
        .catch(() => ({ error: "Unknown error" }));
      throw new Error(error.error || `HTTP ${response.status}`);
    }

    return response.json();
  }

  // Templates
  async getTemplates(): Promise<Template[]> {
    return this.request<Template[]>("/templates");
  }

  async getTemplate(id: number): Promise<Template> {
    return this.request<Template>(`/templates/${id}`);
  }

  async createTemplate(data: CreateTemplateRequest): Promise<Template> {
    return this.request<Template>("/templates", {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  async updateTemplate(
    id: number,
    data: Partial<CreateTemplateRequest>,
  ): Promise<Template> {
    return this.request<Template>(`/templates/${id}`, {
      method: "PUT",
      body: JSON.stringify(data),
    });
  }

  async deleteTemplate(id: number): Promise<void> {
    return this.request<void>(`/templates/${id}`, {
      method: "DELETE",
    });
  }

  // Projects
  async getProjects(): Promise<Project[]> {
    return this.request<Project[]>("/projects");
  }

  async getProject(id: number): Promise<Project> {
    return this.request<Project>(`/projects/${id}`);
  }

  async createProject(data: CreateProjectRequest): Promise<Project> {
    return this.request<Project>("/projects", {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  async deleteProject(id: number): Promise<void> {
    return this.request<void>(`/projects/${id}`, {
      method: "DELETE",
    });
  }

  streamProjectLogs(id: number, onMessage: (msg: string) => void): EventSource {
    const url = `${API_BASE_URL}/projects/${id}/logs`;
    const ev = new EventSource(url);
    ev.onmessage = (e) => {
      onMessage(e.data);
    };
    return ev;
  }
}

export const apiClient = new ApiClient();
