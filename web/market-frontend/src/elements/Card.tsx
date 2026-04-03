interface ProductCardProps {
    name: string;
    description: string;
    price: string;
    imageUrl: string;
    onClick?: () => void; // обработчик клика
}

export default function ProductCard({ name, description, price, imageUrl, onClick }: ProductCardProps) {
    return (
        <div
            onClick={onClick}
            className="cursor-pointer max-w-xs bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300"
        >
            <img
                src={imageUrl}
                alt={name}
                className="w-full h-48 object-cover"
            />

            <div className="p-4">
                <h2 className="text-lg font-semibold text-gray-800">{name}</h2>
                <p className="text-gray-600 mt-2 text-sm">{description}</p>
                <p className="text-blue-600 font-bold mt-4 text-lg">{price}</p>
            </div>
        </div>
    );
}