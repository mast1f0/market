import { useMemo } from "react";
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

function resolveImageSrc(imageRef?: string | null): string {
  const ref = imageRef?.trim() ?? "";
  if (!ref) return PLACEHOLDER;
  if (isDirectImageSource(ref)) return ref;
  return minioLinkUrl(`/view/${encodeURIComponent(ref)}`);
}

type Props = {
  /** Полный URL, data URL или id объекта в minio-link-service. */
  imageRef?: string | null;
  alt?: string;
  className?: string;
};

export default function ResolvedImage({ imageRef, alt = "", className }: Props) {
  const src = useMemo(() => resolveImageSrc(imageRef), [imageRef]);
  const imgClass = className ? `${className} block` : "block";

  return (
    <img
      src={src}
      alt={alt}
      className={imgClass}
      decoding="async"
      onError={(e) => {
        e.currentTarget.onerror = null;
        e.currentTarget.src = PLACEHOLDER;
      }}
    />
  );
}
