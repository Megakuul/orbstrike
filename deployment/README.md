# Deployment

In this directory you will find instructions to deploy the orbstrike-cluster on various environments.

There are some key points you will always need to consider on every deployment:

- **Log Rotation**: The orchestrator and gameservers are responsible for managing their own log retention for the *.log files. However, when collecting logs from the console, proper handling is crucial. During major outages or failures (e.g., when games cannot be reallocated due to no gserver being online), logs can quickly become voluminous. While this allows for immediate issue detection, it requires cautious management within your logging software to prevent the rapid overflow and displacement of other logs.

- **Database Latency**: The database is a crucial component of the cluster system. Although games are managed statefully by the gameservers, maintaining low latency to the database (<10ms) remains crucial. Latency issues are particularly noticeable on *focontroller* operations.
