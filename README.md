
# api-go

Projeto simples trabalhado com GoLang. É possível Criar, editar, buscar, editar e deletar
dados usando API.

## Aprendizados

Primeira experiência na criação de API com Go. O principal desafio foi a própria sintaxe da linguagem.


## Stack utilizada

**Back-end:** Go, Gorm, Docker


## Documentação da API

#### Retorna todas as personalidades

```
  GET /api/personalidades
```

#### Retorna uma personalidade

```
  GET /api/personalidades/{id}
```

#### Cria uma personalidade

```
  POST /api/personalidades
```
| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `nome` | `string` |Nome da personalidade |
| `historia` | `string` |historia da personalidade |

#### Edita uma personalidade passando o id

```
  PUT /api/personalidades/{id}
```
| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `nome` | `string` |Alteração |
| `historia` | `string` |Alteração |

#### Deleta uma personalidade passando o id

```
  DELETE /api/personalidades/{id}
```

## Autores

- [Eduardo Lemes](https://github.com/eduardo-lemes)

