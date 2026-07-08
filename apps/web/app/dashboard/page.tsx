import { redirect } from 'next/navigation';

export default async function DashboardRouter({
  searchParams,
}: {
  searchParams: Promise<{ role?: string }>;
}) {
  const params = await searchParams;
  const role = params.role ?? 'reader';

  switch (role) {
    case 'journalist':
    case 'auditor':
    case 'reader':
    case 'admin':
      redirect(`/dashboard/${role}`);
    default:
      redirect('/dashboard/reader');
  }
}
