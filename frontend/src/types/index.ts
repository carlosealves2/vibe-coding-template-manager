export interface Template {
  id: number;
  name: string;
  description: string;
  git_url: string;
  language: string;
  tags: string;
  created_at: string;
  updated_at: string;
}

export interface CreateTemplateRequest {
  name: string;
  description: string;
  git_url: string;
  language: string;
  tags: string;
}

export interface Project {
  id: number;
  name: string;
  git_url: string;
  template_id: number;
  template: Template;
  status: 'creating' | 'ready' | 'error';
  created_at: string;
  updated_at: string;
}

export interface CreateProjectRequest {
  name: string;
  template_id: number;
}
