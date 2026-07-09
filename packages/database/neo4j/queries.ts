export const linkArticleToAuthor = `
MERGE (j:Journalist {id: $journalistId})
MERGE (a:Article {id: $articleId})
MERGE (j)-[:AUTHORED]->(a)
`;

export const relateLineage = `
MATCH (parent:Article {id: $parentId})
MATCH (child:Article {id: $childId})
MERGE (parent)-[:LINEAGE_OF]->(child)
`;
