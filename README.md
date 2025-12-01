# Chameleon Auth API

Este é o microsserviço de **Identidade e Autenticação** (IAM - Identity and Access Management) do sistema Chameleon Agent AI. Ele atua como o Provedor Central de Identidade (SSO), sendo o único responsável por armazenar senhas e emitir tokens de acesso (JWT).

O serviço é construído com Go, GORM e segue a arquitetura Hexagonal/Clean Architecture, utilizando a biblioteca compartilhada `chameleon-common` para padronização.

---

## 📦 1. Configuração e Instalação

### Pré-requisitos
- Docker e Docker Compose.
- Variáveis de ambiente no `.env` (DB_HOST, DB_PORT=5433, JWT_SECRET).

### Execução Local

1. **Garantir Dependências:** Baixe os módulos (incluindo a lib `chameleon-common`):
```bash
go mod tidy
```