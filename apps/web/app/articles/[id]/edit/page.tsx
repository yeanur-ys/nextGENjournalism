import { ArticleForm } from '@/components/article-form';

export default async function EditArticlePage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;

  return (
    <main>
      <h1>Article editing interface (FR-3b)</h1>
      <ArticleForm mode="edit" articleId={id} />
    </main>
  );
}
