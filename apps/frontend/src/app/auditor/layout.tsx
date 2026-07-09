export default function AuditorLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <section>
      <h1>Auditor Workspace</h1>
      {children}
    </section>
  );
}
