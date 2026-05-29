import { getOrderStatusMeta, ORDER_STATUSES, orderStatusLabel } from "../lib/orders.ts";
import OrderStatusIcon from "./OrderStatusIcon.tsx";
import { STATUS_TONE_STYLES } from "./order-status-styles.ts";

type Props = {
  value: string;
  onChange: (status: string) => void;
  disabled?: boolean;
};

export default function OrderStatusPicker({ value, onChange, disabled }: Props) {
  return (
    <div
      className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-2"
      role="radiogroup"
      aria-label="Статус заказа"
    >
      {ORDER_STATUSES.map((status) => {
        const meta = getOrderStatusMeta(status);
        const styles = STATUS_TONE_STYLES[meta.tone];
        const selected = value === status;

        return (
          <button
            key={status}
            type="button"
            disabled={disabled}
            role="radio"
            aria-checked={selected}
            onClick={() => onChange(status)}
            className={`relative flex items-start gap-3 rounded-xl border-2 p-3 text-left transition-all disabled:opacity-50 disabled:cursor-not-allowed ${
              selected
                ? `border-emerald-500 bg-emerald-50/50 shadow-sm ring-2 ${styles.ring}`
                : "border-slate-100 bg-slate-50/50 hover:border-slate-200 hover:bg-white"
            }`}
          >
            <span
              className={`flex h-9 w-9 shrink-0 items-center justify-center rounded-lg bg-white ring-1 ring-inset ${
                selected ? "ring-emerald-200" : "ring-slate-100"
              } ${styles.icon}`}
            >
              <OrderStatusIcon status={status} className="h-5 w-5" />
            </span>
            <span className="min-w-0 flex-1">
              <span className={`block text-sm font-semibold ${selected ? "text-slate-900" : "text-slate-800"}`}>
                {orderStatusLabel(status)}
              </span>
              <span className="block text-xs text-slate-500 mt-0.5 line-clamp-2">{meta.description}</span>
            </span>
            {selected ? (
              <span className="absolute top-2 right-2 flex h-5 w-5 items-center justify-center rounded-full bg-emerald-600 text-white">
                <svg className="h-3 w-3" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="3" aria-hidden>
                  <path d="m5 12 4 4 10-10" strokeLinecap="round" strokeLinejoin="round" />
                </svg>
              </span>
            ) : null}
          </button>
        );
      })}
    </div>
  );
}
