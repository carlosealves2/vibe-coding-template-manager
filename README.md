# Template Manager

Uma aplicação completa para gerenciar templates de repositórios Git e criar novos projetos a partir deles.

## Funcionalidades

- **Gerenciamento de Templates**: Cadastre repositórios Git como templates
- **Criação de Projetos**: Crie novos projetos a partir de templates existentes
- **Integração com GitHub**: Criação automática de repositórios no GitHub
- **Interface Moderna**: Frontend responsivo com Next.js e shadcn/ui
- **API RESTful**: Backend robusto em Go com Fiber v2

## Arquitetura

### Backend (Go + Fiber v2)
- **Clean Architecture**: Separação clara de responsabilidades
- **SOLID Principles**: Código maintível e extensível
- **Repository Pattern**: Abstração da camada de dados
- **Use Cases**: Lógica de negócio isolada
- **SQLite**: Banco de dados leve para desenvolvimento

### Frontend (Next.js + TypeScript)
- **React 18**: Com App Router
- **TypeScript**: Tipagem estática
- **shadcn/ui**: Componentes modernos e acessíveis
- **Tailwind CSS**: Estilização utilitária

## Pré-requisitos

- Go 1.21+
- Node.js 18+
- Git
- Token do GitHub (para criação de repositórios)

## Configuração

### Backend

1. Navegue para o diretório do backend:
```bash
cd backend
```

2. Instale as dependências:
```bash
go mod tidy
```

3. Configure as variáveis de ambiente:
```bash
cp .env.example .env
```

4. Edite o arquivo `.env` com suas credenciais do GitHub:
```env
PORT=8080
GITHUB_TOKEN=seu_token_aqui
GITHUB_USERNAME=seu_usuario_aqui
```

5. Execute o servidor:
```bash
go run cmd/main.go
```

### Frontend

1. Navegue para o diretório do frontend:
```bash
cd frontend
```

2. Instale as dependências:
```bash
npm install
```

3. Execute o servidor de desenvolvimento:
```bash
npm run dev
```

## Uso

1. Acesse `http://localhost:3000` no seu navegador
2. Vá para "Templates" para cadastrar seus repositórios favoritos
3. Vá para "Projects" para criar novos projetos a partir dos templates
4. Ao criar um projeto, um novo repositório será criado no seu GitHub

## API Endpoints

### Templates
- `GET /api/v1/templates` - Lista todos os templates
- `POST /api/v1/templates` - Cria um novo template
- `GET /api/v1/templates/:id` - Busca um template por ID
- `PUT /api/v1/templates/:id` - Atualiza um template
- `DELETE /api/v1/templates/:id` - Remove um template

### Projects
- `GET /api/v1/projects` - Lista todos os projetos
- `POST /api/v1/projects` - Cria um novo projeto
- `GET /api/v1/projects/:id` - Busca um projeto por ID
- `DELETE /api/v1/projects/:id` - Remove um projeto

## Estrutura do Projeto

```
template-manager/
├── backend/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── domain/
│   │   ├── usecase/
│   │   ├── repository/
│   │   ├── handler/
│   │   └── config/
│   └── pkg/
│       ├── database/
│       └── github/
└── frontend/
    ├── src/
    │   ├── app/
    │   ├── components/
    │   ├── lib/
    │   └── types/
    └── public/
```

## Tecnologias Utilizadas

### Backend
- Go 1.21
- Fiber v2 (Framework web)
- GORM (ORM)
- SQLite (Banco de dados)
- go-github (API do GitHub)
- go-git (Manipulação de repositórios Git)

### Frontend
- Next.js 14
- TypeScript
- React 18
- shadcn/ui
- Tailwind CSS
- Lucide React (Ícones)

## Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.
