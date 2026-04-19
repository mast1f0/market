export default function AboutPage() {
  return (
    <div className="max-w-3xl mx-auto px-4 md:px-8 py-10 md:py-14">
      <h1 className="text-3xl font-bold text-slate-900 tracking-tight">О проекте</h1>
      <p className="mt-4 text-slate-600 leading-relaxed">
        «Маркет» — витрина: каталог и карточки товаров ходят в маркет API. В режиме{" "}
        <code className="text-xs bg-slate-100 px-1.5 py-0.5 rounded">npm run dev</code> Vite проксирует{" "}
        <code className="text-xs bg-slate-100 px-1.5 py-0.5 rounded">/api</code> на маркет (порт 8080),{" "}
        <code className="text-xs bg-slate-100 px-1.5 py-0.5 rounded">/auth</code> на auth-microservice (8081),{" "}
        <code className="text-xs bg-slate-100 px-1.5 py-0.5 rounded">/minio</code> на minio-link-service (8082).
      </p>
      <p className="mt-4 text-slate-600 leading-relaxed">
        В продакшене задайте переменные окружения{" "}
        <code className="text-xs bg-slate-100 px-1.5 py-0.5 rounded">VITE_API_URL</code>,{" "}
        <code className="text-xs bg-slate-100 px-1.5 py-0.5 rounded">VITE_AUTH_URL</code>,{" "}
        <code className="text-xs bg-slate-100 px-1.5 py-0.5 rounded">VITE_MINIO_LINK_URL</code> на публичные базовые URL
        сервисов (без завершающего слэша).
      </p>
      <p className="mt-4 text-slate-600 leading-relaxed">
        JWT выдаёт auth-microservice и принимается маркетом при совпадении секрета подписи. В поле{" "}
        <code className="text-xs bg-slate-100 px-1.5 py-0.5 rounded">image_url</code> можно хранить полный URL или id
        объекта: тогда картинка подгружается через minio-link-service (<code className="text-xs bg-slate-100 px-1.5 py-0.5 rounded">GET /file/{"{id}"}</code>).
      </p>
    </div>
  );
}
