import { useEffect, useState } from "react";
import { minioLinkUrl } from "../lib/endpoints.ts";

const PLACEHOLDER =
  "data:image/svg+xml," +
  encodeURIComponent(
    `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="300" viewBox="0 0 400 300"><rect fill="#e2e8f0" width="400" height="300"/><text x="50%" y="50%" dominant-baseline="middle" text-anchor="middle" fill="#64748b" font-family="system-ui" font-size="14">Нет фото</text></svg>`
  );

function isDirectImageSource(raw: string): boolean {
  const s = raw.trim();
  return /^https?:\/\//i.test(s) || s.startsWith("data:") || s.startsWith("blob:");
}

type Props = {
  /** Полный URL, data URL или id объекта в minio-link-service. */
  imageRef?: string | null;
  alt?: string;
  className?: string;
};

export default function ResolvedImage({ imageRef, alt = "", className }: Props) {
  const [src, setSrc] = useState<string>(() => {
    const t = imageRef?.trim() ?? "";
    if (!t) return PLACEHOLDER;
    return isDirectImageSource(t) ? t : PLACEHOLDER;
  });

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

    let cancelled = false;
    (async () => {
      try {
        const res = await fetch(minioLinkUrl(`/file/${encodeURIComponent(ref)}`));
        if (!res.ok) throw new Error(String(res.status));
        const data = (await res.json()) as { link?: string };
        if (!cancelled && data.link) setSrc(data.link);
        else if (!cancelled) setSrc(PLACEHOLDER);
      } catch {
        if (!cancelled) setSrc(PLACEHOLDER);
      }
    })();

    return () => {
      cancelled = true;
    };
  }, [imageRef]);

  return (
    <img
      src={src}
      alt={alt}
      className={className}
      onError={() => setSrc(PLACEHOLDER)}
    />
  );
}
