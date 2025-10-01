# API Wheels

<p align="center">
  <strong>Uma API RESTful completa para a gestão de serviços em oficinas de veículos.</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go 1.22">
  <img src="https://img.shields.io/badge/Gin_Gonic-1.9-008ECF?style=for-the-badge&logo=gin&logoColor=white" alt="Gin Gonic">
  <img src="https://img.shields.io/badge/PostgreSQL-16-336791?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/Docker-26-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
</p>

## Sobre o Projeto

A **API Wheels** é uma solução de backend desenvolvida em Go para centralizar e otimizar o gerenciamento de uma oficina de veículos. O sistema permite o cadastro e controle de veículos e suas respectivas ordens de serviço através de uma API RESTful robusta, utilizando uma arquitetura limpa e containerização com Docker para garantir portabilidade e escalabilidade.

---

## Funcionalidades Principais

-   **Gestão de Veículos**: CRUD completo (Criar, Ler, Atualizar, Deletar) para os veículos da oficina.
-   **Gestão de Ordens de Serviço**: CRUD completo para as ordens de serviço vinculadas aos veículos.
-   **Consultas Relacionais**: Recupere facilmente todas as ordens de serviço associadas a um veículo específico através da sua placa.
-   **Arquitetura Limpa**: Código organizado em camadas (`controller`, `usecase`, `repository`) para promover a separação de responsabilidades, testabilidade e fácil manutenção.
-   **Ambiente Containerizado**: O ambiente de desenvolvimento e o banco de dados são totalmente gerenciados pelo Docker e Docker Compose, simplificando a configuração e o deploy.

---

## Guia de Instalação e Execução

Para executar o projeto localmente, siga os passos abaixo.

### Pré-requisitos

-   Docker
-   Docker Compose

### Instalação e Execução

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/cauamapurunga/wheels-api
    cd wheels-api
    ```

2.  **Inicie os containers:**
    Este comando irá construir a imagem da aplicação Go, iniciar o container do banco de dados PostgreSQL e executar a API.
    ```bash
    docker-compose up --build -d
    ```
    -   A flag `--build` força a reconstrução da imagem, garantindo que as últimas alterações de código sejam aplicadas.
    -   A flag `-d` executa os containers em segundo plano.

3.  **Disponibilidade da API:**
    A API estará em execução e acessível em `http://localhost:8000`.

Para parar a aplicação e remover os containers, execute:
```bash
docker-compose down
```

---

## Documentação da API

As seções a seguir descrevem os endpoints disponíveis.

### **Recurso: Veículos**

| Método   | Rota              | Descrição                                         |
| :------- | :---------------- | :------------------------------------------------ |
| `GET`    | `/veiculos`       | Retorna uma lista de todos os veículos cadastrados. |
| `GET`    | `/veiculo/{id}`   | Retorna um veículo específico pelo seu ID.        |
| `POST`   | `/veiculo`        | Cria um novo veículo.                             |
| `PUT`    | `/veiculo/{id}`   | Atualiza os dados de um veículo existente.        |
| `DELETE` | `/veiculo/{id}`   | Remove um veículo do sistema.                     |

**Exemplo de corpo (`body`) para requisições `POST` e `PUT`:**
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

### **Recurso: Ordens de Serviço**

| Método   | Rota               | Descrição                                                      |
| :------- | :----------------- | :------------------------------------------------------------- |
| `GET`    | `/servicos`        | Retorna uma lista de todos os serviços cadastrados.            |
| `POST`   | `/servicos`        | Cria uma nova ordem de serviço.                                |
| `GET`    | `/servicos/{placa}`| Retorna todas as ordens de serviço de um veículo específico.   |
| `PUT`    | `/servicos/{id}`   | Atualiza uma ordem de serviço existente.                       |
| `DELETE` | `/servicos/{id}`   | Remove uma ordem de serviço do sistema.                        |

**Exemplo de corpo (`body`) para requisições `POST` e `PUT`:**
```json
{
    "descricao_servico": "Instalação do software Autopilot",
    "custo": 15000.00,
    "data_servico": "2024-10-20T00:00:00Z",
    "veiculo_placa": "NEW-1A23"
}
```

---

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
├── init.sql              # Script de inicialização do banco de dados (criação de tabelas).
└── README.md             # Este arquivo.
```
