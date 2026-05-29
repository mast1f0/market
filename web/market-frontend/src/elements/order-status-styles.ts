import type { OrderStatusTone } from "../lib/orders.ts";

export const STATUS_TONE_STYLES: Record<
  OrderStatusTone,
  {
    badge: string;
    dot: string;
    ring: string;
    icon: string;
    stepDone: string;
    stepActive: string;
    lineDone: string;
  }
> = {
  amber: {
    badge: "bg-amber-50 text-amber-900 ring-amber-200/80",
    dot: "bg-amber-500",
    ring: "ring-amber-400/40",
    icon: "text-amber-600",
    stepDone: "bg-amber-500 border-amber-500",
    stepActive: "bg-amber-500 border-amber-500 ring-4 ring-amber-200",
    lineDone: "bg-amber-300",
  },
  sky: {
    badge: "bg-sky-50 text-sky-900 ring-sky-200/80",
    dot: "bg-sky-500",
    ring: "ring-sky-400/40",
    icon: "text-sky-600",
    stepDone: "bg-sky-500 border-sky-500",
    stepActive: "bg-sky-500 border-sky-500 ring-4 ring-sky-200",
    lineDone: "bg-sky-300",
  },
  violet: {
    badge: "bg-violet-50 text-violet-900 ring-violet-200/80",
    dot: "bg-violet-500",
    ring: "ring-violet-400/40",
    icon: "text-violet-600",
    stepDone: "bg-violet-500 border-violet-500",
    stepActive: "bg-violet-500 border-violet-500 ring-4 ring-violet-200",
    lineDone: "bg-violet-300",
  },
  emerald: {
    badge: "bg-emerald-50 text-emerald-900 ring-emerald-200/80",
    dot: "bg-emerald-500",
    ring: "ring-emerald-400/40",
    icon: "text-emerald-600",
    stepDone: "bg-emerald-500 border-emerald-500",
    stepActive: "bg-emerald-500 border-emerald-500 ring-4 ring-emerald-200",
    lineDone: "bg-emerald-300",
  },
  rose: {
    badge: "bg-rose-50 text-rose-900 ring-rose-200/80",
    dot: "bg-rose-500",
    ring: "ring-rose-400/40",
    icon: "text-rose-600",
    stepDone: "bg-rose-400 border-rose-400",
    stepActive: "bg-rose-500 border-rose-500 ring-4 ring-rose-200",
    lineDone: "bg-rose-200",
  },
  slate: {
    badge: "bg-slate-100 text-slate-700 ring-slate-200/80",
    dot: "bg-slate-400",
    ring: "ring-slate-300/40",
    icon: "text-slate-500",
    stepDone: "bg-slate-400 border-slate-400",
    stepActive: "bg-slate-500 border-slate-500 ring-4 ring-slate-200",
    lineDone: "bg-slate-300",
  },
};
