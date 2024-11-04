# Blockchain-Go

🚀 **Blockchain-Go** é um projeto simples de blockchain escrito em Go, usando o banco de dados embutido **BadgerDB** para persistência. Este projeto implementa funcionalidades básicas de uma blockchain, como criação de blocos e transações entre contas. Abaixo estão as instruções para configuração, execução e uso de comandos essenciais.

## 📋 Pré-requisitos

- **Go** versão 1.23.1 ou superior.
- Configuração de ambiente Go adequada.
- Recomendado: Experiência básica com blockchain e linha de comando.

## 📁 Estrutura do Projeto

O projeto tem a seguinte estrutura básica:

- `main.go`: Arquivo principal com a interface de linha de comando (CLI) para a blockchain.
- `blockchain.go`: Define a estrutura da blockchain e os métodos para adicionar blocos.
- `badger` (biblioteca): Usada para armazenamento persistente de dados.

## 🚀 Começando

Clone este repositório e navegue até a pasta do projeto.

```bash
git clone https://github.com/danubiobwm/blockchain-go
cd blockchain-go
```
## ⚙️ Configuração e Execução
1. Criar uma Blockchain
Para iniciar uma nova blockchain e criar um bloco gênese (primeiro bloco), use o comando createblockchain. O parâmetro -address indica o endereço que receberá a recompensa inicial:
```bash
go run main.go createblockchain -address "jonh"
```
🎉 Isso cria uma nova blockchain com o endereço "jonh" como minerador inicial, recebendo a recompensa do bloco gênese.


## 🔍 2. Consultar Saldo
Para verificar o saldo de um endereço específico, utilize o comando getbalance:

```bash
go run main.go getbalance -address "jonh"
```
## 🔍 Exemplo de Saída:

```bash
badger 2024/11/03 INFO: All 1 tables opened in 2ms
badger 2024/11/03 INFO: Replay took: 0s
Balance of jonh: 100
```

##  🤑 3. Enviar Moedas
Para enviar moedas de um endereço a outro, use o comando send. Certifique-se de que o endereço de origem possui saldo suficiente para cobrir o valor da transação.

```bash
go run main.go send -from "jonh" -to "danubio" -amount 10
```
## 📦 Exemplo de Saída:
```bash
badger 2024/11/03 INFO: All 1 tables opened in 2ms
000008947f0942d7d69a034b8e781a5deee6bd5977737809833d37a007873b72
Sucess!
Balance of jonh: 10
```

## 🎮 Exemplos de Comandos
Aqui estão alguns exemplos de comandos que você pode utilizar para interagir com a blockchain:

### Criar Blockchain:
```bash
go run main.go createblockchain -address "alice"
```
### Consultar Saldo:
```bash
go run main.go getbalance -address "alice"
```
### Enviar Moedas:
```bash
go run main.go send -from "alice" -to "bob" -amount 20
```

##  🛠 Solução de Problemas
#### Erro: not enough funds: Esse erro indica que o endereço de origem não possui saldo suficiente para a transação solicitada.
#### Exibição de Usage ao executar um comando: Certifique-se de que os comandos e parâmetros foram inseridos corretamente, conforme a documentação.

## 📜 Mensagens de Log do BadgerDB

Durante a execução, o BadgerDB exibe mensagens de log informando o status do banco de dados. Essas mensagens são normais e indicam atividades como abertura de tabelas, execução de compacções, e replay de arquivos.

##
✨ Pronto! Agora você está apto a criar sua própria blockchain, consultar saldos e realizar transações básicas entre endereços! Divirta-se explorando as possibilidades e experimentando com os diferentes comandos. Se precisar de ajuda ou encontrar problemas, sinta-se à vontade para revisar o código ou consultar esta documentação.
##

## 💡 Dicas:

Este projeto é uma implementação básica de blockchain e é ideal para estudos.
Para produção, considere adicionar validação de transação, sistema de Proof-of-Work, e mais segurança no armazenamento e manipulação dos dados.

## Desenvolvido com ❤️ e Go! 🐹
Criado por **Danubio**.




Esse `README.md` inclui detalhes para instalação, execução, comandos e exemplos de saída do console, além de dicas úteis para o uso do projeto. Sinta-se à vontade para modificar ou expandir conforme o projeto evoluir!


