# Empregufu
Projeto em **Go** para gerenciamento de vagas de emprego, utilizando **Gorilla Mux** e **GORM** com **MySQL**.

---

## 🚀 Pré-requisitos
Antes de começar, certifique-se de ter instalado em sua máquina:

- [Go](https://go.dev/dl/) **1.18 ou superior**  
- [MySQL](https://dev.mysql.com/downloads/) rodando e com um banco criado (ex: `empregUFU`)  
- [Git](https://git-scm.com/)  

---

## ⚙️ Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/empregufu.git
   cd empregufu

2. Inicialize o módulo Go (caso ainda não exista):
   ```bash
   go mod init empregufu

3. Instale as dependencias (caso ainda não exista):
   ```bash
   go mod tidy

4. Instale as dependencias (caso ainda não exista):
    Configure a conexão com o banco de dados.
    No arquivo pkg/config/app.go, ajuste a string de conexão (dsn) com seu usuário, senha e nome do banco.

5. Execute o projeto
   ```bash
    go run main.go

----------------------------------------------
🌐 Servidor

Após rodar o projeto, o servidor estará disponível em: http://localhost:8080

----------------------------------------------

📦 Dependências principais

Gorilla Mux
 — roteador HTTP

GORM
 — ORM para Go

GORM MySQL Driver

---------------------------------------------

🗄️ Banco de Dados

O projeto utiliza MySQL como banco de dados.
Certifique-se de criar previamente o banco (exemplo: empregUFU) antes de rodar a aplicação.

--------------------------------------------

📌 Exemplos de Requisições

Criar um Job
POST /api/v1/jobs

{
  "title": "Desenvolvedor Backend Go",
  "description": "Atuar no desenvolvimento de microserviços em Go",
  "company": "Tech Solutions",
  "location": "Remoto",
  "salary": 8000.50,
  "posted_at": "2025-09-13T10:00:00Z"
}

Listar todos os Jobs
GET /api/v1/jobs

Resposta esperada:

[
  {
    "id": 1,
    "title": "Desenvolvedor Backend Go",
    "description": "Atuar no desenvolvimento de microserviços em Go",
    "company": "Tech Solutions",
    "location": "Remoto",
    "salary": 8000.5,
    "posted_at": "2025-09-13T10:00:00Z"
  }
]