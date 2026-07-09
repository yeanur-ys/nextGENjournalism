export default function JournalistLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <section>
      <h1>Journalist Workspace</h1>
      {children}
    </section>
  );
}
