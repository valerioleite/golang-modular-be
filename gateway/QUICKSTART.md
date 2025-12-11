# Guia Rápido - API Gateway KrakenD

## Iniciando o Gateway

### Pré-requisitos

1. Docker e Docker Compose instalados
2. Variáveis de ambiente configuradas no arquivo `.env` na raiz do projeto:
   - `BACKEND_URL`: URL do backend (padrão: `http://hr-app-backend:8080`)
   - `OIDC_ISSUER_URL`: URL do provedor OIDC (ex: `https://your-domain.auth0.com`)
   - `OIDC_CLIENT_ID`: Client ID do OIDC

### Iniciar todos os serviços

```bash
docker-compose up -d
```

Isso iniciará:
- PostgreSQL (porta 5432)
- Backend HR App (porta 8080)
- Gateway KrakenD (porta 8000)

### Iniciar apenas o gateway

```bash
docker-compose up hr-app-gateway
```

## Testando o Gateway

### Health Check

```bash
curl http://localhost:8000/v1/actuator/health
```

### Endpoint público (sem autenticação)

```bash
curl http://localhost:8000/v1/authentication/authorize
```

### Endpoint protegido (requer JWT)

```bash
curl -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  http://localhost:8000/v1/authentication/userinfo
```

## Configurando Backend na Máquina Local

Se o backend está rodando na sua máquina local (não no Docker), use:

**macOS/Windows:**
```bash
export BACKEND_URL=http://host.docker.internal:8081
```

**Linux:**
```bash
export BACKEND_URL=http://172.17.0.1:8081
# ou use o IP da sua interface de rede
```

Depois inicie o gateway:
```bash
docker-compose up hr-app-gateway
```

## Verificando Logs

### Logs do gateway

```bash
docker-compose logs -f hr-app-gateway
```

### Logs de todos os serviços

```bash
docker-compose logs -f
```

## Estrutura de Rotas

O gateway usa roteamento por prefixos. Todas as rotas são automaticamente redirecionadas:

- **Autenticação**: `/api/authentication/*` → Backend (Público)
- **Usuários**: `/api/user/*` → Backend (Protegido)
- **Tenants**: `/api/tenant/*` → Backend (Protegido)
- **Storage**: `/api/storage/*` → Backend (Protegido)

**Exemplo:**
- `GET /api/user/v1/users/sub/123` → `http://hr-app-backend:8080/api/user/v1/users/sub/123`
- `POST /api/tenant/v1/tenants` → `http://hr-app-backend:8080/api/tenant/v1/tenants`

## Desenvolvimento

Para desenvolvimento, o gateway usa a imagem `krakend:watch` que recarrega automaticamente quando o arquivo `krakend.json` é alterado.

Para produção, altere no `docker-compose.yaml` para usar `krakend:2.12.0`.

## Troubleshooting

### Gateway não inicia

1. Verifique se as variáveis de ambiente estão configuradas
2. Verifique se o backend está rodando: `docker-compose ps`
3. Verifique os logs: `docker-compose logs hr-app-gateway`

### Erro de autenticação

1. Verifique se `OIDC_ISSUER_URL` e `OIDC_CLIENT_ID` estão corretos
2. Verifique se o token JWT é válido
3. Verifique os logs do gateway para mais detalhes

### Timeout em requisições

1. Verifique se o backend está respondendo: `curl http://localhost:8080/api/user/v1/actuator/health`
2. Aumente o timeout no `krakend.json` se necessário
3. Verifique os logs para identificar o problema

