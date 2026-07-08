import { AuthForm } from '@/components/auth-form';

export default function LoginPage() {
  return (
    <main>
      <h1>User login</h1>
      <AuthForm mode="login" />
    </main>
  );
}
