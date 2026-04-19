import { formatRub } from "../lib/format.ts";
import ResolvedImage from "./ResolvedImage.tsx";

interface ProductCardProps {
  name: string;
  description: string;
  price: number;
  imageUrl: string;
  onClick?: () => void;
}

export default function ProductCard({ name, description, price, imageUrl, onClick }: ProductCardProps) {
  const excerpt =
    description.length > 100 ? `${description.slice(0, 100).trim()}…` : description;

  return (
    <article
      role="button"
      tabIndex={0}
      onClick={onClick}
      onKeyDown={(e) => {
        if (e.key === "Enter" || e.key === " ") {
          e.preventDefault();
          onClick?.();
        }
      }}
      className="cursor-pointer max-w-xs bg-white rounded-xl shadow-sm border border-slate-100 overflow-hidden hover:shadow-md hover:border-slate-200 transition-all duration-200"
    >
      <ResolvedImage
        imageRef={imageUrl}
        className="w-full h-48 object-cover bg-slate-100 pointer-events-none"
      />

      <div className="p-4">
        <h2 className="text-lg font-semibold text-slate-800 leading-snug">{name}</h2>
        {excerpt ? <p className="text-slate-600 mt-2 text-sm max-h-[4.5rem] overflow-hidden">{excerpt}</p> : null}
        <p className="text-emerald-700 font-semibold mt-4 text-lg">{formatRub(price)}</p>
      </div>
    </article>
  );
}
