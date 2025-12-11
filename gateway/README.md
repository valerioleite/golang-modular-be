# API Gateway - Kong Gateway

Este diretório contém a configuração do API Gateway usando Kong Gateway.

## Funcionalidades Implementadas

### 1. Roteamento Inteligente ✅
- Roteamento baseado em path
- Suporte a versionamento através dos paths
- Mapeamento automático de rotas do gateway para os serviços backend

### 3. Timeouts e Retry ✅
- Timeouts configuráveis por serviço
- Retry automático com backoff exponencial
- Retry configurado para erros 5xx e timeouts
- Número de tentativas configurável por serviço

### 5. Autenticação e Autorização ✅
- Validação JWT usando plugin jwt do Kong
- Validação de expiração (exp)
- Proteção de endpoints protegidos

### 7. Logging Centralizado ✅
- Logging estruturado via file-log plugin
- Correlation ID único por requisição (header `X-Correlation-ID`)
- Logs via stdout para integração com sistemas de log centralizados

## Estrutura de Rotas

O gateway expõe as seguintes rotas:

- `/api/authentication/*` → Backend (Público - sem JWT)
- `/api/user/*` → Backend (Protegido - requer JWT)
- `/api/tenant/*` → Backend (Protegido - requer JWT)
- `/api/storage/*` → Backend (Protegido - requer JWT)

## Configuração

### Variáveis de Ambiente Necessárias

O gateway precisa das seguintes variáveis de ambiente:

- `BACKEND_URL`: URL do backend (padrão: `http://host.docker.internal:8081`)

**Exemplo de configuração:**
```bash
export BACKEND_URL=http://localhost:8081
```

### Timeouts Configurados

- User service: `5s`
- Tenant service: `10s`
- Storage service: `30s`

### Retry Configurado

- User service: `3` tentativas
- Tenant service: `3` tentativas
- Storage service: `2` tentativas (devido ao timeout maior)
- Backoff: `exponential`
- Retry em: erros `5xx` e `429`

## Desenvolvimento

Para desenvolvimento:

```bash
docker-compose up hr-app-gateway
```

O gateway estará disponível em:
- Proxy: `http://localhost:8000`
- Admin API: `http://localhost:8001`

## Produção

Para produção, use a mesma imagem `kong:3.4` ou uma versão específica.

## Logs

Os logs são formatados e enviados para stdout. Cada requisição recebe um Correlation ID único no header `X-Correlation-ID`.

## Correlation ID

Cada requisição recebe automaticamente um Correlation ID único no header `X-Correlation-ID`. Este ID é:
- Gerado automaticamente pelo gateway (UUID)
- Incluído nos logs
- Propagado para o backend
- Útil para rastrear requisições entre serviços

## Próximos Passos

Os seguintes recursos ainda precisam ser implementados:
- [ ] Configuração completa de JWT com OIDC (atualmente apenas validação básica)
- [ ] Rate limiting (por IP, token, rota, tenant)
- [ ] TLS/HTTPS terminado no gateway
- [ ] CORS e transformações de headers mais avançadas
- [ ] Balanceamento de carga (round-robin, least-connections, IP hash)
