import { useEffect, useState } from "react";
import { minioLinkUrl } from "../lib/endpoints.ts";

const PLACEHOLDER =
  "data:image/svg+xml," +
  encodeURIComponent(
    `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="300" viewBox="0 0 400 300"><rect fill="#e2e8f0" width="400" height="300"/><text x="50%" y="50%" dominant-baseline="middle" text-anchor="middle" fill="#64748b" font-family="system-ui" font-size="14">Нет фото</text></svg>`
  );

/** Кэш minio id → URL, чтобы при переходе с карточки на товар не мигал плейсхолдер. */
const linkCache = new Map<string, string>();

function isDirectImageSource(raw: string): boolean {
  const s = raw.trim();
  return /^https?:\/\//i.test(s) || s.startsWith("data:") || s.startsWith("blob:");
}

function initialSrc(imageRef?: string | null): string {
  const ref = imageRef?.trim() ?? "";
  if (!ref) return PLACEHOLDER;
  if (isDirectImageSource(ref)) return ref;
  return linkCache.get(ref) ?? PLACEHOLDER;
}

type Props = {
  /** Полный URL, data URL или id объекта в minio-link-service. */
  imageRef?: string | null;
  alt?: string;
  className?: string;
};

export default function ResolvedImage({ imageRef, alt = "", className }: Props) {
  const [src, setSrc] = useState(() => initialSrc(imageRef));

  useEffect(() => {
    const ref = imageRef?.trim() ?? "";
    if (!ref) {
      setSrc(PLACEHOLDER);
      return;
    }
    if (isDirectImageSource(ref)) {
      setSrc(ref);
      return;
    }

    const cached = linkCache.get(ref);
    if (cached) {
      setSrc(cached);
      return;
    }

    let cancelled = false;
    (async () => {
      try {
        const res = await fetch(minioLinkUrl(`/file/${encodeURIComponent(ref)}`));
        if (!res.ok) throw new Error(String(res.status));
        const data = (await res.json()) as { link?: string };
        if (cancelled) return;
        if (data.link) {
          linkCache.set(ref, data.link);
          setSrc(data.link);
        } else {
          setSrc(PLACEHOLDER);
        }
      } catch {
        if (!cancelled) setSrc(PLACEHOLDER);
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [imageRef]);

  const imgClass = className ? `${className} block` : "block";

  return (
    <img
      src={src}
      alt={alt}
      className={imgClass}
      decoding="async"
      onError={() => setSrc(PLACEHOLDER)}
    />
  );
}
