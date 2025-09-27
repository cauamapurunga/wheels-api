# API Wheels

Uma API RESTful desenvolvida em Go para o gerenciamento de uma oficina de veículos. O sistema permite o cadastro e controle de veículos e suas respectivas ordens de serviço, utilizando uma arquitetura em camadas e containerização com Docker.

## Funcionalidades

- **Gerenciamento de Veículos**: CRUD completo (Criar, Ler, Atualizar, Deletar) para veículos.
- **Gerenciamento de Ordens de Serviço**: CRUD completo (Criar, Ler, Atualizar, Deletar) para ordens de serviço.
- **Consultas Relacionais**: Recuperação de ordens de serviço associadas a um veículo específico por sua placa.
- **Arquitetura Limpa**: Código organizado em camadas (Controller, Usecase, Repository) para separação de responsabilidades e manutenibilidade.
- **Containerização**: Ambiente de desenvolvimento e banco de dados totalmente gerenciado pelo Docker e Docker Compose.

## Tecnologias Utilizadas

- **Linguagem**: Go
- **Framework Web**: Gin Gonic
- **Banco de Dados**: PostgreSQL
- **Driver de Banco de Dados**: pq
- **Containerização**: Docker & Docker Compose

## Como Executar o Projeto

### Pré-requisitos

As seguintes ferramentas devem estar instaladas em seu sistema:
- Docker
- Docker Compose

### Instalação e Execução

1.  **Clone o repositório:**
    ```bash
    git clone <url-do-repositorio>
    cd wheels-api
    ```

2.  **Inicie os containers:**
    Este comando irá construir a imagem da aplicação Go, iniciar o container do banco de dados PostgreSQL e executar a API.
    ```bash
    docker-compose up --build -d
    ```
    - A flag `--build` força a reconstrução da imagem da aplicação, garantindo que as últimas alterações de código sejam aplicadas.
    - A flag `-d` executa os containers em modo "detached" (segundo plano).

3.  **Disponibilidade da API:**
    A API estará em execução e acessível em `http://localhost:8000`.

Para parar a aplicação, execute o seguinte comando:
```bash
docker-compose down
```

## Documentação da API

As seções a seguir descrevem os endpoints disponíveis na API.

### Recurso: Veículos

| Método | Rota | Descrição |
| :--- | :--- | :--- |
| `GET` | `/veiculos` | Retorna uma lista de todos os veículos cadastrados. |
| `GET` | `/veiculo/{id}` | Retorna um veículo específico pelo seu ID. |
| `POST` | `/veiculo` | Cria um novo veículo. |
| `PUT` | `/veiculo/{id}` | Atualiza os dados de um veículo existente. |
| `DELETE` | `/veiculo/{id}` | Remove um veículo do sistema. |

**Exemplo de corpo para requisições `POST` e `PUT`:**
```json
{
    "placa": "NEW-1A23",
    "marca": "Tesla",
    "modelo": "Model Y",
    "ano_fabricacao": 2024,
    "cor": "Branco Perolado",
    "nome_proprietario": "Elon Musk"
}
```

---

### Recurso: Ordens de Serviço

| Método | Rota | Descrição |
| :--- | :--- | :--- |
| `POST` | `/servicos` | Cria uma nova ordem de serviço. |
| `GET` | `/servicos/{placa}` | Retorna todas as ordens de serviço de um veículo específico. |
| `PUT` | `/servicos/{id}` | Atualiza uma ordem de serviço existente. |
| `DELETE` | `/servicos/{id}` | Remove uma ordem de serviço do sistema. |

**Exemplo de corpo para requisições `POST` e `PUT`:**
```json
{
    "descricao_servico": "Instalação do software Autopilot",
    "custo": 15000.00,
    "data_servico": "2024-10-20T00:00:00Z",
    "veiculo_placa": "NEW-1A23"
}
```

## Estrutura do Projeto

O projeto segue uma arquitetura em camadas para promover a organização e o desacoplamento do código:

```
/
├── cmd/
│   └── main.go           # Ponto de entrada da aplicação, configuração do Gin e injeção de dependências.
├── controller/           # Camada de apresentação (HTTP Handlers). Recebe requisições e retorna respostas.
├── db/
│   └── conn.go           # Lógica de conexão com o banco de dados.
├── model/                # Structs que representam as entidades do domínio (Veiculo, OrdemServico).
├── repository/           # Camada de acesso a dados. Contém a lógica de queries SQL.
├── usecase/              # Camada de negócio. Orquestra as operações e regras de negócio.
├── .gitignore
├── docker-compose.yml    # Orquestração para os containers da aplicação e do banco de dados.
├── Dockerfile            # Define o processo de build da imagem Docker da aplicação Go.
├── go.mod
├── go.sum
├── init.sql              # Script de inicialização do banco de dados (criação de tabelas e dados iniciais).
└── README.md             # Este arquivo.
```
