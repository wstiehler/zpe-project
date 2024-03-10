# metrics

## Context and Problem Statement

Alertas do grafana podem repetir o fingerprint, mesmo que seja outro alerta. Isso faz com que a regra de 1 alerta por numero quebre.


## Decision Drivers

- Iremos usar o fingerprint do alerta + uma parte da data para identificar se o alerta ja foi disparado ou nao.
- Não iremos fazer update de alerta no banco, que seria o caminho ideal, porem, vai ser uma alteração grande e que pode quebrar o fluxo de alertas.
- Deixaremos uma trava por hora, ou seja, se o alerta ja foi disparado, ele não sera disparado novamente na mesma hora.