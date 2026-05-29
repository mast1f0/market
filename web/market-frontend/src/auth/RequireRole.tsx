import { type ReactNode } from "react";
import { Navigate } from "react-router-dom";
import { useAuth } from "./useAuth.ts";

type Props = {
  roles: string[];
  children: ReactNode;
};

export default function RequireRole({ roles, children }: Props) {
  const { token, profile, profileLoading } = useAuth();

  if (!token) {
    return <Navigate to="/login" replace />;
  }

  if (profileLoading) {
    return (
      <div className="max-w-6xl mx-auto p-6 md:p-8">
        <div className="h-10 w-56 bg-slate-200 rounded animate-pulse" />
      </div>
    );
  }

  if (!profile || !roles.includes(profile.role)) {
    return (
      <div className="max-w-3xl mx-auto p-6 md:p-8">
        <h1 className="text-2xl font-bold text-slate-900 mb-2">Нет доступа</h1>
        <p className="text-slate-600">Эта страница недоступна для вашей роли.</p>
      </div>
    );
  }

  return <>{children}</>;
}
