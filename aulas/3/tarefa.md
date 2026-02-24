# Tarefas - aula 3

## Criar camada view

Crie uma nova branch contendo o diretório `cmd` contendo os elementos que farão parte da camada de view da aplicação.


## Criar struct MessageCommand e seus parsers

* Crie o cmd de messagens (MessageCommand) fazendo com que ele saiba interpretar os comandos vindos da linha de comando, incluindo:
    * Validação do nome do comando
    * Validação dos dados respectivos ao comando
    * Crie uma opção extra para que o usuário possa escolher o formato de saída. As opções devem ser: json, yaml e table

## Dicas

* Você pode buscar uma lib para parsear para os formatos yaml e table (markdown)
