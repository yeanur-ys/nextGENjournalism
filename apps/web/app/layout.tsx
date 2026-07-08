import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "nextGENjournalism",
  description: "Sprint 1 MVC foundation scaffold",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
