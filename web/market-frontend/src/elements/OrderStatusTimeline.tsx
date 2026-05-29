import {
  getOrderStatusMeta,
  ORDER_FLOW_STEPS,
  orderStatusLabel,
} from "../lib/orders.ts";
import OrderStatusIcon from "./OrderStatusIcon.tsx";
import { STATUS_TONE_STYLES } from "./order-status-styles.ts";

type Props = {
  status: string;
};

export default function OrderStatusTimeline({ status }: Props) {
  const current = getOrderStatusMeta(status);

  if (status === "cancelled") {
    const styles = STATUS_TONE_STYLES.rose;
    return (
      <div
        className={`rounded-xl border border-rose-100 bg-gradient-to-r from-rose-50 to-white p-4 flex items-start gap-3 ${styles.ring}`}
        role="status"
      >
        <span className={`flex h-10 w-10 items-center justify-center rounded-full bg-white ring-1 ring-rose-200 ${styles.icon}`}>
          <OrderStatusIcon status="cancelled" className="h-5 w-5" />
        </span>
        <div>
          <p className="font-semibold text-rose-900">{current.label}</p>
          <p className="text-sm text-rose-700/90 mt-0.5">{current.description}</p>
        </div>
      </div>
    );
  }

  const activeIndex = current.flowIndex ?? 0;

  return (
    <div className="rounded-xl border border-slate-100 bg-white p-4 sm:p-5 shadow-sm" aria-label="Прогресс заказа">
      <ol className="flex items-start w-full">
        {ORDER_FLOW_STEPS.map((step, index) => {
          const meta = getOrderStatusMeta(step);
          const styles = STATUS_TONE_STYLES[meta.tone];
          const isDone = index < activeIndex;
          const isActive = index === activeIndex;
          const isLast = index === ORDER_FLOW_STEPS.length - 1;

          let nodeClass =
            "flex h-8 w-8 sm:h-9 sm:w-9 items-center justify-center rounded-full border-2 transition-colors ";
          if (isDone) {
            nodeClass += `${styles.stepDone} text-white`;
          } else if (isActive) {
            nodeClass += `${styles.stepActive} text-white`;
          } else {
            nodeClass += "border-slate-200 bg-white text-slate-400";
          }

          let lineClass = "h-0.5 flex-1 min-w-[0.5rem] mx-1 sm:mx-2 mt-4 sm:mt-[1.125rem] rounded-full ";
          if (isDone) {
            lineClass += styles.lineDone;
          } else if (isActive) {
            lineClass += `bg-gradient-to-r ${styles.lineDone} to-slate-100`;
          } else {
            lineClass += "bg-slate-100";
          }

          return (
            <li key={step} className={`flex items-start ${isLast ? "shrink-0" : "flex-1 min-w-0"}`}>
              <div className="flex flex-col items-center shrink-0 w-8 sm:w-auto sm:min-w-[4.5rem]">
                <div className={nodeClass}>
                  {isDone ? (
                    <svg className="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.5" aria-hidden>
                      <path d="m5 12 4 4 10-10" strokeLinecap="round" strokeLinejoin="round" />
                    </svg>
                  ) : (
                    <OrderStatusIcon
                      status={step}
                      className={`h-4 w-4 ${isActive ? "" : "opacity-70"}`}
                    />
                  )}
                </div>
                <span
                  className={`mt-2 text-center text-[10px] sm:text-xs leading-tight max-w-[4.5rem] sm:max-w-none ${
                    isActive ? "font-semibold text-slate-900" : isDone ? "font-medium text-slate-600" : "text-slate-400"
                  }`}
                >
                  {orderStatusLabel(step)}
                </span>
              </div>
              {!isLast ? <div className={lineClass} aria-hidden /> : null}
            </li>
          );
        })}
      </ol>
      <p className="mt-4 text-sm text-slate-600 border-t border-slate-50 pt-3">{current.description}</p>
    </div>
  );
}
