import { Link } from "react-router-dom";
import ResolvedImage from "./ResolvedImage.tsx";
import { formatRub } from "../lib/format.ts";
import {
  formatOrderDate,
  orderPreviewImages,
  orderStatusLabel,
  orderStatusTone,
  type OrderDTO,
} from "../lib/orders.ts";

const toneClasses = {
  neutral: "bg-slate-100 text-slate-700",
  active: "bg-sky-100 text-sky-800",
  success: "bg-emerald-100 text-emerald-800",
  danger: "bg-red-100 text-red-800",
} as const;

type Props = {
  order: OrderDTO;
};

export default function OrderCard({ order }: Props) {
  const itemCount = order.items.reduce((s, i) => s + i.quantity, 0);
  const tone = orderStatusTone(order.status);
  const previews = orderPreviewImages(order);
  const extraCount = Math.max(0, order.items.length - previews.length);

  return (
    <Link
      to={`/orders/${order.id}`}
      className="block rounded-xl border border-slate-100 bg-white p-5 shadow-sm hover:border-slate-200 hover:shadow-md transition-all"
    >
      <div className="flex flex-wrap items-start justify-between gap-3">
        <div>
          <p className="font-semibold text-slate-900">Заказ №{order.id}</p>
          <p className="mt-1 text-sm text-slate-500">{formatOrderDate(order.created_at)}</p>
        </div>
        <span className={`inline-flex rounded-full px-2.5 py-0.5 text-xs font-medium ${toneClasses[tone]}`}>
          {orderStatusLabel(order.status)}
        </span>
      </div>

      {previews.length > 0 ? (
        <div className="mt-4 flex items-center gap-2">
          {previews.map((img, i) => (
            <div
              key={`${order.id}-${img}-${i}`}
              className="h-12 w-12 shrink-0 overflow-hidden rounded-lg bg-slate-100 ring-1 ring-slate-100"
            >
              <ResolvedImage imageRef={img} alt="" className="h-full w-full object-cover" />
            </div>
          ))}
          {extraCount > 0 ? (
            <span className="text-xs font-medium text-slate-500">+{extraCount}</span>
          ) : null}
        </div>
      ) : null}

      <p className="mt-4 text-sm text-slate-600">
        {itemCount} {itemCount === 1 ? "товар" : itemCount < 5 ? "товара" : "товаров"}
      </p>
      <p className="mt-2 text-lg font-semibold text-emerald-700">{formatRub(order.total_price)}</p>
    </Link>
  );
}
