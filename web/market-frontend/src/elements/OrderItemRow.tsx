import { Link } from "react-router-dom";
import ResolvedImage from "./ResolvedImage.tsx";
import { formatRub } from "../lib/format.ts";
import type { OrderItemDTO } from "../lib/orders.ts";

type Props = {
  item: OrderItemDTO;
  linkToProduct?: boolean;
};

export default function OrderItemRow({ item, linkToProduct = true }: Props) {
  const title = linkToProduct ? (
    <Link
      to={`/product/${item.product_id}`}
      className="font-medium text-slate-900 hover:text-emerald-700 line-clamp-2"
    >
      {item.name_snapshot}
    </Link>
  ) : (
    <p className="font-medium text-slate-900 line-clamp-2">{item.name_snapshot}</p>
  );

  return (
    <li className="flex gap-4 p-4 rounded-xl border border-slate-100 bg-white shadow-sm">
      <div className="w-20 h-20 sm:w-24 sm:h-24 shrink-0 rounded-lg overflow-hidden bg-slate-100 ring-1 ring-slate-100">
        <ResolvedImage
          imageRef={item.image_snapshot}
          alt={item.name_snapshot}
          className="w-full h-full object-cover"
        />
      </div>
      <div className="flex-1 min-w-0 flex flex-col justify-center">
        {title}
        <p className="mt-1 text-sm text-slate-500">
          {item.quantity} × {formatRub(item.price_snapshot)}
        </p>
      </div>
      <p className="font-semibold text-slate-900 shrink-0 self-center">
        {formatRub(item.price_snapshot * item.quantity)}
      </p>
    </li>
  );
}
