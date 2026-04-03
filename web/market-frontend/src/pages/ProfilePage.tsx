import { Link } from "react-router-dom";

// Тестовые данные пользователя (JWT)
const jwt = {
    role: "seller", // "seller" или "admin"
    userId: 2,
    name: "Иван Иванов",
    email: "ivan@example.com",
};

export default function ProfilePage() {
    return (
        <div className="max-w-3xl mx-auto p-6">
            <h1 className="text-3xl font-bold mb-6">Профиль пользователя</h1>

            <div className="bg-white shadow rounded p-6 mb-6">
                <p className="mb-2"><span className="font-semibold">Имя:</span> {jwt.name}</p>
                <p className="mb-2"><span className="font-semibold">Email:</span> {jwt.email}</p>
                <p className="mb-2"><span className="font-semibold">Роль:</span> {jwt.role === "admin" ? "Администратор" : "Продавец"}</p>
            </div>

            {(jwt.role === "seller" || jwt.role === "admin") && (
                <Link
                    to="/seller-panel"
                    className="inline-block bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
                >
                    Перейти в панель управления товарами
                </Link>
            )}
        </div>
    );
}