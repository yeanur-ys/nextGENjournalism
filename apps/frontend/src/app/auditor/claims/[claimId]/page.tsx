interface ClaimPageProps {
  params: Promise<{ claimId: string }>;
}

export default async function ClaimPage({ params }: ClaimPageProps) {
  const { claimId } = await params;

  return (
    <section>
      <h1>Claim Review</h1>
      <p>Claim ID: {claimId}</p>
    </section>
  );
}
