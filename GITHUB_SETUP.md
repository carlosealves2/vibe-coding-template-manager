# Como enviar o projeto para o GitHub

## Opção 1: Via GitHub CLI (gh)

Se você tem o GitHub CLI instalado:

```bash
# Navegar para o diretório do projeto
cd template-manager

# Criar repositório no GitHub e fazer push
gh repo create template-manager --public --source=. --remote=origin --push
```

## Opção 2: Via interface web do GitHub

1. **Criar repositório no GitHub:**
   - Acesse https://github.com/new
   - Nome do repositório: `template-manager`
   - Descrição: `Template Manager - Manage Git repository templates and create new projects`
   - Marque como público
   - NÃO inicialize com README, .gitignore ou licença (já temos esses arquivos)
   - Clique em "Create repository"

2. **Conectar repositório local:**
   ```bash
   cd template-manager
   git remote add origin https://github.com/SEU_USERNAME/template-manager.git
   git branch -M main
   git push -u origin main
   ```

## Opção 3: Via SSH (se configurado)

```bash
cd template-manager
git remote add origin git@github.com:SEU_USERNAME/template-manager.git
git branch -M main
git push -u origin main
```

## Estrutura do projeto que será enviada:

```
template-manager/
├── README.md                    # Documentação completa
├── .gitignore                   # Arquivos ignorados
├── start.sh                     # Script de inicialização
├── backend/                     # API Go + Fiber v2
│   ├── cmd/main.go             # Aplicação principal
│   ├── internal/               # Código interno
│   │   ├── domain/             # Entidades e interfaces
│   │   ├── usecase/            # Lógica de negócio
│   │   ├── repository/         # Persistência
│   │   ├── handler/            # Controllers HTTP
│   │   └── config/             # Configuração
│   └── pkg/                    # Pacotes reutilizáveis
│       ├── database/           # Conexão DB
│       └── github/             # Serviço Git
└── frontend/                   # Interface Next.js
    ├── src/
    │   ├── app/                # Páginas (App Router)
    │   ├── components/         # Componentes React
    │   ├── lib/                # Utilitários
    │   └── types/              # Tipos TypeScript
    └── public/                 # Arquivos estáticos
```

## Após o upload:

1. Configure as variáveis de ambiente no backend (.env)
2. Instale as dependências: `./start.sh`
3. Configure seu token do GitHub
4. Execute o projeto conforme instruções no README.md

## URL do repositório:
Após criar, o repositório estará disponível em:
`https://github.com/SEU_USERNAME/template-manager`
