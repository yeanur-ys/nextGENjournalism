'use client';

import { useState } from 'react';
import type { User } from '@/lib/models';

const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? 'http://localhost:8080';

export default function ProfilePage() {
  const [profile, setProfile] = useState<User | null>(null);
  const [error, setError] = useState('');

  async function loadProfile() {
    const token = localStorage.getItem('nextgen_token');
    if (!token) {
      setError('Missing auth token. Please login first.');
      return;
    }

    const res = await fetch(`${apiBaseUrl}/api/v1/profile`, {
      headers: { Authorization: 'Bearer ' + token },
    });

    if (!res.ok) {
      setError('Unable to load profile');
      return;
    }

    const data = await res.json();
    setProfile(data);
  }

  return (
    <main>
      <h1>User profile display (FR-1b)</h1>
      <button onClick={loadProfile}>Load Profile</button>
      {error && <p>{error}</p>}
      {profile && (
        <div>
          <p>Name: {profile.displayName}</p>
          <p>Email: {profile.email}</p>
          <p>Role: {profile.role}</p>
          <p>Verification: {profile.verification}</p>
        </div>
      )}
    </main>
  );
}
