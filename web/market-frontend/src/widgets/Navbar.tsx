export default function Navbar() {
    return (
        <header className="flex items-center justify-between px-6 py-4 bg-gray-100 shadow-md">
            <img
                src="https://via.placeholder.com/50"
                alt="Logo"
                className="h-10 w-10"
            />

            <nav className="flex gap-6">
                <a href="/" className="text-gray-700 hover:text-blue-500 font-medium">
                    Главная
                </a>
                <a href="/categories" className="text-gray-700 hover:text-blue-500 font-medium">
                    Категории
                </a>
                <a href="/about" className="text-gray-700 hover:text-blue-500 font-medium">
                    О нас
                </a>
                <a href="/cart" className="text-gray-700 hover:text-blue-500 font-medium">
                    Корзина
                </a>
            </nav>

            <a href="/profile">
                <img
                    src="https://via.placeholder.com/40"
                    alt="Profile"
                    className="h-10 w-10 rounded-full"
                />
            </a>
        </header>
    );
}