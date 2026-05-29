import { bearerHeaders } from "./api.ts";
import { authApiUrl } from "./endpoints.ts";

export type AuthUser = {
  id: number;
  login: string;
  role: string;
  created_at: string;
};

export const AUTH_ROLES = ["buyer", "seller", "admin"] as const;

export const AUTH_ROLE_LABELS: Record<string, string> = {
  buyer: "Покупатель",
  seller: "Продавец",
  admin: "Администратор",
};

function pick<T>(obj: Record<string, unknown>, ...keys: string[]): T | undefined {
  for (const k of keys) {
    if (obj[k] !== undefined && obj[k] !== null) return obj[k] as T;
  }
  return undefined;
}

export function normalizeAuthUser(raw: Record<string, unknown>): AuthUser {
  return {
    id: Number(pick(raw, "id", "Id") ?? 0),
    login: String(pick(raw, "login", "Login") ?? ""),
    role: String(pick(raw, "role", "Role") ?? ""),
    created_at: String(pick(raw, "createdAt", "created_at", "CreatedAt") ?? ""),
  };
}

export async function fetchAuthUsers(): Promise<AuthUser[]> {
  const res = await fetch(authApiUrl("/users"), {
    headers: { Accept: "application/json", ...bearerHeaders() },
  });
  const text = await res.text();
  if (!res.ok) {
    throw new Error(text || `Ошибка ${res.status}`);
  }
  const raw = JSON.parse(text) as unknown;
  if (!Array.isArray(raw)) return [];
  return raw.map((u) => normalizeAuthUser(u as Record<string, unknown>));
}

export async function updateAuthUserRole(userId: number, role: string): Promise<void> {
  const res = await fetch(authApiUrl(`/users/${userId}`), {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
      ...bearerHeaders(),
    },
    body: JSON.stringify({ role }),
  });
  if (!res.ok) {
    const text = await res.text().catch(() => "");
    throw new Error(text || `Ошибка ${res.status}`);
  }
}

export async function deleteAuthUser(userId: number): Promise<void> {
  const res = await fetch(authApiUrl(`/users/${userId}`), {
    method: "DELETE",
    headers: bearerHeaders(),
  });
  if (!res.ok) {
    const text = await res.text().catch(() => "");
    throw new Error(text || `Ошибка ${res.status}`);
  }
}

export function formatAuthUserDate(iso: string): string {
  if (!iso) return "—";
  const d = new Date(iso);
  if (Number.isNaN(d.getTime())) return "—";
  return new Intl.DateTimeFormat("ru-RU", {
    day: "numeric",
    month: "short",
    year: "numeric",
  }).format(d);
}
