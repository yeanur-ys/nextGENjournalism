export type UserRole = 'journalist' | 'auditor' | 'reader' | 'admin';

export interface User {
  id: string;
  email: string;
  displayName: string;
  role: UserRole;
  bio?: string;
  verification?: string;
  credentialUrl?: string;
  createdAt?: string;
}

export interface Article {
  id: string;
  authorId: string;
  title: string;
  content: string;
  status: 'draft' | 'published' | 'retracted';
  syncedToGraph: boolean;
  createdAt: string;
  updatedAt: string;
}
