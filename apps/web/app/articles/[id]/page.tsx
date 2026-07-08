import type { Article } from '@/lib/models';

const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? 'http://localhost:8080';

export default async function ArticleViewPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;
  const res = await fetch(`${apiBaseUrl}/api/v1/articles/${id}`, { cache: 'no-store' });

  if (!res.ok) {
    return (
      <main>
        <h1>Article view/display (FR-3c)</h1>
        <p>Article not found.</p>
      </main>
    );
  }

  const article: Article = await res.json();

  return (
    <main>
      <h1>{article.title}</h1>
      <p>Status: {article.status}</p>
      <article>{article.content}</article>
    </main>
  );
}
