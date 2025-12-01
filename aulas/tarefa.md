# Tarefas - aula 1

## Criar estrutura

Crie uma nova branch contendo a estrutura de diretórios mencionada na aula.


## Criar entidade Message e repositório MessageRepository

* Crie a entidade Message contendo os atributos descritos em aula.
* Crie o repositório MessageRepository que seja capaz de:
  - buscar uma mensagem pelo seu ID
  - buscar várias mensagens filtrando por content, createdAt e timesSent
  - salvar uma nova mensagem no banco
  - deletar uma mensagem do banco

## Dicas

* Lembre-se que nosso banco de dados a princípio será um arquivo JSON. Ele pode ficar dentro de uma pasta `database` ou `db`.
* Lembre-se que para criar ou deletar mensagens do banco JSON você vai precisar sobrescrever todos os dados do arquivo JSON após inserir/deletar mensagens
