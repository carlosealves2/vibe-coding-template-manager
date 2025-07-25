#!/bin/bash

# Script para inicializar a aplicação Template Manager

echo "🚀 Iniciando Template Manager..."

# Verificar se Go está instalado
if ! command -v go &> /dev/null; then
    echo "❌ Go não está instalado. Por favor, instale Go 1.21+ primeiro."
    exit 1
fi

# Verificar se Node.js está instalado
if ! command -v node &> /dev/null; then
    echo "❌ Node.js não está instalado. Por favor, instale Node.js 18+ primeiro."
    exit 1
fi

# Verificar se npm está instalado
if ! command -v npm &> /dev/null; then
    echo "❌ npm não está instalado. Por favor, instale npm primeiro."
    exit 1
fi

echo "✅ Dependências verificadas"

# Configurar backend
echo "📦 Configurando backend..."
cd backend

# Verificar se .env existe
if [ ! -f .env ]; then
    echo "⚠️  Arquivo .env não encontrado. Copiando .env.example..."
    cp .env.example .env
    echo "📝 Por favor, edite o arquivo backend/.env com suas credenciais do GitHub"
    echo "   GITHUB_TOKEN=seu_token_aqui"
    echo "   GITHUB_USERNAME=seu_usuario_aqui"
fi

# Instalar dependências do Go
echo "📥 Instalando dependências do Go..."
go mod tidy

# Voltar para o diretório raiz
cd ..

# Configurar frontend
echo "📦 Configurando frontend..."
cd frontend

# Instalar dependências do Node.js
echo "📥 Instalando dependências do Node.js..."
npm install

# Voltar para o diretório raiz
cd ..

echo "✅ Configuração concluída!"
echo ""
echo "Para iniciar a aplicação:"
echo "1. Configure suas credenciais do GitHub no arquivo backend/.env"
echo "2. Execute o backend: cd backend && go run cmd/main.go"
echo "3. Em outro terminal, execute o frontend: cd frontend && npm run dev"
echo "4. Acesse http://localhost:3000 no seu navegador"
echo ""
echo "📚 Consulte o README.md para mais informações"
