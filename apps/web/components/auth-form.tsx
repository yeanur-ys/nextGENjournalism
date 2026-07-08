'use client';

import { FormEvent, useState } from 'react';

const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? 'http://localhost:8080';

type Mode = 'register' | 'login';

export function AuthForm({ mode }: { mode: Mode }) {
  const [message, setMessage] = useState('');

  async function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);

    const payload: Record<string, string> = {
      email: String(formData.get('email') ?? ''),
      password: String(formData.get('password') ?? ''),
    };

    if (mode === 'register') {
      payload.displayName = String(formData.get('displayName') ?? '');
      payload.role = String(formData.get('role') ?? 'reader');
    }

    const path = mode === 'register' ? '/api/v1/auth/register' : '/api/v1/auth/login';
    const res = await fetch(`${apiBaseUrl}${path}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });

    if (!res.ok) {
      const data = await res.json().catch(() => ({ error: 'request failed' }));
      setMessage(data.error ?? 'request failed');
      return;
    }

    const data = await res.json();
    if (data.token) {
      localStorage.setItem('nextgen_token', data.token);
    }
    setMessage('Success');
  }

  return (
    <form onSubmit={onSubmit}>
      <label htmlFor="email">Email</label>
      <input id="email" name="email" type="email" required />
      <label htmlFor="password">Password</label>
      <input id="password" name="password" type="password" required minLength={8} />

      {mode === 'register' && (
        <>
          <label htmlFor="displayName">Display name</label>
          <input id="displayName" name="displayName" required />
          <label htmlFor="role">Role</label>
          <select id="role" name="role" defaultValue="reader">
            <option value="reader">Reader</option>
            <option value="journalist">Journalist</option>
            <option value="auditor">Auditor</option>
            <option value="admin">Admin</option>
          </select>
        </>
      )}

      <button type="submit">{mode === 'register' ? 'Register' : 'Login'}</button>
      {message && <p>{message}</p>}
    </form>
  );
}
