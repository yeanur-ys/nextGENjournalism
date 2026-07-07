// Run once against the Neo4j instance (via Neo4j Browser at http://localhost:7474,
// or `cat constraints.cypher | cypher-shell -u neo4j -p nextgen_dev`).
// Go's neo4j-go-driver also runs CREATE CONSTRAINT IF NOT EXISTS safely on every
// api boot (see apps/api/internal/store/neo4j.go), so this file is mostly for
// manual inspection / grading demo purposes.

CREATE CONSTRAINT article_id IF NOT EXISTS 
  FOR (a:Article) REQUIRE a.id IS UNIQUE;
  
CREATE CONSTRAINT user_id IF NOT EXISTS 
  FOR (u:User) REQUIRE u.id IS UNIQUE;

// Sprint 2 will add: (:Article)-[:SEQUENCE_OF]->(:Article)
