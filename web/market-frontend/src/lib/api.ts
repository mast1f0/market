import { marketApiUrl } from "./endpoints.ts";
import { ACCESS_TOKEN_KEY } from "./token.ts";

export { marketApiUrl as apiUrl, authApiUrl, minioLinkUrl } from "./endpoints.ts";
export { ACCESS_TOKEN_KEY } from "./token.ts";

export function bearerHeaders(): Record<string, string> {
  const t = localStorage.getItem(ACCESS_TOKEN_KEY);
  return t ? { Authorization: `Bearer ${t}` } : {};
}

export async function fetchJson<T>(path: string, init?: RequestInit, opts?: { auth?: boolean }): Promise<T> {
  const headers = new Headers(init?.headers);
  if (!headers.has("Accept")) headers.set("Accept", "application/json");
  if (opts?.auth) {
    const t = localStorage.getItem(ACCESS_TOKEN_KEY);
    if (t) headers.set("Authorization", `Bearer ${t}`);
  }

  const res = await fetch(marketApiUrl(path), { ...init, headers });
  if (!res.ok) {
    const text = await res.text().catch(() => "");
    throw new Error(text || `HTTP ${res.status}`);
  }
  return res.json() as Promise<T>;
}
