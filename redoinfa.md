Current Architecture Limitations
The existing service has several bottlenecks that prevent it from scaling to 10K RPS:

- Single Instance: Runs on one server with limited CPU/memory
- In-Memory Storage: Data lost on restart, no persistence
- No Load Distribution: Single point of failure
- Blocking Operations: Report generation can block other requests
- No Caching: Every request triggers full report generation


## Scalable Architecture

<h2> Horizontal Scaling Strategy</h2>

<h4>Service Tier Architecture</h4>

Internet → Load Balancer → API Gateway → gRPC Services → Message Queue → Workers

<h3>Components:</h3>

- Load Balancer - NGINX with gRPC load balancing
- API Gateway -   Envoy Proxy for routing, rate limiting, and observability
- gRPC Service Instances - StateLess Replicas of Grpc
- Message Queue -  Apache Kafka for async report generation
- Worker Pools - Dedicated report generation workers
- Caching Layer -  Redis cluster for report caching
- Database - PostgreSQL cluster for persistent storage Architecture