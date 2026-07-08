import type { UserRole } from '@/lib/models';

export default async function RoleDashboard({
  params,
}: {
  params: Promise<{ role: UserRole }>;
}) {
  const { role } = await params;

  return (
    <main>
      <h1>Role dashboard routing (FR-2b)</h1>
      <p>Active dashboard: {role}</p>
    </main>
  );
}
