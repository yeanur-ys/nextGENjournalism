import { AuthForm } from '@/components/auth-form';

export default function RegisterPage() {
  return (
    <main>
      <h1>User registration (FR-1a)</h1>
      <AuthForm mode="register" />
    </main>
  );
}
