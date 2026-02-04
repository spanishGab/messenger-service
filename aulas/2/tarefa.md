# Tarefas - aula 2

## Criar camada controller

Crie uma nova branch contendo o diretório `handlers` contendo os elementos que farão parte da camada de controle da aplicação.

## Criar structs MessageHandler, Command e Result

- Crie a handler MessageHandler que seja capaz de:
  - buscar uma mensagem pelo seu ID
  - buscar várias mensagens filtrando por content, createdAt e timesSent. Além disso fará a paginação dos elementos listados. (Pode ser que vc precise alterar o repository para fazer isso)
  - salvar uma nova mensagem
  - deletar uma mensagem

## Dicas

- Ao invés de retorna uma string e um error na handler, como nosso exemplo da aula, utilize a struct Result para servir como retorno, assim você garante uma resiliência melhor da sua camada de controle
