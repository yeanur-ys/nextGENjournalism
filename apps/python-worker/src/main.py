from __future__ import annotations

import time

from config import load_config
from louvain import compute_louvain_clusters


def run() -> None:
    config = load_config()
    while True:
        clusters = compute_louvain_clusters([])
        print(
            f"worker heartbeat interval={config.poll_interval_seconds}s clusters={len(clusters)}",
            flush=True,
        )
        time.sleep(config.poll_interval_seconds)


if __name__ == "__main__":
    run()
