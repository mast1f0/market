import { formatRub } from "../lib/format.ts";
import ResolvedImage from "./ResolvedImage.tsx";

interface ProductCardProps {
  name: string;
  description: string;
  price: number;
  imageUrl?: string | null;
  onClick?: () => void;
  onAddToCart?: () => void;
}

export default function ProductCard({
                                      name,
                                      description,
                                      price,
                                      imageUrl,
                                      onClick,
                                      onAddToCart,
                                    }: ProductCardProps) {
  const excerpt =
      description.length > 100
          ? `${description.slice(0, 100).trim()}…`
          : description;

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
          className="cursor-pointer w-full h-full bg-white rounded-xl shadow-sm border border-slate-100 overflow-hidden hover:shadow-md hover:border-slate-200 transition-all duration-200 flex flex-col"
      >
        <div className="w-full h-48 bg-slate-100 flex items-center justify-center">
          <ResolvedImage
              imageRef={imageUrl}
              className="w-full h-full object-contain bg-slate-100 pointer-events-none"
          />
        </div>

        <div className="p-4 flex-1 flex flex-col">
          <h2 className="text-lg font-semibold text-slate-800 leading-snug">
            {name}
          </h2>

          {excerpt ? (
              <p className="text-slate-600 mt-2 text-sm min-h-[3rem] max-h-[3rem] overflow-hidden">
                {excerpt}
              </p>
          ) : null}

          <p className="text-emerald-700 font-semibold mt-auto pt-4 text-lg">
            {formatRub(price)}
          </p>

          {onAddToCart ? (
            <button
              type="button"
              onClick={(e) => {
                e.stopPropagation();
                onAddToCart();
              }}
              className="mt-4 w-full py-2 rounded-lg bg-emerald-600 text-white text-sm font-medium hover:bg-emerald-700 transition-colors"
            >
              В корзину
            </button>
          ) : null}
        </div>
      </article>
  );
}