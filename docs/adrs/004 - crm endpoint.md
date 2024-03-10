# folders

## Context and Problem Statement

O endpoint `/v2/internal-notifications` foi desconstruido, porém, ainda estamos usando como o endpoint pricipal, porque o novo ainda não esta funcioanndo, mas deixamos um fallback para o novo quando o antigo parar. Enquanto isso, vamos buscar mais informações sobre o novo endpoint.

## Arquivo

internal/infrastructure/communicatorapi/api.go::sendMessage