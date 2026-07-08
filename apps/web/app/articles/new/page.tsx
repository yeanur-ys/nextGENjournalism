import { ArticleForm } from '@/components/article-form';

export default function NewArticlePage() {
  return (
    <main>
      <h1>Article creation form (FR-3a)</h1>
      <ArticleForm mode="create" />
    </main>
  );
}
