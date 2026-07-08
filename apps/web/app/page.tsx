import styles from "./page.module.css";

export default function Home() {
  const pillars = [
    {
      title: "Traceable Story Lineage",
      description:
        "Every article carries linked context, revisions, and source ancestry so readers can audit where claims came from.",
    },
    {
      title: "Relationship Graph Engine",
      description:
        "Entity and account connections are mapped across PostgreSQL and Neo4j to expose influence patterns in real time.",
    },
    {
      title: "Consensus Fact Signals",
      description:
        "Journalists, reviewers, and readers contribute weighted verification signals that strengthen trust over time.",
    },
  ];

  const workflow = [
    "Capture source notes and first draft in the newsroom workspace.",
    "Run lineage checks to connect prior reporting and referenced entities.",
    "Publish with confidence indicators, source map, and live trust score.",
  ];

  const roadmap = [
    { phase: "Week 1", focus: "Core publishing shell, authentication, and role setup." },
    { phase: "Week 2", focus: "Article lineage timeline and revision comparison views." },
    { phase: "Week 3", focus: "Graph-based relationship explorer with WebGL rendering." },
    { phase: "Week 4+", focus: "Consensus moderation, reputation weighting, and analytics dashboard." },
  ];

  return (
    <main className={styles.page}>
      <section className={styles.hero}>
        <p className={styles.kicker}>nextGENjournalism</p>
        <h1>Transparent journalism with visible trust and lineage.</h1>
        <p className={styles.subtitle}>
          A frontend foundation to visualize functionality, track credibility, and present journalism workflows
          clearly from day one.
        </p>
        <div className={styles.ctas}>
          <a href="#platform">Explore Platform</a>
          <a href="#roadmap" className={styles.secondary}>
            View Roadmap
          </a>
        </div>
      </section>

      <section id="platform" className={styles.section}>
        <h2>Platform Pillars</h2>
        <div className={styles.grid}>
          {pillars.map((pillar) => (
            <article key={pillar.title} className={styles.card}>
              <h3>{pillar.title}</h3>
              <p>{pillar.description}</p>
            </article>
          ))}
        </div>
      </section>

      <section className={styles.section}>
        <h2>Publishing Workflow</h2>
        <ol className={styles.timeline}>
          {workflow.map((step) => (
            <li key={step}>{step}</li>
          ))}
        </ol>
      </section>

      <section id="roadmap" className={styles.section}>
        <h2>Build Plan Beyond Week 1</h2>
        <div className={styles.roadmap}>
          {roadmap.map((item) => (
            <article key={item.phase} className={styles.roadmapItem}>
              <h3>{item.phase}</h3>
              <p>{item.focus}</p>
            </article>
          ))}
        </div>
      </section>
    </main>
  );
}
