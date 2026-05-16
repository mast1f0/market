import { formatRub } from "../lib/format.ts";
import ResolvedImage from "./ResolvedImage.tsx";

interface CartItemProps {
  id: number;
  name: string;
  price: number;
  quantity: number;
  image: string;
  onRemove: (id: number) => void;
  onChangeQuantity: (id: number, quantity: number) => void;
}

export default function CartItem({
  id,
  name,
  price,
  quantity,
  image,
  onRemove,
  onChangeQuantity,
}: CartItemProps) {
  return (
    <div className="flex flex-col sm:flex-row sm:items-center gap-4 p-4 bg-white rounded-xl border border-slate-100 shadow-sm mb-4">
      <ResolvedImage
        imageRef={image}
        alt={name}
        className="w-full sm:w-24 h-40 sm:h-24 object-cover rounded-lg bg-slate-100 shrink-0"
      />

      <div className="flex-1 min-w-0">
        <h3 className="text-lg font-semibold text-slate-900">{name}</h3>
        <p className="text-emerald-700 font-semibold mt-1">{formatRub(price)} за шт.</p>
      </div>

      <div className="flex items-center justify-between sm:justify-end gap-4 sm:gap-2">
        <div className="flex items-center gap-2">
          <button
            type="button"
            className="px-3 py-1.5 rounded-lg border border-slate-200 text-slate-700 hover:bg-slate-50 disabled:opacity-40"
            onClick={() => onChangeQuantity(id, quantity - 1)}
            disabled={quantity <= 1}
            aria-label="Уменьшить количество"
          >
            −
          </button>
          <span className="min-w-[2ch] text-center font-medium tabular-nums">{quantity}</span>
          <button
            type="button"
            className="px-3 py-1.5 rounded-lg border border-slate-200 text-slate-700 hover:bg-slate-50"
            onClick={() => onChangeQuantity(id, quantity + 1)}
            aria-label="Увеличить количество"
          >
            +
          </button>
        </div>

        <button
          type="button"
          className="text-sm font-medium text-red-600 hover:text-red-700"
          onClick={() => onRemove(id)}
        >
          Удалить
        </button>
      </div>
    </div>
  );
}
