interface ProfilePageProps {
  params: Promise<{ journalistId: string }>;
}

export default async function ProfilePage({ params }: ProfilePageProps) {
  const { journalistId } = await params;

  return (
    <section>
      <h1>Journalist Profile</h1>
      <p>ID: {journalistId}</p>
    </section>
  );
}
