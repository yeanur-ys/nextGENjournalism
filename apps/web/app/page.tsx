import Link from 'next/link';

export default function Home() {
  return (
    <main>
      <h1>nextGENjournalism</h1>
      <p>MVC foundation sprint scaffold.</p>
      <ul>
        <li><Link href="/register">Register</Link></li>
        <li><Link href="/login">Login</Link></li>
        <li><Link href="/profile">Profile</Link></li>
        <li><Link href="/dashboard">Dashboard</Link></li>
        <li><Link href="/articles/new">Create Article</Link></li>
      </ul>
    </main>
  );
}
