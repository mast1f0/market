/** Маркет API (каталог, корзина, товары). */
export function marketApiUrl(path: string): string {
  const trimmed = path.startsWith("/") ? path : `/${path}`;
  const env = import.meta.env.VITE_API_URL as string | undefined;
  if (env) return `${env.replace(/\/$/, "")}${trimmed}`;
  return `/api${trimmed}`;
}

/** auth-microservice: /login, /register, /profile. */
export function authApiUrl(path: string): string {
  const trimmed = path.startsWith("/") ? path : `/${path}`;
  const env = import.meta.env.VITE_AUTH_URL as string | undefined;
  if (env) return `${env.replace(/\/$/, "")}${trimmed}`;
  return `/auth${trimmed}`;
}

/** minio-link-service: /upload, /file/{id}, … */
export function minioLinkUrl(path: string): string {
  const trimmed = path.startsWith("/") ? path : `/${path}`;
  const env = import.meta.env.VITE_MINIO_LINK_URL as string | undefined;
  if (env) return `${env.replace(/\/$/, "")}${trimmed}`;
  return `/minio${trimmed}`;
}
