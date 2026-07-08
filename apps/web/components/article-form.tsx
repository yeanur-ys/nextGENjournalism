'use client';

import { FormEvent, useState } from 'react';

const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? 'http://localhost:8080';

type ArticleFormProps = {
  mode: 'create' | 'edit';
  articleId?: string;
};

export function ArticleForm({ mode, articleId }: ArticleFormProps) {
  const [message, setMessage] = useState('');

  async function submitArticle(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const token = localStorage.getItem('nextgen_token');
    if (!token) {
      setMessage('Missing auth token. Please login first.');
      return;
    }

    const formData = new FormData(event.currentTarget);
    const payload = {
      title: String(formData.get('title') ?? ''),
      content: String(formData.get('content') ?? ''),
      status: String(formData.get('status') ?? 'draft'),
    };

    const endpoint = mode === 'create' ? '/api/v1/articles' : `/api/v1/articles/${articleId}`;
    const method = mode === 'create' ? 'POST' : 'PUT';

    const res = await fetch(`${apiBaseUrl}${endpoint}`, {
      method,
      headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + token,
      },
      body: JSON.stringify(payload),
    });

    if (!res.ok) {
      const data = await res.json().catch(() => ({ error: 'request failed' }));
      setMessage(data.error ?? 'request failed');
      return;
    }

    setMessage(mode === 'create' ? 'Article created' : 'Article updated');
  }

  return (
    <form onSubmit={submitArticle}>
      <label htmlFor="title">Title</label>
      <input id="title" name="title" required />
      <label htmlFor="content">Content</label>
      <textarea id="content" name="content" rows={8} required />
      <label htmlFor="status">Status</label>
      <select id="status" name="status" defaultValue="draft">
        <option value="draft">Draft</option>
        <option value="published">Published</option>
        <option value="retracted">Retracted</option>
      </select>
      <button type="submit">{mode === 'create' ? 'Create article' : 'Update article'}</button>
      {message && <p>{message}</p>}
    </form>
  );
}
