# GEMINI.md - Mandato Fundacional

Este documento contém as diretrizes fundamentais para o desenvolvimento do projeto **Codex Mundi**. Estas instruções têm precedência absoluta sobre fluxos genéricos.

## Diretrizes de Desenvolvimento

1. **Arquitetura Orientada a Dados**: O jogo deve ser construído sobre uma base estatística sólida (Política, Economia, População).
2. **Deterministic Engine**: Cálculos de motor (crescimento, taxas, migração) devem ser determinísticos sempre que possível para garantir consistência e economizar tokens de IA.
3. **IA como Narradora e Modificadora**: A IA deve ser usada para interpretar ações textuais e gerar "Modificadores de Contexto" e narrativas, não para cálculos matemáticos brutos.
4. **Classes Sociais Dinâmicas**: O sistema deve permitir a definição de classes sociais variadas por Era, suportando migração entre elas.
5. **pkg/dice**: Utilize este pacote para injetar fatores de aleatoriedade controlada em eventos determinísticos.

## Regras do Repositório

- **Organização Go**: Siga os padrões do `golang-standards/project-layout` (cmd, internal, pkg, api).
- **Domínio Rico**: As entidades em `internal/domain` devem encapsular a lógica de negócio e estados complexos.
- **Skills**: Utilize ativamente as skills em `.gemini/skills/` para garantir padrões de elite em arquitetura, Go e prompts.

## Referências de Skills Importadas
- `golang-pro`: Padronização de código Go idiomático.
- `architecture-patterns`: Implementação de Clean Architecture e DDD.
- `prompt-engineering-patterns`: Otimização de prompts para reações sistêmicas.
