# 🚀 Curso FullCycle - Go Expert

Repositório oficial do curso Go Expert da FullCycle. Contém códigos, exercícios, projetos e recursos para dominar tópicos avançados de Go.

## 📋 Conteúdo Programático

1. **Concorrência Avançada**  
   - Goroutines, Channels, `sync.WaitGroup`, `select`
   - Padrões: Worker Pools, Fan-in/Fan-out
2. **Sistemas Distribuídos**  
   - GRPC, Protocol Buffers, REST APIs
   - Service Discovery, Load Balancing
3. **Testes**  
   - Unitários, Mocks (Testify), Benchmark, Fuzzing
4. **Observabilidade**  
   - OpenTelemetry, Metrics (Prometheus), Logging (Zap/Slog)
5. **Armazenamento de Dados**  
   - SQL (PostgreSQL), NoSQL (MongoDB), Caching (Redis)
6. **Event-Driven & Message Brokers**  
   - RabbitMQ, Kafka, NATS
7. **Deploy & Cloud**  
   - Docker, Kubernetes, AWS/GCP
8. **Segurança**  
   - JWT, OAuth2, Certificados TLS
9. **Boas Práticas**  
   - Clean Architecture, DDD, SOLID

## ⚙️ Pré-requisitos

- Go 1.22+ ([Instalação](https://go.dev/dl/))
- Docker ([Instalação](https://docs.docker.com/engine/install/))
- Ferramentas recomendadas:
  - [Taskfile](https://taskfile.dev/) (para automação)
  - [gRPCurl](https://github.com/fullstorydev/grpcurl)
  - [buf](https://buf.build/) (Protocol Buffers)

## 🏗️ Estrutura do Repositório

```bash
/
├── module-1-concurrency/       # Exemplos de concorrência
├── module-2-grpc/              # Serviços gRPC + Protobuf
├── module-3-observability/     # OpenTelemetry + Prometheus
├── module-4-event-driven/      # Kafka/RabbitMQ consumers
├── challenges/                 # Desafios práticos
├── labs/                       # Laboratórios guiados
└── resources/                  # Slides, diagramas, links
