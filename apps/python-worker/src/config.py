from dataclasses import dataclass
import os


@dataclass(frozen=True)
class WorkerConfig:
    poll_interval_seconds: int
    neo4j_uri: str


def load_config() -> WorkerConfig:
    return WorkerConfig(
        poll_interval_seconds=int(os.getenv("POLL_INTERVAL_SECONDS", "30")),
        neo4j_uri=os.getenv("NEO4J_URI", "bolt://localhost:7687"),
    )
